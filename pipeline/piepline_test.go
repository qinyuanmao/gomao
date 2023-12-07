package pipeline

import (
	"context"
	"testing"
	"time"

	"e.coding.net/tssoft/repository/gomao/pipeline/sync"
	"github.com/stretchr/testify/assert"
)

func TestAsync(t *testing.T) {
	assert.Equal(t, true, "true")
}

func TestSync(t *testing.T) {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var channel = make(chan any)

	for _, v := range array {
		go func(v int) {
			channel <- v
			if v == 8 {
				close(channel)
			} else {
				time.Sleep(1 * time.Second)
			}
		}(v)
	}

	var result = 1

	NewPipeline[sync.Pipeline]().
		SetInput(func(ctx context.Context) (<-chan any, error) {
			return channel, nil
		}).
		AddProcessor(func(ctx context.Context, params any) (any, error) {
			return params.(int) * 2, nil
		}).
		AddProcessor(func(ctx context.Context, params any) (any, error) {
			return params.(int) * 2, nil
		}).
		SetOutput(func(ctx context.Context, params any) error {
			assert.Equal(t, result*4, params.(int))
			result++
			return nil
		}).
		Run(context.Background())
}
