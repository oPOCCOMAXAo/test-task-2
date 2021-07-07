package service

import (
	"encoding/json"
	"encoding/xml"
	"test-task-20210707/types"
)

func MakeAPI1(host string) *Client {
	return &Client{
		inputMarshal: func(input *types.Input) ([]byte, error) {
			return json.Marshal(types.Request1FromInput(input))
		},
		outputUnmarshal: func(data []byte) (res *types.Output) {
			var response types.Response1
			_ = json.Unmarshal(data, &response)
			return &types.Output{
				OfferIndex: 1,
				Amount:     response.Total,
			}
		},
		url:         host + "/api1",
		contentType: "application/json",
	}
}

func MakeAPI2(host string) *Client {
	return &Client{
		inputMarshal: func(input *types.Input) ([]byte, error) {
			return json.Marshal(types.Request2FromInput(input))
		},
		outputUnmarshal: func(data []byte) (res *types.Output) {
			var response types.Response2
			_ = json.Unmarshal(data, &response)
			return &types.Output{
				OfferIndex: 2,
				Amount:     response.Amount,
			}
		},
		url:         host + "/api2json",
		contentType: "application/json",
	}
}

func MakeAPI3(host string) *Client {
	return &Client{
		inputMarshal: func(input *types.Input) ([]byte, error) {
			return xml.Marshal(types.Request2FromInput(input))
		},
		outputUnmarshal: func(data []byte) (res *types.Output) {
			var response types.Response2
			_ = xml.Unmarshal(data, &response)
			return &types.Output{
				OfferIndex: 3,
				Amount:     response.Amount,
			}
		},
		url:         host + "/api2xml",
		contentType: "application/xml",
	}
}
