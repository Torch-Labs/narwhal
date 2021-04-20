package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func CheckConfigExists() bool {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(dir + ".narwhal_config.yml"); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetConfigBytes() []byte {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(dir + ".narwhal_config.yml")
	if err != nil {
		panic(err)
	}
	return data
}

func PrintError(name string) {
	fmt.Printf("Error: %s\n", name)
}
