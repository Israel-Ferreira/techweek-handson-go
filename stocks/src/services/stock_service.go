package services

type StockService interface {
	NewStockItem(interface{}) (interface{}, error)
	RemoveStockItem(any) error
	AddStockUpdate(interface{}) (interface{}, error)
	SubstractStockUpdate(interface{}) (interface{}, error)
	GetStock(sku string) (interface{}, error)
	GetStocks() ([]any, error)
}

type stockService struct{}

func (s stockService) AddStockUpdate(stockAmountBody interface{}) (any, error) {
	return nil, nil
}

func (s stockService) SubstractStockUpdate(stockAmountBody interface{}) (any, error) {
	return nil, nil
}

func (s stockService) NewStockItem(item any) (any, error) {
	return nil, nil
}

func (s stockService) RemoveStockItem(msg any) error {
	return nil
}

func (s stockService) GetStock(sku string) (any, error) {
	return nil, nil
}

func (s stockService) GetStocks() ([]any, error) {
	return nil, nil
}

func NewStockService() *stockService {
	return &stockService{}
}
