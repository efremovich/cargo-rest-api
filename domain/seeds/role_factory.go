package seeds

import (
	"cargo-rest-api/domain/entity"
	"fmt"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

// roleFactory is a function uses to create []seed.Seed.
func roleFactory() []Seed {
	roles := []*entity.Role{
		{UUID: uuid.New().String(), Name: "Super Administrator"},
		{UUID: uuid.New().String(), Name: "Administrator"},
		{UUID: uuid.New().String(), Name: "User"},
	}

	fakerFactories := make([]Seed, len(roles))
	for i, r := range roles {
		cr := r
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", cr.Name),
			Run: func(db *gorm.DB) error {
				_, err := createRole(db, cr)
				return err
			},
		}
	}

	return fakerFactories
}
