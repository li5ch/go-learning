package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type gormV2Client struct {
	db *gorm.DB
}

var (
	dbV2Client *gormV2Client
)

func BatchInsertV2CollectionLinks(collectionLinkTabs []*CollectionLinkTab) error {
	if len(collectionLinkTabs) == 0 {
		return nil
	}
	return dbV2Client.db.Create(&collectionLinkTabs).Error
}

type UserTab struct {
	ID       uint64
	UserName string  // link name
}

func NewMysqlV2Client() (*gormV2Client, error) {
	address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%ds",
		"root", "", "localhost", "3306", "gozero", defaultConnectTimeoutSecond)
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN: address, // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &gormV2Client{db: gdb}, nil
}

func main()  {

	tabs := []*CollectionLinkTab{
		{
			LinkName: "a",
		},
		{
			LinkName: "b",
		},
		{
			LinkName: "c",
		},
	}
	dbV2Client, _ = NewMysqlV2Client()
	err := BatchInsertV2CollectionLinks(tabs)
	if err!=nil{
		log.Println(err)
	}
	for i:=range tabs{
		log.Println(tabs[i].ID)
	}
	tab1s := []*UserTab{
		{
			UserName: "a",
		},
		{
			UserName: "b",
		},
		{
			UserName: "c",
		},
	}
	dbV2Client.db.AutoMigrate(&UserTab{})
	log.Println("123")
	dbV2Client.db.Create(&tab1s)
	for i:= range tab1s{
		log.Println(tabs[i].ID)
	}
}