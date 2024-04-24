package integration

type Integrations struct {
	CarsInfoClient *CarsInfoClient
}

type DepIntegrations struct {
	CarsInfoURL string
}

func NewIntegrations(dep DepIntegrations) *Integrations {
	return &Integrations{CarsInfoClient: NewCarsInfoClient(dep.CarsInfoURL)}
}
