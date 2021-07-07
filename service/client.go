package service

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"test-task-20210707/types"
)

type Client struct {
	inputMarshal    func(*types.Input) ([]byte, error)
	outputUnmarshal func([]byte) *types.Output
	url             string
	contentType     string
}

func (c *Client) doRequest(url, contentType string, data []byte) (status int, body []byte) {
	resp, err := http.DefaultClient.Post(url, contentType, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	status = resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return
}

func (c *Client) Request(input *types.Input) *types.Output {
	data, _ := c.inputMarshal(input)
	_, body := c.doRequest(c.url, c.contentType, data)
	return c.outputUnmarshal(body)
}
