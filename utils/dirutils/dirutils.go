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
		log.Print(err)
	}
	for _, f := range files {
		filenames = append(filenames, prefix+dirName+f.Name())
	}
	return filenames
}

//ReadImagesDir reads the folder and return those images that match the typeOfImage provided
func ReadImagesDir(directoryName string, externalDevicePath string, typeOfImage string) (filenames []string) {
	dirName := "imgs/" + directoryName

	//read local files
	filenames = readFilesFromDir(dirName, "")

	//read external (usb) files
	existsExternal, err := fileutils.Exists(externalDevicePath)
	if err != nil {
		log.Print("While checking for external devices:", err)
	}
	if existsExternal {
		externalFilenames := readFilesFromDir(externalDevicePath+dirName, "EXT")
		filenames = append(filenames, externalFilenames...)
	}

	//filter the filenames by tag
	filenames = listutils.FilterByTag(filenames, typeOfImage)

	//return randomized filenames order
	return listutils.RandomList(filenames)
}
