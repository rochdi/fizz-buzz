package fizzbuzz_test

import (
	"testing"

	"github.com/rochdi/fizz-buzz/common"
	"github.com/rochdi/fizz-buzz/services/fizzbuzz"
)

func TestFizzBuzzItHappyPath(t *testing.T) {
	fizbuzz := fizzbuzz.NewService()
	params := common.RequestParams{Int1: 2, Int2: 3, Limit: 7, Str1: "fizz", Str2: "buzz"}
	actual, err := fizbuzz.FizzBuzzIt(&params)
	if err != nil {
		t.Fatal("no error must be returned")
	}

	expected := "1fizzbuzzfizz5fizzbuzz7"
	if actual != expected {
		t.Fatalf("actual: %s\nexpected: %s", actual, expected)
	}

}

func TestFizzBuzzItWithInt2MultipleofInt1(t *testing.T) {
	fizbuzz := fizzbuzz.NewService()
	params := common.RequestParams{Int1: 2, Int2: 4, Limit: 7, Str1: "fizz", Str2: "buzz"}
	actual, err := fizbuzz.FizzBuzzIt(&params)
	if err != nil {
		t.Fatal("no error must be returned")
	}

	expected := "1fizz3fizzbuzz5fizz7"
	if actual != expected {
		t.Fatalf("actual: %s\nexpected: %s", actual, expected)
	}

}

func TestFizzBuzzItWithNoMultipleWithinTheLimit(t *testing.T) {
	fizbuzz := fizzbuzz.NewService()
	params := common.RequestParams{Int1: 200, Int2: 400, Limit: 7, Str1: "fizz", Str2: "buzz"}
	actual, err := fizbuzz.FizzBuzzIt(&params)
	if err != nil {
		t.Fatal("no error must be returned")
	}

	expected := "1234567"
	if actual != expected {
		t.Fatalf("actual: %s\nexpected: %s", actual, expected)
	}
}
