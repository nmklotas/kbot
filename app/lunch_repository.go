package app

import (
	"errors"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type LunchRepository struct {
	db *gorm.DB
}

type Lunch struct {
	gorm.Model
}

type LunchPredicate func(Lunch) bool

func OpenLunchRepository() (*LunchRepository, error) {
	db, err := gorm.Open("sqlite3", "orders.db")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Lunch{})
	return &LunchRepository{db}, nil
}

func (o LunchRepository) Save(lunch Lunch) error {
	return o.db.Create(&lunch).Error
}

func (o LunchRepository) Find(predicate LunchPredicate) (*Lunch, error) {
	var lunches []Lunch

	if result := o.db.Find(&lunches); result.Error != nil {
		return nil, result.Error
	}

	if result, ok := From(lunches).FirstWithT(predicate).(Lunch); ok {
		return &result, nil
	}

	return nil, errors.New("Lunch not found")
}

func (o LunchRepository) Close() error {
	if o.db != nil {
		return o.db.Close()
	}

	return nil
}
