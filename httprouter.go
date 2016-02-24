package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"models"
	"net/http"
)

type Person struct {
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Phone string
}

type Env struct {
	db *mgo.Session
}

func (env *Env) Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := env.db.DB("test").C("people")

	// var results []Person
	// err := c.Find(bson.M{"name": "Cla"}).All(&results)
	// if err != nil {
	// 	panic(err)
	// }

	result := Person{}
	_ = c.Find(bson.M{"name": "Ale"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	u := models.User{
		ID: result.ID,
		Name: result.Name,
		Phone: result.Phone,
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", uj)
}

func (env *Env) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c := env.db.DB("test").C("people")
	var results := 
}

func main() {
	db, err := mgo.Dial("192.168.99.100:27017")
	db.SetMode(mgo.Monotonic, true)

	if err != nil {
		panic(err)
	}

	env := &Env{db: db}

	router := httprouter.New()
	router.GET("/", env.Get)

	log.Fatal(http.ListenAndServe(":8080", router))
	env.db.Close()
}

