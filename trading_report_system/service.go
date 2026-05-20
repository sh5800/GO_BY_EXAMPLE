package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Filters  ReportFilters
	RespChan chan JobResult
}

type JobResult struct {
	Report *PaginatedResponse
	Error  error
}

type Service struct {
	repo    *Repository
	jobChan chan Job
	cache   sync.Map // Simple in-memory cache
	mu      sync.Mutex
}

func NewService(repo *Repository) *Service {
	s := &Service{
		repo:    repo,
		jobChan: make(chan Job, 100),
	}

	// Start worker pool
	for i := 0; i < 10; i++ {
		go s.worker(i)
	}

	return s
}

func (s *Service) worker(id int) {
	fmt.Printf("Worker %d started\n", id)
	for job := range s.jobChan {
		// 1. Check Cache
		cacheKey := fmt.Sprintf("%v", job.Filters)
		if val, ok := s.cache.Load(cacheKey); ok {
			job.RespChan <- JobResult{Report: val.(*PaginatedResponse)}
			continue
		}

		// 2. Execute Query with Timeout (30s)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		report, err := s.repo.ExecuteReport(ctx, job.Filters)
		cancel()

		if err == nil {
			// 3. Update Cache
			s.cache.Store(cacheKey, report)
		}

		job.RespChan <- JobResult{Report: report, Error: err}
	}
}

func (s *Service) GetReport(filters ReportFilters) (*PaginatedResponse, error) {
	respChan := make(chan JobResult)
	s.jobChan <- Job{
		Filters:  filters,
		RespChan: respChan,
	}

	result := <-respChan
	return result.Report, result.Error
}
