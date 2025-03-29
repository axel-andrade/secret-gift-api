package postgres_mappers

import (
	postgres_models "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/models"
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type BasePostgresMapper struct{}

func (m *BasePostgresMapper) toDomain(model postgres_models.BaseModel) *domain.Base {
	return &domain.Base{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (m *BasePostgresMapper) toPersistence(entity domain.Base) *postgres_models.BaseModel {
	return &postgres_models.BaseModel{
		ID:        entity.ID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
