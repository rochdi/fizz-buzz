package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rochdi/fizz-buzz/common"
	"github.com/rochdi/fizz-buzz/services/fizzbuzz"
	"github.com/rochdi/fizz-buzz/services/health"
)

// FizzBuzzController defines the health controller
type FizzBuzzController struct {
	fizzbuzzSrv fizzbuzz.Service
	healthSrv   health.Service
}

// NewFizzBuzzController return a new instance of FizzBuzzController
func NewFizzBuzzController(fsvc fizzbuzz.Service, hsrv health.Service) *FizzBuzzController {
	return &FizzBuzzController{fizzbuzzSrv: fsvc, healthSrv: hsrv}
}

// Stats allows users to know what the most frequent request has been
func (c *FizzBuzzController) FizzBuzz(rw http.ResponseWriter, r *http.Request) {
	request := common.RequestParams{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil || !request.IsValid() {
		http.Error(rw, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // invalid request
	}

	result, err := c.fizzbuzzSrv.FizzBuzzIt(&request)
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	go c.healthSrv.AddRequest(&request)
	json.NewEncoder(rw).Encode(result)
}
