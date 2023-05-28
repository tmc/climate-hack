package graph

import "go-thorium/thoriumfacts"

type Resolver struct {
	service *thoriumfacts.Service
}

func NewResolver(service *thoriumfacts.Service) *Resolver {
	return &Resolver{service}
}
