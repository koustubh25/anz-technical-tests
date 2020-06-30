package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

const (
	versionFilePath       = "./resource/version.json"
	schemaVersionFilePath = "./resource/version_schema.json"
)

func Test_versionFileExists(t *testing.T) {

	absPath, _ := filepath.Abs(versionFilePath)
	fmt.Println(absPath)
	_, err := ioutil.ReadFile(absPath)
	if err != nil {
		t.Errorf("Expected version.json file. Couldn't read version.json file")
	}

}

func Test_validateVersionSchema(t *testing.T) {

	absPathVersion, _ := filepath.Abs(versionFilePath)
	absPathVersionSchema, _ := filepath.Abs(schemaVersionFilePath)

	documentLoader := gojsonschema.NewReferenceLoader("file:///" + absPathVersion)
	schemaLoader := gojsonschema.NewReferenceLoader("file:///" + absPathVersionSchema)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	// fmt.Println(version.Valid())
	if err != nil {
		t.Errorf("The version file format is not present in correct format")
		return
	}

	if !result.Valid() {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		t.Errorf("The version file format is not present in correct format")
	}

}
