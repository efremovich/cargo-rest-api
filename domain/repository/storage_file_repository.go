package repository

import "cargo-rest-api/domain/entity"

type StorageFileRepository interface {
	SaveFile(file *entity.StorageFile) (*entity.StorageFile, map[string]string, error)
	GetFile(string) (*entity.StorageFile, error)
}
