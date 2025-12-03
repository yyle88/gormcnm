// Package main demonstrates basic gormcnm usage with Account CRUD operations
// Auto shows column definitions, WHERE queries, UPDATE operations, and expression updates
// Runs with SQLite in-memory database to showcase type-safe column operations
//
// main 包演示 gormcnm 与 Account 增删改查操作的基本用法
// 自动展示列定义、WHERE 查询、UPDATE 操作和表达式更新
// 使用 SQLite 内存数据库运行以展示类型安全的列操作
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

type Account struct {
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

	//CREATE TABLE `accounts` (`username` varchar(100),`nickname` text,`age` integer,PRIMARY KEY (`username`))
	done.Done(db.AutoMigrate(&Account{}))
	//INSERT INTO `accounts` (`username`,`nickname`,`age`) VALUES ("alice","Alice",17)
	done.Done(db.Create(&Account{Username: "alice", Nickname: "Alice", Age: 17}).Error)

	//SELECT * FROM `accounts` WHERE username="alice" ORDER BY `accounts`.`username` LIMIT 1
	var account Account
	done.Done(db.Where(columnUsername.Eq("alice")).First(&account).Error)
	fmt.Println(neatjsons.S(account))

	//UPDATE `accounts` SET `nickname`="Alice-2" WHERE `username` = "alice"
	done.Done(db.Model(&account).Update(columnNickname.Kv("Alice-2")).Error)
	//SELECT * FROM `accounts` WHERE username="alice" ORDER BY `accounts`.`username` LIMIT 1
	done.Done(db.Where(columnUsername.Eq("alice")).First(&account).Error)
	fmt.Println(neatjsons.S(account))

	//UPDATE `accounts` SET `age`=18,`nickname`="Alice-3" WHERE `username` = "alice"
	done.Done(db.Model(&account).Updates(columnNickname.Kw("Alice-3").Kw(columnAge.Kv(18)).AsMap()).Error)
	//SELECT * FROM `accounts` WHERE username="alice" ORDER BY `accounts`.`username` LIMIT 1
	done.Done(db.Where(columnUsername.Eq("alice")).First(&account).Error)
	fmt.Println(neatjsons.S(account))

	//UPDATE `accounts` SET `age`=age + 1 WHERE `username` = "alice"
	done.Done(db.Model(&account).Update(columnAge.KeAdd(1)).Error)
	//SELECT * FROM `accounts` WHERE username="alice" ORDER BY `accounts`.`username` LIMIT 1
	done.Done(db.Where(columnUsername.Eq("alice")).First(&account).Error)
	fmt.Println(neatjsons.S(account))
}
