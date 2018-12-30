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
	Body  []byte
	ImageList string
	ImageRefresh int
}


func loadPage(title string) (*Page, error) {
	filename := title 
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Body: body, ImageRefresh: 60000}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[len("/"):]
	s := strings.SplitAfterN(url, "/", 2)
	typePage, query := s[0], s[1]
	log.Print("Serving: " + typePage)
	log.Print("Query: " +  query)
	p, _ := loadPage(query)
	if strings.Contains(query, ".htm"){
		p.ImageList = readImagesDir(typePage)
		t, _ := template.ParseFiles(query)
    	t.Execute(w, p)
	}else{
		fmt.Fprintf(w, "%s", p.Body)
	}
}


func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readImagesDir(directoryName string)(filesString string){
	dirname := "imgs/" + directoryName
	files, err := ioutil.ReadDir(dirname)
    if err != nil {
        log.Fatal(err)
	}
	filenames := make([]string, 0)
	for _, f := range files {
		log.Print(dirname+f.Name())
		filenames=append(filenames, dirname + f.Name())
	}
	filesString = strings.Join(filenames,",")
	return
}
