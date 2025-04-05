package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Info struct {
	RollNo int    `json:"rollno"`
	Name   string `json:"name"`
	DOB    string `json:"dob"`
	Email  string `json:"email"`
}

func connection() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "College"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost)/"+dbName)
	if err != nil {
		fmt.Println("Error in connection!!!", err)
		return
	}
	return db
}

func main() {
	http.HandleFunc("/", rootUrl)
	http.HandleFunc("/delete", deleteStudent)
	http.HandleFunc("/update", update)
	http.HandleFunc("/updateStudent", updateStudent)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server starts on Port 5000!!!")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Error to start server!!!", err)
		return
	}
}

func rootUrl(res http.ResponseWriter, req *http.Request) {
	db := connection()
	defer db.Close()

	// Query to get all students from database
	var students []Info
	rows, err := db.Query("SELECT RollNo, Name, DOB, Email FROM students")
	if err != nil {
		fmt.Println("Error in select query!!!", err)
		return
	}
	defer rows.Close()

	// Loop through results and add to slice
	for rows.Next() {
		var student Info
		if err := rows.Scan(&student.RollNo, &student.Name, &student.DOB, &student.Email); err != nil {
			fmt.Println("Error scanning rows:", err)
			return
		}
		students = append(students, student)
	}

	// Convert to JSON
	studentJson, err := json.Marshal(students)
	if err != nil {
		fmt.Println("Error in JSON marshaling:", err)
		return
	}

	// Parse and inject JSON into template
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}
	tmpl.Execute(res, string(studentJson)) // Pass JSON string to HTML
}

func deleteStudent(res http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		http.Error(res, "Invalid request METHOD", http.StatusMethodNotAllowed)
		return
	}

	rollno, _ := strconv.Atoi(req.URL.Query().Get("rollno"))

	db := connection()
	defer db.Close()

	_, err := db.Exec("Delete from students where rollno= ?", rollno)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error in delete Query!! ", err)
		json.NewEncoder(res).Encode(map[string]bool{"success": false})
		return
	}
	json.NewEncoder(res).Encode(map[string]bool{"success": true})
}

func update(res http.ResponseWriter, req *http.Request) {
	rollno, err := strconv.Atoi(req.URL.Query().Get("rollno"))
	if err != nil {
		fmt.Println("Invalid rollno!!!", err)
		return
	}

	db := connection()
	defer db.Close()

	var student Info
	err = db.QueryRow("select rollno, name, dob, email from students where rollno= ?", rollno).Scan(&student.RollNo, &student.Name, &student.DOB, &student.Email)
	if err != nil {
		fmt.Println("Error in select Query!!!", err)
		return
	}

	temp, _ := template.ParseFiles("update.html")
	temp.Execute(res, student)
}

func updateStudent(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		fmt.Println("Envalid method called!!!")
		return
	}

	req.ParseForm()
	rollno := req.FormValue("rollno")
	name := req.FormValue("name")
	dob := req.FormValue("dob")
	email := req.FormValue("email")

	db := connection()
	defer db.Close()

	_, err := db.Exec("UPDATE students SET Name = ?, DOB = ?, Email = ? WHERE RollNo = ?", name, dob, email, rollno)
	if err != nil {
		fmt.Println("Error in update query!!", err)
		return
	}

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
