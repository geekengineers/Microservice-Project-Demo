package sqlite_adapter

import (
	"sync"

	"gorm.io/gorm"
)

var (
	mu       sync.Mutex
	instance *gorm.DB
)

func GORM(dialector gorm.Dialector) (*gorm.DB, error) {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()

		db, err := gorm.Open(dialector, &gorm.Config{})
		if err != nil {
			return nil, err
		}

		instance = db

		return db, nil
	}

	return instance, nil
}
