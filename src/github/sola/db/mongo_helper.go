apackage db

import (
	"fmt"
	"github/sola/repository"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

const (
	DB_Server = "localhost:50000"
	DB_Table  = "sola"
	DB_COL    = "col"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial(DB_Server)
	CheckError(err)
	// defer session.Close()
	// 切换场景到一个monotonic
	session.SetMode(mgo.Monotonic, true)
	return session
}

func GetCollection(db_name string, col_name string) *mgo.Collection {
	session, err := mgo.Dial(DB_Server)
	CheckError(err)
	defer session.Close()
	// 切换场景到一个monotonic
	// session.SetMode(mgo.Monotonic, true)
	c := session.DB(db_name).C(col_name)
	return c
}

func GetCollections(col_name string) *mgo.Collection {
	return GetCollection(DB_Table, col_name)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Find(condition interface{}) (query *mgo.Query, session *mgo.Session) {
	session = getSession()
	query = session.DB(DB_Table).C(DB_COL).Find(condition)
	return query, session
}

func Insert(docs ...interface{}) (err error, session *mgo.Session) {
	session = getSession()
	err = session.DB(DB_Table).C(DB_COL).Insert(docs...)
	return err, session
}

// func Insert(docs ...interface{}) error {
// 	session, err := mgo.Dial(DB_Server)
// 	CheckError(err)
// 	defer session.Close()
// 	// 切换场景到一个monotonic
// 	session.SetMode(mgo.Monotonic, true)
// 	c := session.DB(DB_Table).C(DB_COL)
// 	return c.Insert(docs...) // 数组传入的方式需要带...
// }
