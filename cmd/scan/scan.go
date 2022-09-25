package main

import (
	_ "flag"
	_ "fmt"
	_ "io/ioutil"
	_ "os"
	_ "time"

	_ "github.com/autonomoosetech/schemacan/api/v1"
	_ "github.com/grafana/loki-client-go/loki"
	_ "github.com/prometheus/common/model"
)

func main() {
	return
}

/*
func getObjectsFromFile(path string) (object []api.Resource, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	obj, err := Unmarshal(file)
	if err != nil {
		return nil, err
	}

	return obj, err
}

func main() {
	schemaFile := flag.String("file", "schemacan.yaml", "file containing schemacan definitions")

	lokiUrl := flag.String("lokiurl", "http://localhost:3100", "loki server base path")
	lokiPath := flag.String("lokipath", "/loki/api/v1/push", "loki server url path")

	flag.Parse()

	// read in schemacan yaml
	object, err := getObjectsFromFile(*schemaFile)
	if err != nil {
		fmt.Printf("failed ot read schemacan file: %v", err)
		os.Exit(1)
	}

	_ = object

	// set up loki client
	c, err := loki.NewWithDefault(*lokiUrl + *lokiPath)
	if err != nil {
		fmt.Println("failed there")
		return
	}

	err = c.Handle(model.LabelSet{
		"name":      "gamma",
		"namespace": "default",
	}, time.Now(),
		"id_extended=1365 data_enabled=1 data_voltage_ok=0 voltage=1.5")

	if err != nil {
		fmt.Println("failed here")
		return
	}

	exit := make(chan string)

	// Spawn all you worker goroutines, and send a message to exit when you're done.
	for {
		select {
		case <-exit:
			c.Stop()

			os.Exit(0)
		}
	}
}
*/
