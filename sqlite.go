package sqlite

import (
	"errors"
	"github.com/go-gdbc/gdbc"
	_ "github.com/mattn/go-sqlite3"
)

const AuthKey = "_auth"
const UserAuthKey = "_auth_user"
const PasswordAuthKey = "_auth_pass"

func init() {
	gdbc.Register("sqlite3", "sqlite", &SqliteSourceNameAdapter{})
}

type SqliteSourceNameAdapter struct {
}

func (dsnAdapter SqliteSourceNameAdapter) GetDataSourceName(dataSource gdbc.DataSource) (string, error) {
	dsn := "file:"
	user := ""
	password := ""

	dataSourceUrl := dataSource.GetURL()

	if dataSourceUrl.Opaque != "" {
		dsn = dsn + dataSourceUrl.Opaque
	} else {
		return "", errors.New("wrong format")
	}

	arguments := dataSourceUrl.Query()
	if arguments.Get(UserAuthKey) != "" {
		user = arguments.Get(UserAuthKey)
	} else {
		user = dataSource.GetUsername()
	}

	if arguments.Get(PasswordAuthKey) != "" {
		password = arguments.Get(PasswordAuthKey)
	} else {
		password = dataSource.GetPassword()
	}

	arguments.Del(AuthKey)
	arguments.Del(UserAuthKey)
	arguments.Del(PasswordAuthKey)

	if user != "" {
		dsn = dsn + "?" + AuthKey + "&" + UserAuthKey + "=" + user + "&" + PasswordAuthKey + "=" + password
	}

	if len(arguments) == 0 {
		return dsn, nil
	}

	if user != "" {
		dsn = dsn + "&"
	} else {
		dsn = dsn + "?"
	}

	for argumentName, values := range arguments {
		dsn = dsn + argumentName + "=" + values[0] + "&"
	}
	dsn = dsn[:len(dsn)-1]

	return dsn, nil
}
