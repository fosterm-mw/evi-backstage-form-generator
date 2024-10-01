package main

import (
	"reflect"
	"testing"
)

func TestCRDIngestion(t *testing.T) {
	names := &CRDSpecNames{
		Categories: []string{
			"crossplane",
			"composition",
			"app",
		},
		Kind:   "ApClaim",
		Plural: "appclaims",
	}
	want := &CRD{
		ApiVersion: "apiextensions.crossplane.io/v1",
		Kind:       "CompositeResourceDefinition",
		Metadata:   map[interface{}]interface{}{"name": "xapps.app.idp.mayo.edu"},
		Spec: CRDSpec{
			ClaimNames:                   *names,
			DefaultCompositeDeletePolicy: "Background",
			DefaultCompositionRef: struct {
				Name string `yaml:"name"`
			}{Name: "xapps-v1alpha1.app.idp.mayo.edu"},
			DefaultCompositionUpdatePolicy: "Automatic",
			Group:                          "app.idp.mayo.edu",
			Names:                          *names,
			Versions: []CRDVersions{
				{
					Name:          "v1alpha1",
					Referenceable: true,
					Schema: struct {
						OpenAPIV3Schema map[interface{}]interface{} `yaml:"openAPIV3Schema"`
					}{},
				},
			},
		},
	}
	got := &CRD{}
	got = got.ingestCRD()
	testCRDEquivalence(t, got, want)
}

func testCRDEquivalence(t *testing.T, got, want *CRD) {
	gotValues := reflect.ValueOf(*got)
	wantValues := reflect.ValueOf(*want)
	gotTypes := gotValues.Type()
	wantTypes := wantValues.Type()

	for i := 0; i < gotValues.NumField(); i++ {
		// TODO: Handle Schema
		if gotTypes.Field(i).Name == "Schema" {
			return
		}
		if !reflect.DeepEqual(gotValues.Field(i).Interface(), wantValues.Field(i).Interface()) {
			t.Fatalf("Parsing CRD failed, got %v for %v, wanted %v for %v", gotValues.Field(i), gotTypes.Field(i).Name, wantValues.Field(i), wantTypes.Field(i).Name)
		}
	}
}
