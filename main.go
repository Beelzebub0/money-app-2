package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id        int
	Name      string
	Job       string
	Notes     string
	Status    int
	Flag      int
	CreatedAt string
	UpdatedAt string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "latihan"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM user ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	usr := User{}
	res := []User{}
	for selDB.Next() {
		var id, status, flag int
		var name, job, notes, created_at, updated_at string
		err = selDB.Scan(&id, &name, &job, &notes, &status, &flag, &created_at, &updated_at)
		if err != nil {
			panic(err.Error())
		}
		usr.Id = id
		usr.Name = name
		usr.Job = job
		usr.Notes = notes
		usr.Status = status
		usr.Flag = flag
		usr.CreatedAt = created_at
		usr.UpdatedAt = updated_at
		res = append(res, usr)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM user WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := User{}
	for selDB.Next() {
		var id, status, flag int
		var name, job, notes, created_at, updated_at string
		err = selDB.Scan(&id, &name, &job, &notes, &status, &flag, &created_at, &updated_at)
		if err != nil {
			panic(err.Error())
		}
		usr.Id = id
		usr.Name = name
		usr.Job = job
		usr.Notes = notes
		usr.Status = status
		usr.Flag = flag
		usr.CreatedAt = created_at
		usr.UpdatedAt = updated_at
	}
	tmpl.ExecuteTemplate(w, "Show", usr)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM user WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := User{}
	for selDB.Next() {
		var id int
		var name, job, notes string
		err = selDB.Scan(&id, &name, &job, &notes)
		if err != nil {
			panic(err.Error())
		}
		usr.Id = id
		usr.Name = name
		usr.Job = job
		usr.Notes = notes
	}
	tmpl.ExecuteTemplate(w, "Edit", usr)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		job := r.FormValue("job")
		notes := r.FormValue("notes")
		insForm, err := db.Prepare("INSERT INTO user(name, job, notes) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, job, notes)
		log.Println("INSERT: Name: " + name + " | Job: " + job + " | Notes: " + notes)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		job := r.FormValue("job")
		notes := r.FormValue("notes")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE user SET name=?, job=?, notes=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, job, notes, id)
		log.Println("UPDATE: Name: " + name + " | Job: " + job + " | Notes: " + notes)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	usr := r.URL.Query().Get("id")
	delForm, err := db.Prepare("UPDATE user SET status=0 WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(usr)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
