package service

import (
	"sync"
	"test-task-20210707/types"
)

type Service struct {
	clients []*Client
}

func New(host string) *Service {
	return &Service{
		clients: []*Client{
			MakeAPI1(host),
			MakeAPI2(host),
			MakeAPI3(host),
		},
	}
}

func (s *Service) GetLowestOffer(input types.Input) types.Output {
	total := len(s.clients)
	outputs := make([]*types.Output, total)

	var wg sync.WaitGroup
	wg.Add(total)
	for i, client := range s.clients {
		go func(i int, client *Client) {
			outputs[i] = client.Request(&input)
			wg.Done()
		}(i, client)
	}
	wg.Wait()

	result := outputs[0]
	for _, output := range outputs {
		if result.Amount > output.Amount {
			result = output
		}
	}

	return *result
}
