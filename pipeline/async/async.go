package async

import (
	"context"
	"sync"

	"e.coding.net/tssoft/repository/gomao/logger"
)

type Input func(ctx context.Context, wg *sync.WaitGroup, errChan chan error) <-chan int

type Output func(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error)

type Processor func(ctx context.Context, wg *sync.WaitGroup, dataChan <-chan int, errChan chan error) <-chan int

type Pipeline struct {
	input   Input
	output  Output
	ps      []Processor
	errChan chan error
}

func (pipeline *Pipeline) AddProcessor(processor Processor) *Pipeline {
	pipeline.ps = append(pipeline.ps, processor)
	return pipeline
}

func (pipeline *Pipeline) SetInput(input Input) *Pipeline {
	pipeline.input = input
	return pipeline
}

func (pipeline *Pipeline) SetOutput(output Output) *Pipeline {
	pipeline.output = output
	return pipeline
}

func (pipeline *Pipeline) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg = sync.WaitGroup{}

	// 组装pipeline
	wg.Add(1)
	dataChan := pipeline.input(ctx, &wg, pipeline.errChan)

	for _, v := range pipeline.ps {
		wg.Add(1)
		dataChan = v(ctx, &wg, dataChan, pipeline.errChan)
	}

	wg.Add(1)
	pipeline.output(ctx, &wg, dataChan, pipeline.errChan)

	go func() {
		wg.Wait()
		close(pipeline.errChan)
	}()

	// 错误通道阻塞，错误处理集中处理
	for {
		select {
		case err, ok := <-pipeline.errChan:
			if !ok {
				logger.Info("error channel closed and exit")
				return
			}

			logger.Errorf("receive error: %s", err)
			cancel()
		}
	}
}
