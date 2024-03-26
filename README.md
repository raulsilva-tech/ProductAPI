    - migrate (F)
	- SQLC (F)
	- Repository DP (F)
	- UseCase DP (F)
	- Presenter DP
	- TDD : testing units
	- Google Wire Dependency Injection
	- REST end points
	- Token Authentication
	- Clean architecture



MIGRATE:
   
    // Install migrate version with tah sqlite3
        go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    // creating files up and down in the directory 
        migrate create -ext=sql -dir=sql/migrations -seq init
    // execute up file (create tables)
        migrate -path=sql/migrations -database "sqlite3://productapi.db" -verbose up
    // execute down file (drop tables)
        migrate -path=sql/migrations -database "sqlite3://productapi.db" -verbose up
    
SQLC:
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

    create sqlc.yaml file

    create and edit file query.sql and write the queries you want to use

    sqlc generate