package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	defaultConnectTimeoutSecond = 5
)

type gormV1Client struct {
	db *gorm.DB
}

var (
	dbClient *gormV1Client
)

func NewMysqlClient() (*gormV1Client, error) {

	address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%ds",
		"root", "", "localhost", "3306", "gozero", defaultConnectTimeoutSecond)
	gdb, err := gorm.Open("mysql", address)
	if err != nil {
		return nil, err
	}

	return &gormV1Client{db: gdb}, nil
}

func (c *gormV1Client) GetDB() *gorm.DB {
	return c.db
}

type TransactionManager interface {
	Commit()
	GetConn() *gorm.DB
}

type MySQLDBManager struct {
	tx *gorm.DB
}

func NewTransaction(dbClient *gormV1Client) TransactionManager {
	tm := &MySQLDBManager{}
	tm.tx = dbClient.GetDB().Begin()
	return tm
}

func (tm *MySQLDBManager) Commit() {
	tm.tx.Commit()
}

func (tm *MySQLDBManager) GetConn() *gorm.DB {
	return tm.tx
}


type CollectionLinkTab struct {
	ID       uint64 `gorm:"column:id"`
	LinkName string `gorm:"column:link"` // link name

}

func (*CollectionLinkTab) TableName() string {
	return "link_tab"
}

//func BatchInsertCollectionLinks(tm TransactionManager, collectionLinkTabs []*CollectionLinkTab) error {
//	if len(collectionLinkTabs) == 0 {
//		return nil
//	}
//
//
//
//}



//func main() {
//	dbClient, _ = NewMysqlClient()
//	tm := NewTransaction(dbClient)
//
//	tabs := []*CollectionLinkTab{
//		{
//			LinkName: "a",
//		},
//		{
//			LinkName: "b",
//		},
//		{
//			LinkName: "c",
//		},
//	}
//
//	err := BatchInsertCollectionLinks(tm, tabs)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	tm.Commit()
//	for i := range tabs {
//		log.Println(tabs[i].ID)
//	}
//}
