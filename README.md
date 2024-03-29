# go-automatic-chainsaw -> bank-app

## Roadmap

#### Section I : Working with database [Postgres]

🗒 1. [dbdiagram.io](./research_and_development_work/DB_DIAGRAM_SCHEMA.sql) Design DB schema and generate SQL code and PDF as a report.

<p align='start'>
<img src='./research_and_development_work/DB_DIAGRAM_SCHEMA.png' width='800'/>
</p>

🐳 2. [init docker](https://www.docker.com/) create container/images and run for test

-   postgres:latest [hub](https://hub.docker.com/_/postgres)
-   run database migration in Golang [golang-migrate](https://github.com/golang-migrate/migrate/tree/master) after this [check ``` migrate --version```]
    <!-- // -->
    ```sh
    mkdir -p db/migration
    migrate create -ext sql -dir db/migration -seq init_schema  #you will see up/down migration files
    ```
    create [Makefile](./Makefile) in your root (to create/drop db with ease)
    ```sh
    make createdb # you should see db by name bank_db
     # after migrate up to find out migration version and stats
    ls -l db/migration
    ```

🗄️ 3. initial type-safe SQL compiler by [sqlc](https://docs.sqlc.dev/en/latest/index.html)
also you can use ORM like [GORM](https://gorm.io/index.html) but after some research [like.1](https://blog.jetbrains.com/go/2023/04/27/comparing-db-packages/), [like.2](https://dev.to/fadhilradh/using-sqlc-library-for-golang-projects-will-make-your-life-easier-5474),... we decide to use sqlc

<p align='start'>
<img src='./research_and_development_work/ORM_COMPARISON.png' width='800'/>
</p>

    -   install sqlc

```sh
docker pull sqlc/sqlc #then use docker run ...
#or
yay sqlc
```

then init sqlc and generate sqlc config like [here](./sqlc.yaml) , `insure use postgres config`

then create all needed schemas files [like here](./db/schema/) and queries files [like here](./db/query/) based on your db [find sample here](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)

then run

```sh
make sqlc # after this ./db/sqlc directory should contain generated go files [account/entry/transfer/.../db/model]
```
