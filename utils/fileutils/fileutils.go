package fileutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jugler/aries/models"
)

//ReadConfig Reads a config file (portrait/landscape) and returns the Config object
func ReadConfig(typeConfig string) (config models.Config) {
	filename := "config/" + typeConfig + ".config"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
	}
	json.Unmarshal([]byte(body), &config)

	return
}

//UpdateConfig saves a config object to a file
func UpdateConfig(typePage string, nextImage string, typeImage string) (jsonConfig []byte) {
	var currentConfig = ReadConfig(typePage)
	nextImagetmp, err := strconv.ParseBool(nextImage)
	if err != nil {
		nextImagetmp = currentConfig.NextImage
	}
	if nextImagetmp {
		nextImagetmp = !currentConfig.NextImage
	}
	if typeImage == "" {
		typeImage = currentConfig.TypeOfImage
	}

	var newConfig = models.Config{TypeOfImage: typeImage, NextImage: nextImagetmp}
	WriteConfigFile(newConfig, typePage)

	jsonConfig, erre := json.Marshal(newConfig)
	if erre != nil {
		log.Print(erre)
	}

	return jsonConfig

}

//WriteConfigFile serializes a config into a file
func WriteConfigFile(config models.Config, typeConfig string) {
	filename := "config/" + typeConfig + ".config"
	configJSON, _ := json.Marshal(config)
	err := ioutil.WriteFile(filename, configJSON, 0644)
	if err != nil {
		log.Print(err)
	}
}

//LoadPage reads a page and returns it
func LoadPage(title string) (*models.Page, error) {
	if strings.Contains(title, "EXT") {
		title = strings.TrimLeft(title, "EXT")
	}

	filename := title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Body: body, ImageRefresh: 600000}, nil
}

//Exists checks if a file or directory exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
