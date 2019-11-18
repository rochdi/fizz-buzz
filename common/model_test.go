package common_test

import (
	"testing"

	"github.com/rochdi/fizz-buzz/common"
)

func TestGetID(t *testing.T) {
	request := common.RequestParams{Int1: 12, Int2: 15, Limit: 20, Str1: "fizz", Str2: "buzz"}
	actual := request.GetID()
	expected := "12-15-20-fizz-buzz"
	if actual != expected {
		t.Fatalf("actual: %s\nexpected: %s", actual, expected)
	}
}

func TestIsValidWihPositiveLimit(t *testing.T) {
	request := common.RequestParams{Int1: 12, Int2: 15, Limit: 20, Str1: "fizz", Str2: "buzz"}
	actual := request.IsValid()
	expected := true
	if actual != expected {
		t.Fatalf("actual: %v\nexpected: %v", actual, expected)
	}
}

func TestIsValidWithNegativeLimit(t *testing.T) {
	request := common.RequestParams{Int1: 12, Int2: 15, Limit: -20, Str1: "fizz", Str2: "buzz"}
	actual := request.IsValid()
	expected := false
	if actual != expected {
		t.Fatalf("actual: %v\nexpected: %v", actual, expected)
	}
}
