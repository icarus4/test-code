package main

import (
	// "context"
	"encoding/json"
	"fmt"
	mux "github.com/julienschmidt/httprouter"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "log"
	log "github.com/Sirupsen/logrus"
	"io"
	"io/ioutil"
	// "models"
	"net/http"
	"upper.io/db"         // Imports the main db package.
	"upper.io/db/mongo"   // Imports the mongo adapter.
)

var settings = mongo.ConnectionURL{
	Database: `test`,        // Database name.
	Address:   db.Host("192.168.99.100:27017"),  // Host's IP.
}

type Env struct {
	db db.Database
	collection db.Collection
}

type User struct {
	// ID    bson.ObjectId `bson:"_id,omitempty"`
	Name  string
	Phone string
}

func (env *Env) Show(w http.ResponseWriter, r *http.Request, p mux.Params) {
	// var res db.Result
	// res = env.collection.Find()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	j, _ := json.Marshal(p)
	fmt.Fprintf(w, "%s", j)
}

func (env *Env) Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	var res db.Result
	res = env.collection.Find()
	var users []User
	_ = res.All(&users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	// Debug
	log.WithFields(log.Fields{
		"FormValue: ": r.FormValue("q"),
	}).Info("Index")

	// j, _ := json.Marshal(users)
	// fmt.Fprintf(w, "%s", j)
	json.NewEncoder(w).Encode(users)
}

func (env *Env) Create(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	var user User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Body: ", r.Body)

	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422)

        if err := json.NewEncoder(w).Encode(err); err != nil {
        	panic(err)
        }
	}

	// Create Post here
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

	// fmt.Println("Method: ", r.Method)
	// fmt.Println("URL: ", r.URL)
	// fmt.Println("Host: ", r.Host)
	// fmt.Println("Header: ", r.Header)
	// fmt.Println("Body: ", r.Body)
	// fmt.Println("Form: ", r.Form)
	// fmt.Println("PostForm: ", r.PostForm)
	// fmt.Println("MultipartForm: ", r.MultipartForm)
	// fmt.Println("Close: ", r.Close)

	// r.ParseForm()

	log.WithFields(log.Fields{
		"Method: ":			r.Method,
		"URL: ":			r.URL,
		"Host: ": 			r.Host,
		"Header: ": 		r.Header,
		"Body: ": 			r.Body,
		"Form: ": 			r.Form,
		"PostForm: ": 		r.PostForm,
		"MultipartForm: ":	r.MultipartForm,
		"Close: ": 			r.Close,
		"TransferEncoding: ": 			r.TransferEncoding,
		"Trailer: ": 			r.Trailer,
		"RemoteAddr: ": 			r.RemoteAddr,
		"RequestURI: ": 			r.RequestURI,
		"TLS: ": 			r.TLS,
		"FormValue: ": 			r.FormValue("key1"),
	}).Info("Create")

}

func main() {
	sess, err := db.Open(mongo.Adapter, settings)
	c, _ := sess.Collection("people")

	if err != nil {
		panic(err)
	}

	env := &Env{db: sess, collection: c}

	router := mux.New()
	router.GET("/users/", env.Index)
	router.GET("/users/:user_id", env.Show)
	router.POST("/users", env.Create)

	log.Fatal(http.ListenAndServe(":8080", router))
	env.db.Close()
}
