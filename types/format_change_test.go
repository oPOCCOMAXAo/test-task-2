package types

import (
	"test-task-20210707/test"
	"testing"
)

func TestFromInput(t *testing.T) {
	sizes := []Size{
		{Width: 1, Height: 2, Length: 3},
	}
	testData := &Input{
		ShippingSourceAddress:      "1",
		ShippingDestinationAddress: "2",
		CartonDimensions:           sizes,
	}

	test.Check(t, "Request1FromInput", Request1{
		ContactAddress:    "2",
		WarehouseAddress:  "1",
		PackageDimensions: sizes,
	}, Request1FromInput(testData))

	test.Check(t, "Request2FromInput", Request2{
		Consignee: "2",
		Consignor: "1",
		Cartons:   sizes,
	}, Request2FromInput(testData))
}
