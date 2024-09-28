package test

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestConnectDb(t *testing.T) {
	vipers := initViper()
	dsn := vipers.GetString("db.USERNAME") + ":" + vipers.GetString("db.PASSWORD") + "@tcp(" + vipers.GetString("db.HOST") + ":" + vipers.GetString("db.PORT") + ")/" + vipers.GetString("db.NAME")
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.Nil(t, err)
}
