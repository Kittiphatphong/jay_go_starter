package partners

import "go_starter/trails"

type Partner interface {
}

type partner struct {
	httpClientTrail trails.HttpClientTrail
}

func NewPartner(httpClientTrail trails.HttpClientTrail) Partner {
	return &partner{
		httpClientTrail: httpClientTrail,
	}
}
