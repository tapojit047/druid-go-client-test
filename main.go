package main

import (
	"fmt"
	"github.com/grafadruid/go-druid"
	"io/ioutil"
	"log"
)

func main() {
	var druidOpts []druid.ClientOption
	druidOpts = append(druidOpts, druid.WithBasicAuth("admin", "7;Nnsc)MqZNMWhYv"))
	d, err := druid.NewClient("http://localhost:8088", druidOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := d.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	DruidHealthCheck(d)
	DruidCreateDataSource(d)
}

func DruidHealthCheck(d *druid.Client) {
	health, _, err := d.Common().Health()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Health: ", *health)
}

func DruidCreateDataSource(d *druid.Client) {
	method := "POST"
	path := "/druid/indexer/v1/task"

	specContent, err := ioutil.ReadFile("spec.json")
	if err != nil {
		log.Fatal(err)
	}

	var s *druid.Status
	res, err := d.ExecuteRequest(method, path, specContent, &s)
	fmt.Println(res)
}
