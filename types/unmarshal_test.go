package types

import (
	"encoding/json"
	"encoding/xml"
	"test-task-20210707/test"
	"testing"
)

// response body
func TestUnmarshal(t *testing.T) {
	var res1 Response1
	_ = json.Unmarshal([]byte(`{"total":123}`), &res1)
	test.Check(t, "Response1 Unmarshal JSON", Response1{Total: 123}, res1)

	var res2 Response2
	_ = json.Unmarshal([]byte(`{"amount":456}`), &res2)
	test.Check(t, "Response2 Unmarshal JSON", Response2{Amount: 456}, res2)

	_ = xml.Unmarshal([]byte(`<result><amount>789</amount></result>`), &res2)
	test.Check(t, "Response2 Unmarshal XML", Response2{Amount: 789}, res2)
}
