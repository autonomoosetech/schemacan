package main

import (
	"github.com/autonomoosetech/schemacan/api/v1"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func objectsFromFilenames(files []string) (objects []api.Object, err error) {
	log.Println("Parsing in:")
	for _, file := range files {
		log.Println("\t" + file)

		dat, err := os.ReadFile(file)
		if err != nil {
			log.Fatalln(err)
		}

		var obj api.Object
		err = yaml.Unmarshal(dat, &obj)
		if err != nil {
			break
		}

		objects = append(objects, obj)
	}

	return
}
