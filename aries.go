package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
	MyVals string
}



func loadPage(title string) (*Page, error) {
	filename := title 
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	p.MyVals = readDir()
	if strings.Contains(title, "htm"){
		t, _ := template.ParseFiles(title)
    	t.Execute(w, p)
	}else{
		fmt.Fprintf(w, "%s", p.Body)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readDir()(filesString string){
	files, err := ioutil.ReadDir("imgs/")
    if err != nil {
        log.Fatal(err)
    }
	filenames := make([]string, 0)
	for _, f := range files {
		filenames=append(filenames, f.Name())
	}
	filesString = strings.Join(filenames,",")
	return
}
