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

func validateJSON(content []byte, jsonSchema bool) error {
	var jObj map[string]interface{}
	var j interface{}

	if jsonSchema {
		err := json.Unmarshal(content, &jObj)
		if err != nil {
			return fmt.Errorf("invalid json schema, %v", err)
		}
	} else {
		err := json.Unmarshal(content, &j)
		if err != nil {
			return err
		}
	}
	return nil
}

func escapeStr(content string) string {
	content = strings.Replace(content, "\\", "\\\\", -1)
	content = strings.Replace(content, "\"", "\\\"", -1)
	return content
}

func main() {
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

	if err = validateJSON(bytes, *schemaPtr); err != nil {
		log.Fatalf("Invalid json file '%s', error: %v", fileName, err)
	}

	var i interface{}
	json.Unmarshal(bytes, &i)
	bytes, _ = json.Marshal(i)
	str := string(bytes)

	if *escapedStringPtr {
		str = escapeStr(str)
	}
	fmt.Println(str)
}
