package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var session ServerVars

type ServerVars struct {
	ImagesPortrait []string
	ImagesLanscape []string
	ImageBatch     int
}

type Page struct {
	Body         []byte
	ImageRefresh int
	PageRefresh  int
	Images       []string
	TypePage     string
	TypeOfImage  string
}

type Config struct {
	NextImage   bool
	TypeOfImage string
	Images      []string
}

func loadPage(title string) (*Page, error) {
	filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Body: body, ImageRefresh: 600000}, nil
}

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
	log.Print("Serving: " + typePage + ", Query: " + query)
	p, _ := loadPage(query)

	if strings.Contains(query, "aries.htm") {
		p.Images = readImagesDir(typePage, readConfig(typePage))
		if len(p.Images) == 0 {
			p.Images = readImagesDir(typePage, Config{TypeOfImage: "all"})
		}
		p.TypePage = typePage
		p.TypeOfImage = readConfig(typePage).TypeOfImage
		p.PageRefresh = (p.ImageRefresh / 1000) * len(p.Images)
		t, _ := template.ParseFiles(query)
		t.Execute(w, p)
	} else if strings.Contains(query, "configs") {
		fmt.Fprintf(w, "%s", getConfig(typePage))
	} else if strings.Contains(query, "updateConfig") {
		if err := r.ParseForm(); err != nil {
			log.Fatal(w, "ParseForm() err: %v", err)
			return
		}

		nextImage := r.FormValue("nextImage")
		typeImage := r.FormValue("typeImage")
		fmt.Fprintf(w, "%s", updateConfig(typePage, nextImage, typeImage))

	} else {
		if p != nil {
			fmt.Fprintf(w, "%s", p.Body)
		} else {
			fmt.Fprintf(w, "%s", "404: Not found")
		}
	}
}

func updateConfig(typePage string, nextImage string, typeImage string) (jsonConfig []byte) {
	var currentConfig = readConfig(typePage)
	nextImagetmp, err := strconv.ParseBool(nextImage)
	if err != nil {
		log.Print("No valid value passed for nextImage, using current one")
		nextImagetmp = currentConfig.NextImage
	}
	if nextImagetmp {
		nextImagetmp = !currentConfig.NextImage
	}
	if typeImage == "" {
		typeImage = currentConfig.TypeOfImage
	}

	var newConfig = Config{TypeOfImage: typeImage, NextImage: nextImagetmp}
	writeConfigFile(newConfig, typePage)

	jsonConfig, erre := json.Marshal(newConfig)
	if erre != nil {
		log.Fatal(erre)
	}

	return jsonConfig

}

func writeConfigFile(config Config, typeConfig string) {
	filename := "config/" + typeConfig[0:len(typeConfig)-1] + ".config"
	configJSON, _ := json.Marshal(config)
	err := ioutil.WriteFile(filename, configJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getConfig(typeConfig string) (jsonConfig []byte) {
	var config = readConfig(typeConfig)

	//get Images by tags on the config
	config.Images = readImagesDir(typeConfig, config)
	if len(config.Images) == 0 {
		config.Images = readImagesDir(typeConfig, Config{TypeOfImage: "all"})
		config.TypeOfImage = "all"
	}
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}

	return jsonConfig
}

func readConfig(typeConfig string) (config Config) {
	filename := "config/" + typeConfig[0:len(typeConfig)-1] + ".config"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(body), &config)

	return
}

func main() {

	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readImagesDir(directoryName string, config Config) (filenames []string) {
	dirname := "imgs/" + directoryName
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if config.TypeOfImage == "all" {
			filenames = append(filenames, dirname+f.Name())

		} else {
			//split image by -
			//filter by tag
			tags := strings.Split(f.Name(), "-")
			for _, tag := range tags {
				if strings.Contains(tag, config.TypeOfImage) {
					filenames = append(filenames, dirname+f.Name())
				}
			}
		}
	}

	for i := range filenames {
		j := rand.Intn(i + 1)
		filenames[i], filenames[j] = filenames[j], filenames[i]
	}
	return filenames
}
