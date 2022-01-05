package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"

	"gorm.io/gorm"
)

// createRole will create predefined role and insert into DB.
func createRole(db *gorm.DB, role *entity.Role) (*entity.Role, error) {
	var roleExists entity.Role
	err := db.Where("name = ?", role.Name).Take(&roleExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(role).Error
			if err != nil {
				return role, err
			}
			return role, err
		}
		return role, err
	}
	return role, err
}
