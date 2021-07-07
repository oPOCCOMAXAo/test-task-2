package service

import (
	"strconv"
	"test-task-20210707/mock"
	"test-task-20210707/test"
	"test-task-20210707/types"
	"testing"
)

type testCase struct {
	Responses mock.Responses
	Output    types.Output
}

func TestService(t *testing.T) {
	data := []testCase{
		{
			Responses: mock.Responses{
				R1: types.Response1{Total: 100},
				R2: types.Response2{Amount: 200},
				R3: types.Response2{Amount: 300},
			},
			Output: types.Output{OfferIndex: 1, Amount: 100},
		},
		{
			Responses: mock.Responses{
				R1: types.Response1{Total: 300},
				R2: types.Response2{Amount: 100},
				R3: types.Response2{Amount: 200},
			},
			Output: types.Output{OfferIndex: 2, Amount: 100},
		},
		{
			Responses: mock.Responses{
				R1: types.Response1{Total: 200},
				R2: types.Response2{Amount: 300},
				R3: types.Response2{Amount: 100},
			},
			Output: types.Output{OfferIndex: 3, Amount: 100},
		},
	}

	for i, datum := range data {
		t.Run("Test"+strconv.Itoa(i), func(t *testing.T) {
			mockServer := mock.New(datum.Responses)
			testService := New("http://localhost:" + strconv.Itoa(mockServer.GetPort()))

			defer mockServer.Close()
			go mockServer.Serve()

			result := testService.GetLowestOffer(types.Input{
				ShippingSourceAddress:      "1",
				ShippingDestinationAddress: "2",
				CartonDimensions:           []types.Size{{Height: 1, Width: 1, Length: 1}},
			})
			test.Check(t, "Service Result", datum.Output, result)
		})
	}
}
