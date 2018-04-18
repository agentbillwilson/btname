package main

import (
	"fmt"
	"log"
	"os"

	bencode "github.com/wwalexander/go-bencode"
)

// NameFromFile decodes the value of the name string from the info dictionary
// of a BitTorrent metainfo file.
func NameFromFile(file *os.File) (string, error) {
	var metainfo struct {
		Info struct {
			Name string `bencode:"name"`
		} `bencode:"info"`
	}
	if err := bencode.NewDecoder(file).Decode(&metainfo); err != nil {
		return "", err
	}
	return metainfo.Info.Name, nil
}

// NameFromFilename decodes the value of the name string from the info
// dictonary of a BitTorrent metainfo file named name.
func NameFromFilename(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	metaname, err := NameFromFile(file)
	if err != nil {
		return "", err
	}
	return metaname, nil
}

func main() {
	if len(os.Args) == 1 {
		name, err := NameFromFile(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
		os.Exit(0)
	}
	for _, arg := range os.Args[1:] {
		name, err := NameFromFilename(arg)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(name)
	}
}
