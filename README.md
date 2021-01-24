# gdbc-sqlite
GDBC Sqlite Driver - It is based on [github.com/mattn/go-sqlite3](github.com/mattn/go-sqlite3)

[![Go Report Card](https://goreportcard.com/badge/github.com/go-gdbc/gdbc-sqlite)](https://goreportcard.com/report/github.com/go-gdbc/gdbc-sqlite)
[![codecov](https://codecov.io/gh/go-gdbc/gdbc-sqlite/branch/main/graph/badge.svg?token=1O1KF6HIHH)](https://codecov.io/gh/go-gdbc/gdbc-sqlite)
[![Build Status](https://travis-ci.com/go-gdbc/gdbc-sqlite.svg?branch=main)](https://travis-ci.com/go-gdbc/gdbc-sqlite)

# Usage
```go
dataSource, err := gdbc.GetDataSource("gdbc:sqlite:test.db?cache=shared&mode=memory")
if err != nil {
    panic(err)
}

var connection *sql.DB
connection, err = dataSource.GetConnection()
if err != nil {
    panic(err)
}
```

Sqlite GDBC URL takes the following form:

```
gdbc:sqlite:database-name?arg1=value1
gdbc:sqlite:database-file-path?arg1=value1
```

Checkout [github.com/mattn/go-sqlite3](github.com/mattn/go-sqlite3) for arguments details.
