package http

import (
	"github.com/panovateam/go-micro/registry"
)

type Options struct {
	Registry registry.Registry
}

type Option func(*Options)

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}
