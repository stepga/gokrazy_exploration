package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

const config = "/perm/foo/config.json"

//                    ^ this directory has to exist, otherwise this will panic

type Config struct {
	FilesToCreate int `json:"files_to_create"`
}

func main() {
	_, err := os.Stat(config)
	if err != nil {
		fsErr := &fs.ErrNotExist
		if !errors.As(err, fsErr) {
			panic(err)
		}

		content, err := json.Marshal(&Config{FilesToCreate: 4})
		if err != nil {
			panic(err)
		}

		if err = ioutil.WriteFile(config, content, 0600); err != nil {
			panic(err)
		}
	}

	content, err := ioutil.ReadFile(config)
	if err != nil {
		panic(err)
	}

	var c Config
	if err := json.Unmarshal(content, &c); err != nil {
		panic(err)
	}

	for i := 0; i < c.FilesToCreate; i++ {
		if err := ioutil.WriteFile(fmt.Sprintf("/tmp/%d.txt", i), []byte("gokrazy rocks"), 0600); err != nil {
			panic(err)
		}
	}
}
