package types

type Size struct {
	Width  float64 `json:"width" xml:"width,attr"`
	Height float64 `json:"height" xml:"height,attr"`
	Length float64 `json:"length" xml:"length,attr"`
}

type Input struct {
	ShippingSourceAddress      string
	ShippingDestinationAddress string
	CartonDimensions           []Size
}

type Output struct {
	OfferIndex int
	Amount     float64
}
