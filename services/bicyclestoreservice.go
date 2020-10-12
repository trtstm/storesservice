package bicyclestoreservice

type BicycleStoreService struct {
	apikey string
}

func NewBicycleStoreService(apiKey string) *BicycleStoreService {
	return &BicycleStoreService{
		apikey: apiKey,
	}
}
