package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/heriant0/financial-api/internal/app/controllers"
	"github.com/heriant0/financial-api/internal/app/repositories"
	"github.com/heriant0/financial-api/internal/app/services"
	"github.com/heriant0/financial-api/internal/pkg/config"
	"github.com/heriant0/financial-api/internal/pkg/middlewares"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var cfg config.Config
var dbConn *sqlx.DB
var err error

// var enforcer *casbin.Enforcer

func init() {
	// load configuration based on app.env
	cfg, err = config.LoadConfig()
	if err != nil {
		panic("failed to load config")
	}

	// Create database connection
	dbConn, err = sqlx.Open(cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		errMsg := fmt.Errorf("err database connect: %w", err)
		panic(errMsg)
	}

	err = dbConn.Ping()
	if err != nil {
		errMsg := fmt.Errorf("err database ping: %w", err)
		panic(errMsg)
	}

	// casebin enforcer
	// enforcer, err = casbin.NewEnforcer("config/rbac_model.conf", "config/rbac_policy.csv")
	// if err != nil {
	// 	errMsg := fmt.Errorf("error enforce casbin: %w", err)
	// 	panic(errMsg)
	// }

	// setup logrus logging
	loglLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		errMsg := fmt.Errorf("parse log level : %s", err)
		panic(errMsg)
	}

	log.SetLevel(loglLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {

	r := gin.New()
	r.Use(middlewares.LogMiddleware())

	// init repository
	categoryRepository := repositories.NewCategoryRepository(dbConn)
	userRepository := repositories.NewUserRepository(dbConn)
	currencyRepository := repositories.NewCurrencyRepository(dbConn)
	authRepository := repositories.NewAuthRepository(dbConn)
	transactionRepository := repositories.NewTransactionRepository(dbConn)

	// init service
	categoryService := services.NewCategoryService(categoryRepository)
	currencyService := services.NewCurrencyService(currencyRepository)

	tokenMaker := services.NewTokenMaker(
		cfg.AccessTokenKey,
		cfg.RefreshTokenKey,
		cfg.AccessTokenDuration,
		cfg.RefreshTokenDuration,
	)
	sessionService := services.NewSessionService(userRepository, authRepository, tokenMaker)
	registrationService := services.NewRegistrationService(userRepository)
	transactionService := services.NewTransactironService(transactionRepository)

	// init controller
	categoryController := controllers.NewCategoryController(categoryService)
	registrationControler := controllers.NewRegistrationConroller(registrationService)
	currencyController := controllers.NewCurrencyController(currencyService)
	sessionController := controllers.NewSessionController(sessionService, tokenMaker)
	transactionController := controllers.NewTransactionController(transactionService)

	// routes
	authRoutes := r.Group("api/v1/auths")
	{
		authRoutes.POST("/register", registrationControler.Register)
		authRoutes.POST("/login", sessionController.Login)
		authRoutes.GET("/refresh", sessionController.Refresh)
		authRoutes.POST("/logout", middlewares.AuthenticationMiddleware(tokenMaker), sessionController.Logout)
	}

	r.Use(middlewares.AuthenticationMiddleware(tokenMaker))
	v1Routes := r.Group("api/v1")
	{

		v1Routes.GET("users/:id", registrationControler.UserProfile)

		v1Routes.GET("categories", categoryController.GetList)
		v1Routes.GET("categories/:id", categoryController.Detail)
		v1Routes.POST("categories", categoryController.Create)
		v1Routes.PATCH("categories/:id", categoryController.Update)
		v1Routes.DELETE("categories/:id", categoryController.Delete)

		v1Routes.GET("/currencies", currencyController.GetList)
		v1Routes.GET("/currencies/:id", currencyController.GetByID)
		v1Routes.POST("/currencies", currencyController.Create)
		v1Routes.PATCH("/currencies/:id", currencyController.Update)
		v1Routes.DELETE("/currencies/:id", currencyController.Delete)

		v1Routes.POST("/transactions", transactionController.Create)
		v1Routes.GET("/transactions", transactionController.GetList)
		v1Routes.GET("/transactions/:types", transactionController.GetDataByType)

	}

	appPort := fmt.Sprintf(":%s", cfg.AppPort)
	err := r.Run(appPort)
	if err != nil {
		log.Panic("cannot start the apps")
	}
}
