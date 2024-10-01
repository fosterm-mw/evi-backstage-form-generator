package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func readYaml() map[string]interface{} {
	crd := make(map[string]interface{})

	yamlFile, err := os.ReadFile("./crds/app-definition.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, crd)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return crd
}

func (crd *CRD) ingestCRD() *CRD {
	yamlFile, err := os.ReadFile("./crds/app-definition.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, crd)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return crd
}
