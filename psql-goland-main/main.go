package main

import (

	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)


var tmpl = template.Must(template.ParseGlob("vista"))

func Index(w http.ResponseWriter, r http.Request) {
	db := getConnection()
	selDB, err := db.Query("SELECT * FROM ebook ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

}
  ebo:= ebook{}
res := []ebook{}
for selDB.Next(){
	var id int
	var name string
	var date time.Time
	var study bool
	err= selDB.Scan(&id, &name, &date, &study)
	if err!= nil{
		panic(err.Error())
}
ebo.id = id
ebo.name = name
ebo.date = date
ebo.study = study
res = append(res, ebo)
}
tmpl.ExecuteTemplate(w, "index", res)
defer db.Close()
}


func Create() {
	e := &ebook{
		name:  "The little fox",
		date:  time.Now(),
		study: false,
	}
	err := CreateEbook(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created! :)")
}



func Show(w http.ResponseWriter, r http.Request)  {
	db := getConnection()
	nid:= r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM ebook WHERE id=$" , nid)
	if err!=nil{
		panic(err.Error())
	}
	tmpl.ExecuteTemplate(w, "Show", ebook{})
	defer db.Close()
}

func main() {
	http.HandlerFunc("/", Index)
	http.HandlerFunc("/", Show)
	http.HandlerFunc("/", CreateEbook)
	http.HandlerFunc("/", Create)
	log.Fatal(http.ListenAndServe("3000", nil))

}
