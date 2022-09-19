package mysql

import (
	"fmt"
	"log"
	"project/dashboard/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBmySql(config *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return db

}

// TEST TO LOCALHOST
// func InitDBMySqlTest() *gorm.DB {

// 	config := map[string]string{
// 		"DB_Username": "",
// 		"DB_Password": "",
// 		"DB_Port":     "",
// 		"DB_Host":     "",
// 		"DB_Name":     "",
// 	}

// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		config["DB_Username"],
// 		config["DB_Password"],
// 		config["DB_Host"],
// 		config["DB_Port"],
// 		config["DB_Name"])

// 	// var e error
// 	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	return db
// }
