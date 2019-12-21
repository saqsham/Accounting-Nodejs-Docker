package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "strings"
	"crypto/md5"
	"time"
	"io"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type logins struct {
	id 			int
	username 	string
	email 		string
	password 	string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
    dbUser := "saqsham"
    dbPass := "1234"
    dbName := "golang"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

// var t = template.Must(template.ParseGlob("templates/*"))

// func assets(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, r.URL.Path[1:])
// }

func sayhelloname(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parse url params passed
	// fmt.Println(r.Form) // print info on server side
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }
	fmt.Fprintf(w, "hewo error here") // writing data to response
}

func signup(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("method:", r.Method) // get request method
	if r.Method == "GET" {
		// MD5 hash
		crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
	 	token := fmt.Sprintf("%x", h.Sum(nil))
		
		t, _ := template.ParseFiles("templates/signup.html")
		t.Execute(w, token)

		// fmt.Println("token:", token)

	} else {
		// POST data
		r.ParseForm() 
		// fmt.Println("thisexecutes")

		token := r.Form.Get("token")
        if token != "" {
            // check token validity
        } else {
            // give error if no token
        }
		// logic part of login
		// fmt.Println("email:", template.HTMLEscapeString(r.Form.Get("email")))
		// fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("2: username"))) // print at server side
		// fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		// template.HTMLEscape(w, []byte(r.Form.Get("username"))) // responded to clients
		
		// https://golang.org/pkg/html/template/
		// t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	    // err = t.ExecuteTemplate(, "T", template.HTML("<script>alert('you have been pwned')</script>"))
	
		// validation
		username := r.Form["username"]
		if len(username) == 0 {
			t, _ := template.ParseFiles("templates/signup.html")
			t.Execute(w, nil)
		}

		// db work
		db := dbConn()
		defer db.Close()
		if r.Method == "POST" {
			username := r.FormValue("username")
			// fmt.Println(username)
			email := r.FormValue("email")
			password := r.FormValue("password")
			insForm, err := db.Prepare("INSERT INTO logins (username, email, password) VALUES (?,?,?)")
        	if err != nil {
            	panic(err.Error())
			}
			// defer insForm
			insForm.Exec(username, email, password)
			// log.Println("INSERT: userame: " + username + " | email: " + email )   
			// fmt.Println("1 row inserted")
			http.Redirect(w, r, "/login", 301)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("method:", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html")
		t.Execute(w, nil)
	} else {
		// db work
		r.ParseForm()
		username := r.FormValue("username")
		password := r.FormValue("password")
		db :=  dbConn()
		defer db.Close()
		selDB, err := db.Query("SELECT password FROM logins WHERE username=?", username)
		if err != nil {
			panic(err.Error())
			// log.Fatal(err)
		}
		logingo := logins{}
		for selDB.Next() {
			var password1 string
			err = selDB.Scan(&password1)
			if (err!= nil) {
				panic(err.Error())
			}
			logingo.password = password1
			if (logingo.password != password) {
				http.Redirect(w, r, "/login", 301)
				break
			} else {
				http.Redirect(w, r, "/home", 301)
				break
			}
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
  	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", sayhelloname)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	// http.HandleFunc(".*.[js|css]", assets) // static files 
	err := http.ListenAndServe(":3000", nil) // port setup
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}