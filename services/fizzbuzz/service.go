package fizzbuzz

import (
	"strconv"

	"github.com/rochdi/fizz-buzz/common"
)

// Service fizz buzz contract
type Service interface {
	FizzBuzzIt(*common.RequestParams) (string, error)
}

type service struct {
}

// NewService returns a new instance of fizz buzz service
func NewService() Service {
	return &service{}
}

func (s *service) FizzBuzzIt(params *common.RequestParams) (string, error) {
	result := ""
	fizzbuzzable := fizzbuzzable(params.Int1, params.Int2, params.Str1, params.Str2)
	for i := 0; i < params.Limit; i++ {
		result += fizzbuzzable()
	}
	return result, nil
}

func fizzbuzzable(int1 int, int2 int, str1 string, str2 string) func() string { // fizz buzz enumerable
	n := 1
	return func() string {
		k := n // capture current step
		n++
		if k%int1 == 0 && k%int2 == 0 { // This use case if not clear in the specification, moving forward with returning both matches
			return str1 + str2
		}

		if k%int1 == 0 {
			return str1
		}

		if k%int2 == 0 {
			return str2
		}

		return strconv.Itoa(k)
	}
}
