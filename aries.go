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
	ImageList string
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
	p.ImageList = readDir("imgs/")
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

func readDir(directoryName string)(filesString string){
	files, err := ioutil.ReadDir(directoryName)
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
