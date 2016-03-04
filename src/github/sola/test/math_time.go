package test

import (
	"fmt"
	"github/sola/db"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
)

type ColResult struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	X        int
	Time     time.Time `bson:"timestamp"`
	TimeUnit string
	Name     string `bson:"type"`
}

func DB_Mongo_Run() {
	result := ColResult{
		X: rand.Intn(2000), Name: "insert_type", Time: time.Now(),
		TimeUnit: time.Now().Format("2016-01-02 15:04:05"),
	}
	err := db.Insert(result)
	CheckError(err)

	list := []ColResult{}
	query, session := db.Find(nil)
	defer session.Close()
	query.All(&list)
	fmt.Println(list)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("test:", err)
	}
}
