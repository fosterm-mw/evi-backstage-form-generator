package main

type CRD struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Metadata interface{} `yaml:"metadata"`
	Spec CRDSpec `yaml:"spec"`
}

type CRDSpec struct {
	ClaimNames CRDSpecNames `yaml:"claimNames"`
	DefaultCompositeDeletePolicy string `yaml:"defaultCompositeDeletePolicy"`
	DefaultCompositionRef struct {Name string `yaml:"name"`} `yaml:"defaultCompositionRef"`
	DefaultCompositionUpdatePolicy string `yaml:"defaultCompositionUpdatePolicy"`
	Group string `yaml:"group"`
	Names CRDSpecNames `yaml:"names"`
	Versions []CRDVersions `yaml:"versions"`
}

type CRDSpecNames struct {
	Categories []string `yaml:"categories"`
	Kind string `yaml:"kind"`
	Plural string `yaml:"plural"`
}

type CRDVersions struct {
	Name string `yaml:"name"`
	Referenceable bool `yaml:"referenceable"`
	Schema struct {OpenAPIV3Schema map[interface{}]interface{} `yaml:"openAPIV3Schema"`} `yaml:"schema"`
}

type Field struct {
	Type string
	Title string
	Description string
	Required bool
}
