package types

type Request2 struct {
	XMLName   struct{} `json:"-" xml:"request"`
	Consignee string   `json:"consignee" xml:"consignee"`
	Consignor string   `json:"consignor" xml:"consignor"`
	Cartons   []Size   `json:"cartons" xml:"cartons>carton"`
}

type Response2 struct {
	XMLName struct{} `json:"-" xml:"result"`
	Amount  float64  `json:"amount" xml:"amount"`
}

func Request2FromInput(input *Input) Request2 {
	return Request2{
		Consignee: input.ShippingDestinationAddress,
		Consignor: input.ShippingSourceAddress,
		Cartons:   input.CartonDimensions,
	}
}
