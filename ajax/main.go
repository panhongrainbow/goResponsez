package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type PageData struct {
	TableData [10][10]int
}

type Data struct {
	Number int
	Reload bool
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	//
	var d Data

	// Create a new source with the current time
	source := rand.NewSource(time.Now().UnixNano())

	// Create a new random number generator with the new source
	d.Number = rand.New(source).Intn(100)
	if d.Number >= 50 {
		d.Reload = true
	}

	// Create json response from struct
	response, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(string(response))

	//
	_, _ = w.Write(response)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
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
	tmpl, err := template.ParseFiles("./templates/ajax.html")
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
}

func main() {
	http.HandleFunc("/ajax", ajaxHandler)
	http.HandleFunc("/", defaultHandler)
	_ = http.ListenAndServe(":8080", nil)
}
