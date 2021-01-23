package sqlite

import (
	"github.com/go-gdbc/gdbc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func getDSN(t *testing.T, dataSourceUrl string) (string, error) {
	adapter := gdbc.GetDataSourceNameAdapter("sqlite")
	dataSource, err := gdbc.GetDataSource(dataSourceUrl)
	assert.Nil(t, err)
	return adapter.GetDataSourceName(dataSource)
}

func getDSNWithUser(t *testing.T, dataSourceUrl string, username string, password string) (string, error) {
	adapter := gdbc.GetDataSourceNameAdapter("sqlite")
	dataSource, err := gdbc.GetDataSource(dataSourceUrl, gdbc.Username(username), gdbc.Password(password))
	assert.Nil(t, err)
	return adapter.GetDataSourceName(dataSource)
}

func TestSqliteDataSourceNameAdapter_GetDataSourceNameWithoutUser(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlite:test.db?cache=shared&mode=memory")
	assert.Nil(t, err)
	assert.Equal(t, dsn, "file:test.db?cache=shared&mode=memory")
}

func TestSqliteDataSourceNameAdapter_GetDataSourceNameWithUserData(t *testing.T) {
	dsn, err := getDSNWithUser(t, "gdbc:sqlite:test.db?mode=memory", "username", "password")
	assert.Nil(t, err)
	assert.Equal(t, dsn, "file:test.db?_auth&_auth_user=username&_auth_pass=password&mode=memory")
}

func TestSqliteDataSourceNameAdapter_GetDataSourceNameWithUserParameter(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlite:test.s3db?_auth&_auth_user=admin&_auth_pass=admin")
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(dsn, "file:test.s3db?"))
	assert.Contains(t, dsn, "_auth")
	assert.Contains(t, dsn, "auth_user=admin")
	assert.Contains(t, dsn, "_auth_pass=admin")
}

func TestSqliteDataSourceNameAdapter_GetDataSourceNameWithWrongFormat(t *testing.T) {
	dsn, err := getDSN(t, "gdbc:sqlite://localhost:3000")
	assert.NotNil(t, err)
	assert.Equal(t, "wrong format", err.Error())
	assert.Empty(t, dsn)
}
