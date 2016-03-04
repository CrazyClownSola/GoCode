package main

import (
	"fmt"
	// "github/sola/net"
	// "github/sola/test"
	// "github/sola/db"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "github/sola/repository"
)

func main() {
	fmt.Println("Main Start")
	// net.Init_Router().RunOnAddr(":51000")
	// mgo.GridFile
	// test.DB_Mongo_Run()
	// Compose(sin, cos)
}

// func Compose(f, g func(x float) float) func(x float) float {
// 	return func(x float) float {
// 		return f(g(x))
// 	}
// }

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
