package dirutils

import (
	"io/ioutil"
	"log"

	"github.com/aries/utils/fileutils"
	"github.com/aries/utils/listutils"
)

func readFilesFromDir(dirName string, prefix string) (filenames []string) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		//log.Print(err)
	}
	for _, f := range files {
		filenames = append(filenames, prefix+dirName+f.Name())
	}
	return filenames
}

//ReadImagesDir reads the folder and return those images that match the typeOfImage provided
func ReadImagesDir(directoryName string, mountLocation string, typeOfImage string) (filenames []string) {
	dirName := "imgs/" + directoryName

	//read local files
	filenames = readFilesFromDir(dirName, "")

	//read external (usb) files
	//discover read locations
	externalDirs := discoverExternalDrives(mountLocation, dirName)
	for _, dir := range externalDirs {
		externalFilenames := readFilesFromDir(dir+dirName, "EXT")
		filenames = append(filenames, externalFilenames...)
	}

	//filter the filenames by tag
	filenames = listutils.FilterByTag(filenames, typeOfImage)

	//return randomized filenames order
	return listutils.RandomList(filenames)
}

func discoverExternalDrives(mountLocation string, imagesDir string) (externalDirs []string) {
	//read mount location, read all dirs, then for each file check if image directory exists
	externalDevices := readFilesFromDir(mountLocation, "")
	for _, externalDevice := range externalDevices {
		//check if imagesDir exists on each external
		externalImages, err := fileutils.Exists(externalDevice + "/" + imagesDir)
		if err != nil {
			log.Printf("Error while reading dir: %v", err)
		}
		if externalImages {
			externalDirs = append(externalDirs, externalDevice+"/")
		}
	}
	return externalDirs
}
