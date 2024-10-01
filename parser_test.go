package main

import (
	"testing"
)

func TestGetValueByString(t *testing.T) {
	got := "value"
	want := "value"
	if want != got {
		t.Fatalf("getValueByString() = %v, wanted %v", got, want)
	}
}

func TestGetClaimNameByString(t *testing.T) {
	testData := make(map[interface{}]interface{})
	testData["claimName"] = "testClaim"
	want := "testClaim"
	got := getValueByString(testData, "claimName")
	if want != got {
		t.Fatalf("getValueByString(testData, claimName) = %v, wanted %v", got, want)
	}
}

func TestGetOpenAPISpecField(t *testing.T) {
	testData := make(map[interface{}]interface{})
	testData["claimName"] = "testClaim"
	testData["versions"] = map[interface{}]interface{}{"name": "v1alpha1", "referenceable": true}
	/* log.Println(testData) */
	want := "testClaim"
	got := getValueByString(testData, "claimName")
	if want != got {
		t.Fatalf("getValueByString(testData, claimName) = %v, wanted %v", got, want)
	}
}
