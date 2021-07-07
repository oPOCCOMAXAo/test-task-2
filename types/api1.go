package types

type Request1 struct {
	ContactAddress    string `json:"contact address"`
	WarehouseAddress  string `json:"warehouse address"`
	PackageDimensions []Size `json:"package dimensions"`
}

type Response1 struct {
	Total float64 `json:"total"`
}

func Request1FromInput(input *Input) Request1 {
	return Request1{
		ContactAddress:    input.ShippingDestinationAddress,
		WarehouseAddress:  input.ShippingSourceAddress,
		PackageDimensions: input.CartonDimensions,
	}
}
