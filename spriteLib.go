package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type AnimationData struct {
	X          int
	Y          int
	W          int
	H          int
	Frame      int
	Speed      int
	Path       string
	Horizontal bool
}

type SpriteLib map[string]AnimationData

func loadSpriteLibrary(path string) SpriteLib {
	content, err := ioutil.ReadFile(path)

	var data map[string]AnimationData

	if err != nil {
		log.Fatal("SpriteLib - Error when opening json file: ", err)
	}

	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Fatal("SpriteLib - Error during Unmarshal(): ", err)
	}

	return data
}
