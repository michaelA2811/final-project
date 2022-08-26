package config

import (
	"final-project/entity"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// dbDriver := viper.GetString("db.dbdriver")
	dbUser := viper.GetString("db.dbuser")
	dbPass := viper.GetString("db.dbpassword")
	dbName := viper.GetString("db.dbname")

	// fmt.Println(viper.GetString("db.dbdriver"))
	// db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)/"+dbName+"?parseTime=true")
	dsn := dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbName + "?parseTime=true"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.Debug().AutoMigrate(entity.User{})
	if err != nil {
		panic(err.Error())
	}
	return db, nil
}
