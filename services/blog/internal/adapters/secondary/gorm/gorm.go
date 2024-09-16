package gorm_adapter

import (
	"sync"

	"github.com/geekengineers/Microservice-Project-Demo/services/blog/internal/core/domain/article"
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

		db.AutoMigrate(&article.Article{})

		instance = db

		return db, nil
	}

	return instance, nil
}
