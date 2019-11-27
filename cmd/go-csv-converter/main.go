package main

import (
	"encoding/csv"
	"go-csv-converter/processor"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

func main() {
	ymlFile, err := ioutil.ReadFile("../../fixtures/sales_mappings_with_headers_and_converters_aliases.yml")
	if err != nil {
		panic(err)
	}

	var yml interface{}
	err = yaml.Unmarshal(ymlFile, &yml)
	if err != nil {
		panic(err)
	}

	yaml, _ := yml.(map[interface{}]interface{})
	entities := yaml["mappings"].(map[interface{}]interface{})

	csvFile, err := ioutil.ReadFile("../../fixtures/sales_with_headers.csv")
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(strings.NewReader(string(csvFile)))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	p := &processor.File{
		Filename:   "sales_with_headers.csv",
		Rows:       records,
		Mappings:   entities,
		HasHeaders: true,
	}

	p.ProcessFile()
}
