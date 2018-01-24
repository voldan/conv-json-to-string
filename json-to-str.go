package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func isJSONFile(fileName string) bool {
	if filepath.Ext(fileName) == ".json" {
		return true
	}
	return false
}

func escapeStr(content string) string {
	content = strings.Replace(content, "\\", "\\\\", -1)
	content = strings.Replace(content, "\"", "\\\"", -1)
	return content
}

func main() {
	var jObj map[string]interface{}
	var j interface{}

	schemaPtr := flag.Bool("s", false, "json file is a json schema")
	escapedStringPtr := flag.Bool("e", false, "convert json to the escape string")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Error: no input file listed")
	}

	fileName := args[0]
	if !isJSONFile(fileName) {
		log.Fatalf("Error: extension of the input file '%s' is not json", fileName)
	}

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if *schemaPtr {
		err = json.Unmarshal(bytes, &jObj)
		if err != nil {
			log.Fatalf("Invalid json file '%s', error: invalid json schema, %v", fileName, err)
		}
		bytes, _ = json.Marshal(jObj)
	} else {
		err = json.Unmarshal(bytes, &j)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		bytes, _ = json.Marshal(j)
	}

	str := string(bytes)
	if *escapedStringPtr {
		str = escapeStr(str)
	}
	fmt.Println(str)
}
