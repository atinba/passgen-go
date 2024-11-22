package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Struct to hold data for the template
type PageData struct {
	Length   int
	Password string
}

// Function to generate a random password

// HTTP handler function to serve the password generator page
func passwordHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type for HTML response
	w.Header().Set("Content-Type", "text/html")

	// Default password length
	length := 12

	// Parse query parameters to get custom password length
	r.ParseForm()
	if lengthParam := r.FormValue("length"); lengthParam != "" {
		if parsedLength, err := strconv.Atoi(lengthParam); err == nil && parsedLength >= 4 && parsedLength <= 128 {
			length = parsedLength
		}
	}

	// Generate password
	password, _ := generate_password(length)

	// Prepare data for the template
	pageData := PageData{
		Length:   length,
		Password: password,
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func main() {
	// Handle the root URL with the passwordHandler function
	http.HandleFunc("/", passwordHandler)

	// Start the web server on port 8080
	fmt.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
