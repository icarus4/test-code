package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	// "reflect"
	"time"
)

type Person struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Phone string
	Timestamp time.Time
}

func f(from string) {
    for i := 0; i < 10; i++ {
        fmt.Println(from, ":", i)
    }
}



func main() {
	session, err := mgo.Dial("192.168.99.100:27017")
	// session, err := mgo.Dial("0.0.0.0:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")

	// err = c.Insert(&Person{"CC", "+55 53 8116 9639"},
	// 	&Person{"AA", "+55 53 8402 8510"},
	// 	&Person{"BB", "17 Media"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Query one
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)

	// Query all
	var results []Person
	err = c.Find(bson.M{"name": "Cla"}).All(&results)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Results All: ", results)

	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

