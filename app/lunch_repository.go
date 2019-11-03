package app

import (
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

func (o LunchRepository) Any(predicate LunchPredicate) bool {
	var lunches []Lunch
	if result := o.db.Find(&lunches); result.Error != nil {
		return false
	}

	return From(lunches).AnyWithT(predicate)
}

func (o LunchRepository) Close() error {
	if o.db != nil {
		return o.db.Close()
	}

	return nil
}
