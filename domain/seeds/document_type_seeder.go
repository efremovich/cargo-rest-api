package seeds

import (
	"cargo-rest-api/domain/entity"
	"errors"
	"fmt"
	"log"

	"github.com/bxcodec/faker"

	"gorm.io/gorm"
)

// documentTypeFactory is a function uses to create []seed.Seed.
func documentTypeFactory() []Seed {
	fakerFactories := make([]Seed, 5)
	for i := 0; i < 5; i++ {
		a := entity.DocumentTypeFaker{}
		errFaker := faker.FakeData(&a)
		if errFaker != nil {
			log.Fatal(errFaker)
		}

		documentType := &entity.DocumentType{
			UUID: a.UUID,
			Type: a.Type,
		}
		fakerFactories[i] = Seed{
			Name: fmt.Sprintf("Create %s", a.Type),
			Run: func(db *gorm.DB) error {
				_, errDB := createDocumentType(db, documentType)
				return errDB
			},
		}
	}

	return fakerFactories
}

// createDocumentType will create fake documentType and insert into DB.
func createDocumentType(db *gorm.DB, documentType *entity.DocumentType) (*entity.DocumentType, error) {
	var documentTypeExists entity.DocumentType
	err := db.Where("type = ?", documentType.Type).Take(&documentTypeExists).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := db.Create(documentType).Error
			if err != nil {
				return documentType, err
			}
			return documentType, err
		}
		return documentType, err
	}
	return documentType, err
}
