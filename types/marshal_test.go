package types

import (
	"encoding/json"
	"encoding/xml"
	"test-task-20210707/test"
	"testing"
)

// request body
func TestMarshal(t *testing.T) {
	data1, _ := json.Marshal(Request1{
		ContactAddress:   "123",
		WarehouseAddress: "456",
		PackageDimensions: []Size{
			{Width: 1, Height: 2, Length: 3},
		},
	})
	test.Check(t, "Request1 Marshal JSON", []byte(`{"contact address":"123","warehouse address":"456","package dimensions":[{"width":1,"height":2,"length":3}]}`), data1)

	data2, _ := json.Marshal(Request2{
		Consignee: "123",
		Consignor: "456",
		Cartons: []Size{
			{Width: 1, Height: 2, Length: 3},
		},
	})
	test.Check(t, "Request2 Marshal JSON", []byte(`{"consignee":"123","consignor":"456","cartons":[{"width":1,"height":2,"length":3}]}`), data2)

	data3, _ := xml.Marshal(Request2{
		Consignee: "123",
		Consignor: "456",
		Cartons: []Size{
			{Width: 1, Height: 2, Length: 3},
			{Width: 4, Height: 5, Length: 6},
		},
	})
	test.Check(t, "Request2 Marshal XML", []byte(`<request><consignee>123</consignee><consignor>456</consignor><cartons><carton width="1" height="2" length="3"></carton><carton width="4" height="5" length="6"></carton></cartons></request>`), data3)
}
