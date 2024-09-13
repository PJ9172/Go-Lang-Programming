package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	Name       string `json:"fullname"`
	Address    string `json:"address"`
	Contact_No string `json:"contactno"`
	Email      string `json:"email"`
	DOB        string `json:"dob"`
	Gender     string `json:"gender"`
}

var student Student

func main() {

	// Handle the root URL to serve the HTML file
	http.HandleFunc("/", serveHTML)

	// Handle the form submission
	http.HandleFunc("/submit", formDataHandler)

	// Server starting on port 3000
	fmt.Println("Server starting on port 3000!!!")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error in server starting", err)
		return
	}
}

// Serve the HTML file
func serveHTML(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

// Handle form submission
func formDataHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()
		student.Name = req.FormValue("fullname")
		student.Address = req.FormValue("address")
		student.Contact_No = req.FormValue("contactnumber")
		student.Email = req.FormValue("mail")
		student.DOB = req.FormValue("dob")
		student.Gender = req.FormValue("gender")

		studentJson, err := json.MarshalIndent(student, "", "\t")
		if err != nil {
			fmt.Fprint(res, http.StatusBadGateway)
			return
		}
		// fmt.Fprint(res, "Student data in JSON : ", string(studentJson))

		// Parse and render the studentData.html template
		tmpl, err := template.ParseFiles("studentData.html")
		if err != nil {
			http.Error(res, "Error loading template", http.StatusInternalServerError)
			return
		}

		// Pass the JSON data as a string to the template
		tmpl.Execute(res, template.JS(studentJson))
	}
}
