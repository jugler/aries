package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aries/models"
	"github.com/aries/utils/dirutils"
	"github.com/aries/utils/fileutils"
)

var session models.ServerVars

var mountLocation = "/media/pi/"

//var mountLocation = "/Users/jugler/code/ariesImgs/"

func viewHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[len("/"):]
	s := strings.SplitAfterN(url, "/", 2)
	var typePage string
	var query string
	if len(s) >= 2 {
		typePage, query = s[0], s[1]
	} else {
		typePage = s[0]

	}
	//log.Print("Serving: " + typePage + ", Query: " + query)
	p, _ := fileutils.LoadPage(query)

	if strings.Contains(query, "aries.htm") {
		p.Images = dirutils.ReadImagesDir(typePage, mountLocation, fileutils.ReadConfig(typePage[0:len(typePage)-1]).TypeOfImage)
		if len(p.Images) == 0 {
			p.Images = dirutils.ReadImagesDir(typePage, mountLocation, models.Config{TypeOfImage: "all"}.TypeOfImage)
		}
		p.TypePage = typePage
		p.TypeOfImage = fileutils.ReadConfig(typePage[0 : len(typePage)-1]).TypeOfImage
		p.PageRefresh = (p.ImageRefresh / 1000) * len(p.Images)
		t, _ := template.ParseFiles(query)
		t.Execute(w, p)
	} else if strings.Contains(query, "configs") {
		fmt.Fprintf(w, "%s", getConfig(typePage))
	} else if strings.Contains(query, "updateConfig") {
		if err := r.ParseForm(); err != nil {
			log.Print(w, "ParseForm() err: %v", err)
			return
		}

		nextImage := r.FormValue("nextImage")
		typeImage := r.FormValue("typeImage")
		fmt.Fprintf(w, "%s", fileutils.UpdateConfig(typePage[0:len(typePage)-1], nextImage, typeImage))

	} else {
		if p != nil {
			fmt.Fprintf(w, "%s", p.Body)
		} else {
			log.Print("404: ", query)
			fmt.Fprintf(w, "%s", "404: Not found")
		}
	}
}

var imagesPortrait = dirutils.ReadImagesDir("portrait", mountLocation, "all")
var imagesLandscape = dirutils.ReadImagesDir("landscape", mountLocation, "all")
var typeOfImage = "all"

func getConfig(typeConfig string) (jsonConfig []byte) {
	typeConfig = typeConfig[0 : len(typeConfig)-1]
	var config = fileutils.ReadConfig(typeConfig)

	//get Images by tags on the config
	if config.TypeOfImage != typeOfImage {
		log.Printf("Config %s not the same as global %s ", config.TypeOfImage, typeOfImage)
		imagesPortrait = dirutils.ReadImagesDir("portrait/", mountLocation, config.TypeOfImage)
		imagesLandscape = dirutils.ReadImagesDir("landscape/", mountLocation, config.TypeOfImage)
		config.Images = imagesLandscape
		if typeConfig == "portrait" {
			config.Images = imagesPortrait
		}
		typeOfImage = config.TypeOfImage
		if len(config.Images) == 0 {
			log.Printf("Didnt found images for %s searching all ", config.TypeOfImage)
			imagesPortrait = dirutils.ReadImagesDir("portrait/", mountLocation, models.Config{TypeOfImage: "all"}.TypeOfImage)
			imagesLandscape = dirutils.ReadImagesDir("landscape/", mountLocation, models.Config{TypeOfImage: "all"}.TypeOfImage)
			config.Images = imagesLandscape
			if typeConfig == "portrait" {
				config.Images = imagesPortrait
			}
			config.TypeOfImage = "all"
		}
	} else {
		config.Images = imagesLandscape
		if typeConfig == "portrait" {
			config.Images = imagesPortrait
		}
	}
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		log.Print(err)
	}
	if config.Reload == true {
		config.Reload = false
		fileutils.WriteConfigFile(config, typeConfig)
	}
	return jsonConfig
}

func main() {
	f, err := os.OpenFile("aries.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Print("Starting aries")

	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
