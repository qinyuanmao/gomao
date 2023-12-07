package pipeline

import (
	"e.coding.net/tssoft/repository/gomao/pipeline/async"
	"e.coding.net/tssoft/repository/gomao/pipeline/sync"
)

func NewPipeline[T sync.Pipeline | async.Pipeline]() *T {
	return new(T)
}
