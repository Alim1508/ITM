package service

import (
	"fmt"
	"strings"
	"sync"
)

type Service struct {
	prod Producer
	pres Presenter
}

func NewService(p Producer, pr Presenter) *Service {
	return &Service{
		prod: p,
		pres: pr,
	}
}
 
func (s *Service) Run() error {
	lines, err := s.prod.Produce()
	if err != nil {
		return fmt.Errorf("error producing data: %w", err)
	}


	inputChan := make(chan string, len(lines))
	outputChan := make(chan string, len(lines))


	limitChan := make(chan struct{}, 10) // Вместо семафора используем канал с буфером 10

	var wg sync.WaitGroup

	for _, line := range lines {
		inputChan <- line
		wg.Add(1)
		go func(text string) {
			defer wg.Done()

		    limitChan <- struct{}{}

			result := maskLinks(text)
			outputChan <- result

			<-limitChan
		}(line)
	}

	close(inputChan)

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	var results []string
	for result := range outputChan {
		results = append(results, result)
	}

	if err := s.pres.Present(results); err != nil {
		return fmt.Errorf("error presenting data: %w", err)
	}

	return nil
}
