# Using MySQL

1. Install MySQL
2. We will need a MySQL driver
   - go get -u github.com/go-sql-driver/mysql
   - [read the documentation](https://github.com/go-sql-driver/mysql#installation)
   - [see all SQL drivers](https://github.com/golang/go/wiki/SQLDrivers)
   - [Astaxie's book](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.2.html)
3. Include driver in your imports
   - _ "github.com/go-sql-driver/mysql"
   - [Read the documentation](https://github.com/go-sql-driver/mysql#usage)
4. Determine the Data Source Name
   - user:password@tcp(localhost:5555)/dbname?charset=utf-8
   - [Read the documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name)
5. Open a connection
   - db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf-8)

[package sql](https://godoc.org/database/sql)
