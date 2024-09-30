package main

import (
	"reflect"
	"log"
	"errors"
)

func (c *CRD) parseCRD() error {
	yaml := readYaml()
	spec := yaml["spec"]
	var claimName interface{}
	if claimName = getValueByString(getValueByString(spec, "claimNames"), "kind"); reflect.TypeOf(claimName) != reflect.TypeOf("") {
		log.Fatal("Error parsing claim name kind. Got no string value.")
		return errors.New("Claim Name not available.")
	}
	c.claimName = claimName.(string)
	return nil
}

func getValueByString(spec interface{}, searchTerm string) interface{} {
	for k, v := range spec.(map[interface{}]interface{}) {
		if k == searchTerm {
			return v
		}
	}
	return nil
}

