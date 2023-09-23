package sca

import (
	"net/http"

	"github.com/kyselabs/kyse-sdk-go/kyse"
	"github.com/kyselabs/kyse-sdk-go/rest"
)

type SCA struct {
	Kyse *kyse.Kyse
}

func NewSCA(kyse *kyse.Kyse) *SCA {
	return &SCA{Kyse: kyse}
}

func (s *SCA) Audit(resources []Resource) (verdictedResources []VerdictedResource) {
	response, err := s.Kyse.RestClient.Request(
		http.MethodPost,
		"/sca/audit",
		map[string]string{"Accept": rest.NewMime("1").Accept()},
		resources,
	)

	if err == nil {
		response.Unmarshal(&verdictedResources)
	}

	return
}
