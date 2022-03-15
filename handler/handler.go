package handler

import (
	"as-golangweb/entity"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"title": "Home Page",
		"content": "Belajar Golang untuk Web Development bersama Imam Rizaldi" ,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "There is an error:(", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)

	// fmt.Fprintln(w, "<h1>Welcome Home</h1>")

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World</h1>")
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hi from Mario</h1>")
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	// get id query
	// id := r.URL.Query().Get("id")

	// convert to int
	// idNumb, err := strconv.Atoi(id)
	// if err != nil || idNumb < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }

	data := []entity.Product{
		{ID: 1, Name: "Laptop ASUS", Price: 10000000, Stock: 5},
		{ID: 2, Name: "Laptop ACER", Price: 11000000, Stock: 8},
		{ID: 3, Name: "iPhone X", Price: 15000000, Stock: 2},
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "There is an error", http.StatusInternalServerError)
		return
	}

}

func FormHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "There is an error", http.StatusBadRequest)
			return
		}

		if err != nil {
			log.Println(err)
			http.Error(w, "There is an error", http.StatusBadRequest)
			return
		}

		// message := r.Form.Get("message")

		tmpl.Execute(w, nil)

	}

}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "There is an error", http.StatusBadRequest)
			return
		}
		name := r.Form.Get("name")
		fmt.Fprintf(w, "%s", name)
		return
	}

	http.Error(w, "There is an error", http.StatusBadRequest)
}