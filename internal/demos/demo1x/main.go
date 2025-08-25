package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	Username string `gorm:"primary_key;type:varchar(100);"`
	Nickname string `gorm:"column:nickname;"`
	Age      int    `gorm:"column:age;"`
}

const (
	columnUsername = gormcnm.ColumnName[string]("username")
	columnNickname = gormcnm.ColumnName[string]("nickname")
	columnAge      = gormcnm.ColumnName[int]("age")
)

func main() {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&User{}))
	done.Done(db.Create(&User{Username: "alice", Nickname: "Alice", Age: 17}).Error)

	var user User
	done.Done(db.Where(columnUsername.Eq("alice")).First(&user).Error)
	fmt.Println(neatjsons.S(user))

	done.Done(db.Model(&user).Update(columnNickname.Kv("SuperAlice")).Error)
	done.Done(db.Where(columnUsername.Eq("alice")).First(&user).Error)
	fmt.Println(neatjsons.S(user))

	done.Done(db.Model(&user).Update(columnAge.KeAdd(1)).Error)
	done.Done(db.Where(columnUsername.Eq("alice")).First(&user).Error)
	fmt.Println(neatjsons.S(user))
}
