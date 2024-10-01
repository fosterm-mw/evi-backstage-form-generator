package main 

import (
	"testing"
	"log"
)

func TestCRDIngestion(t *testing.T) {
	want := &CRD{}
	got := &CRD{}
	got = got.ingestCRD()
	log.Println(got.ApiVersion)
	log.Println(got.Spec.ClaimNames.Kind)
	if want != got {
		t.Fatalf("getValueByString(testData, claimName) = %v, wanted %v", got, want)
	}
}
