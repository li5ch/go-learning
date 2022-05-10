package main

import (
	"log"
	"sync"
)

type (
	// ParallelProcessorOptions is the configs for ParallelProcessor
	ParallelProcessorOptions struct {
		QueueSize   int
		WorkerCount int
	}

	ParallelProcessor struct {
		status  int32
		options *ParallelProcessorOptions

		logger       log.Logger

		//tasksChan    chan Task
		shutdownChan chan struct{}
		workerWG     sync.WaitGroup
	}
)

