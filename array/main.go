package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// PageData contains a 10x10 matrix.
type PageData struct {
	TableData [10][10]int
}

// the main function generates a 10x10 matrix on the web page.
// Currently, in order to test, it reloads every 10 seconds.
// Later it needs to be corrected.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Initialize tableData with 10x10 array of zeros
		tableData := [10][10]int{}

		// Create a new source with the current time
		source := rand.NewSource(time.Now().UnixNano())

		// Create a new random number generator with the new source
		rand2 := rand.New(source)

		// Fill tableData with random numbers
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				tableData[i][j] = rand2.Intn(100)
			}
		}

		// Create a new PageData struct with the tableData
		pageData := PageData{TableData: tableData}

		// Parse the template
		tmpl, err := template.ParseFiles("./templates/template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template with the page data
		err = tmpl.Execute(w, pageData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Execute the template with the page data
	_ = http.ListenAndServe(":8080", nil)
}
