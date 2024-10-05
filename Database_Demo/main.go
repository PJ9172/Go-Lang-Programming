package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/submit", submitfunc)
	http.HandleFunc("/", rootUrl)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server Starts on Port 3000!!!")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}

func rootUrl(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func submitfunc(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()
		name := req.FormValue("name")
		dob := req.FormValue("dob")
		email := req.FormValue("email")

		dbDriver := "mysql"
		dbUser := "root"
		dbPass := "root"
		dbName := "College"

		db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
		if err != nil {
			fmt.Println("Error in Database connection :\n", err)
			http.Error(res, "Internal server error", http.StatusInternalServerError) // Return HTTP 500 error to the client
			return
		}

		insert, err := db.Prepare("insert into students(name,dob,email) values (?,?,?)")
		if err != nil {
			fmt.Println("Error to prepare insert query : \n", err)
			http.Error(res, "Internal server error", http.StatusInternalServerError) // Return HTTP 500 error to the client
			return
		}

		_, err = insert.Exec(name, dob, email)
		if err != nil {
			fmt.Println("Error to execute insert query : \n", err)
		}
		defer db.Close()
		fmt.Fprint(res, "Data inserted into Database successfully!!!")
	}
}

// func connect(