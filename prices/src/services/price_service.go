package services

import (
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/models"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/repositories"
)

type PriceService interface {
	GetAll() ([]models.Price, error)
	GetBySku(string) (models.Price, error)
	AddItem(interface{}) (string, error)
	Update(string, data.UpdatePrice) (models.Price, error)
}

type priceService struct {
	Repo repositories.PriceRepository
}

func (p *priceService) GetAll() ([]models.Price, error) {
	prices, err := p.Repo.GetAll()

	if err != nil {
		return nil, err
	}

	return prices, nil
}

func (p *priceService) GetBySku(sku string) (models.Price, error) {
	price, err := p.Repo.GetBySku(sku)

	if err != nil {
		return models.Price{}, nil
	}

	return price, nil
}

func (p *priceService) AddItem(priceDTO interface{}) (string, error) {
	return "", nil
}

func (p *priceService) Update(sku string, priceDTO data.UpdatePrice) (models.Price, error) {

	if err := priceDTO.Validate(); err != nil {
		return models.Price{}, err
	}

	price, err := p.Repo.Update(sku, priceDTO)

	if err != nil {
		return models.Price{}, err
	}

	return price, nil
}

func NewPriceService(repo repositories.PriceRepository) *priceService {
	return &priceService{
		Repo: repo,
	}
}
