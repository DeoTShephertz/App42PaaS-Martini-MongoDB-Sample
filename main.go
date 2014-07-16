package main

import (
	"fmt"
	"github.com/go-martini/martini" // loading in the Martini package
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"      // loading in the mgo MongoBD driver package
	"labix.org/v2/mgo/bson" // loading in the bson package
	"net/http"
)

// Create a struct which matches the BSON documents in the database collection you want to access
type User struct {
	Name  string "name"
	Email string "email"
	Desc  string "desc"
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	databaseName = "DATABASENAME" // "DATABASE NAME"
	collection   = "users"
)

func main() {
	// if you are new to Go the := is a short variable declaration
	m := martini.Classic()

	// reads "templates" directory by default
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	// the func() call is creating an anonymous function that retuns a string
	m.Post("/users", func(ren render.Render, r *http.Request) {

		name := r.FormValue("name")
		email := r.FormValue("email")
		description := r.FormValue("description")

		//mgoSession, err := mgo.Dial("localhost")
		mgoSession, err := mgo.Dial("USER:PASSWORD@VM IP:VM PORT/DATABASE NAME")
		c := mgoSession.DB(databaseName).C(collection)

		err = c.Insert(&User{name, email, description})

		if err != nil {
			panic(err)
		}

		users := []User{}
		err = c.Find(bson.M{}).All(&users)
		PanicIf(err)

		fmt.Println(users)

		ren.HTML(200, "users", users)
	})
	m.Run()
}
