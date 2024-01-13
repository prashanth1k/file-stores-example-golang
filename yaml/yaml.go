package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Request struct {
	Body    map[string]interface{} `yaml:"body"`
	Headers map[string]string      `yaml:"headers"`
}

type Response struct {
	Body    []ResponseBody    `yaml:"body"`
	Headers map[string]string `yaml:"headers"`
}

type ResponseBody struct {
	UserId int    `yaml:"userId"`
	ID     int    `yaml:"id"`
	Title  string `yaml:"title"`
	Body   string `yaml:"body"`
}

type API struct {
	Name   string   `yaml:"name"`
	URL    string   `yaml:"url"`
	Method string   `yaml:"method"`
	Req    Request  `yaml:"req"`
	Res    Response `yaml:"res"`
}

func main() {
	err := filepath.Walk("./data/yaml/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".yaml" {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				log.Fatal(err)
			}

			var api API
			err = yaml.Unmarshal(content, &api)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("API struct of %s: %+v\n", path, api)
			fmt.Printf("URL / Method: %s, %+s\n", api.URL, api.Method)
			fmt.Println("=====================================")
			fmt.Printf("Response Body: %s \n", api.Res.Body[0].Body)
			fmt.Println("=====================================")

		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
