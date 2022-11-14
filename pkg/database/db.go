package database

//go:generate mockgen -destination=../mocks/mock_db.go -package=mocks gigmile/pkg/database DB

import "gigmile/pkg/model"

type DB interface {
	CreateCountry(user *model.Country) (*model.Country, error)
	FindCountryByID(ID string) (*model.Country, error)
	UpdateCountry(id string, update map[string]interface{}) error
	GetCountries() []model.Country
}
