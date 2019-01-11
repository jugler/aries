package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
	"strings"
	"math/rand"

)
var session ServerVars

type ServerVars struct{
	ImagesPortrait[] string
	ImagesLanscape[] string
	ImageBatch int
}

type Page struct {
	Body  []byte
	ImageRefresh int
	PageRefresh int
	Images []string
	TypePage string
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
		p.Images = readImagesDir(typePage)
		p.TypePage = typePage
		p.PageRefresh = (p.ImageRefresh / 1000) * len(p.Images)
		t, _ := template.ParseFiles(query)
    	t.Execute(w, p)
	}else{
		fmt.Fprintf(w, "%s", p.Body)
	}
}



func main() {
	session := loadServer()
	log.Print(session.ImageBatch)
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadServer()(*ServerVars){
	return &ServerVars{ImagesPortrait: readImagesDir("portrait"), ImagesLanscape: readImagesDir("landscape"), ImageBatch:5}
}

func readImagesDir(directoryName string)(filenames []string){
	dirname := "imgs/" + directoryName
	files, err := ioutil.ReadDir(dirname)
    if err != nil {
        log.Fatal(err)
	}
	for _, f := range files {
		log.Print(dirname+f.Name())
		filenames=append(filenames, dirname + f.Name())
	}
	
	for i := range filenames {
		j := rand.Intn(i + 1)
		filenames[i], filenames[j] = filenames[j], filenames[i]
	}
	return filenames[0:10]
}
