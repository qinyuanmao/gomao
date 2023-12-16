package pipeline

import (
	"e.coding.net/tssoft/repository/gomao/pipeline/async"
	"e.coding.net/tssoft/repository/gomao/pipeline/sync"
)

type Pipeline interface {
	sync.Pipeline | async.Pipeline
}

func NewPipeline[T Pipeline]() *T {
	return new(T)
}
