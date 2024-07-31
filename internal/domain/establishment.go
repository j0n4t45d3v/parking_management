package domain

type Addrees struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	City         string `json:"city"`
	State        string `json:"state"`
	Neighborhood string `json:"neighborhood"`
}

type Establishment struct {
	Name           string  `json:"name"`
	Document       string  `json:"document"`
	Addrees        Addrees `json:"addrees"`
	Phone          string  `json:"phone"`
	QtdMotorcycles int32   `json:"max_quantity_motocycles"`
	QtdCars        int32   `json:"max_quantity_cars"`
}
