gopostgresql
============

The **gopostgresql** package provides a simple and easy-to-use PostgreSQL repository implementation in Go. It handles the connection to the PostgreSQL database and provides basic functionality for querying and managing the database connection.

Installation
------------

    go get github.com/ashkenazi1/gopostgresql


Usage
-----

To use the **gopostgresql** package, follow the steps below:

1.  Import the package into your Go project.
2.  Use the `GetPostgresql` function to establish a connection to your PostgreSQL database.
3.  Use the `Query` method to execute SQL queries.
4.  Use the `Ping` method to check the connection status.
5.  Use the `Close` method to close the database connection when done.

### Example

    
    package main
    
    import (
        "fmt"
        "log"
        "gopostgresql"
    )
    
    func main() {
        db, err := gopostgresql.GetPostgresql("localhost", 5432, "user", "password", "dbname", true)
        if err != nil {
            log.Fatal(err)
        }
        defer db.Close()
    
        rows, err := db.Query("SELECT * FROM my_table")
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()
    
        for rows.Next() {
            var column1 string
            var column2 int
            if err := rows.Scan(&column1, &column2); err != nil {
                log.Fatal(err)
            }
            fmt.Println(column1, column2)
        }
    
        if err := db.Ping(); err != nil {
            log.Fatal("Cannot connect to database:", err)
        }
    }
    

Functions
---------

### GetPostgresql

`func GetPostgresql(Host string, Port int, User string, Password string, Dbname string, SSLMode bool) (*PostgresqlRepository, error)`

Establishes a connection to the PostgreSQL database and returns a singleton instance of the `PostgresqlRepository`. If the connection is already established, it returns the existing instance.

### Query

`func (p *PostgresqlRepository) Query(query string) (*sql.Rows, error)`

Executes a SQL query on the database and returns the result set.

### Ping

`func (p *PostgresqlRepository) Ping() error`

Pings the database to check the connection status.

### Close

`func (p *PostgresqlRepository) Close()`

Closes the connection to the database.

License
-------

This package is available under the MIT License.
