package sync

import (
	"context"
	"log"
	"sync"
)

type Input func(ctx context.Context) (<-chan any, error)

type Output func(ctx context.Context, parapipelines any) error

type Processor func(ctx context.Context, parapipelines any) (any, error)

type Pipeline struct {
	input  Input
	output Output
	ps     []Processor
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

func (pipeline *Pipeline) Run(ctx context.Context) error {
	var err error

	in, err := pipeline.input(ctx)
	if err != nil {
		return err
	}

	for data := range in {
		for _, v := range pipeline.ps {
			data, err = v(ctx, data)

			if err != nil {
				log.Printf("Process err %s\n", err)
				return nil
			}
		}

		err := pipeline.output(ctx, data)
		if err != nil {
			log.Printf("Output err %s\n", err)
			return nil
		}
	}

	return nil
}

func (pipeline *Pipeline) RunN(ctx context.Context, pipelineaxCnt int) error {
	var err error

	in, err := pipeline.input(ctx)
	if err != nil {
		return err
	}

	// pipeline构建和执行
	syncProcess := func(data any) {
		for _, v := range pipeline.ps {
			data, err = v(ctx, data)

			// 错误集中处理，这里选择提前退出
			if err != nil {
				log.Printf("Process err %s\n", err)
				return
			}
		}

		err := pipeline.output(ctx, data)
		if err != nil {
			log.Printf("Output err %s\n", err)
			return
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(pipelineaxCnt)

	// 多个协程消费同一个channel
	for i := 0; i < pipelineaxCnt; i++ {
		go func() {
			defer wg.Done()

			for data := range in {
				syncProcess(data)
			}
		}()
	}

	wg.Wait()

	return nil
}
