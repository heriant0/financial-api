# purplestore

## Pkg
- gin   (http fremwork) : github.com/gin-gonic
- sqlx (db interaction) : github.com/jmoiron
- lib pq (db potgres)
- viper 

```
.
├── app.env.sample ( contains application configuration )
├── cmd ( main binary )
├── Makefile ( simplify project commands )
├── docs ( swagger documentation )
├── db
│   └── migrations ( database migrations )
└── internal
    ├── app
    │   ├── controllers ( request response handler )
    │   ├── models ( all about database table )
    │   ├── repository ( database/cache operation )
    │   ├── router ( http router )
    │   ├── schema ( request/response schema )
    │   └── service ( business logic )
    └── pkg ( private lib )
```

```
router --> middleware --> controllers(use schema) --> service --> repository(user model)
```