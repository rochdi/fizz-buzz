package health

import (
	"github.com/rochdi/fizz-buzz/common"
)

// Service defines the health operations
type Service interface {
	Health() error
	GetStats() (*common.RequestStats, error)
	AddRequest(*common.RequestParams)
}

type service struct {
	analytics       map[string]*common.RequestStats
	mostUsedRequest *common.RequestStats
}

// NewService returns an instance of the health service
func NewService() Service {
	return &service{analytics: make(map[string]*common.RequestStats), mostUsedRequest: &common.RequestStats{}}
}

func (s *service) Health() error {
	return nil // Naive implementation of a watch dog, we may need to run a couple of IO checks in real life
}

// GetStats returns the precomputed version of most used request, this will enable async processing and increase performance
// but in this case we loose conistancy which is not an issue in this case based on the spec
func (s *service) GetStats() (*common.RequestStats, error) {
	return s.mostUsedRequest, nil // return the precomputed version of most used request, this will enable async processing and increase performance
	// but in this case we loose conistancy
}

// AddRequest add a new request to the analytics collection, set the right Hits and update the mosr used request
// this method is meant to be used in async way
func (s *service) AddRequest(params *common.RequestParams) {
	id := params.GetID()
	if stats, ok := s.analytics[id]; ok {
		stats.Hits++
	}
	s.analytics[id] = &(common.RequestStats{Request: params, Hits: 1})
	//update themost visited (not the most efficent way in term of computing working woth a sorted map will make it faster)
	for _, stats := range s.analytics {
		if stats.Hits > s.mostUsedRequest.Hits {
			s.mostUsedRequest = stats
		}
	}
}
