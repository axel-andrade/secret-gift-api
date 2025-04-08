package postgres_mappers

import (
	postgres_models "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/models"
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type AuthorizedGiftMapper struct {
	BaseMapper BasePostgresMapper
}

func (m *AuthorizedGiftMapper) ToDomain(model postgres_models.AuthorizedGiftModel) *domain.AuthorizedGift {
	entity := &domain.AuthorizedGift{
		Base:           *m.BaseMapper.toDomain(model.BaseModel),
		GiftID:         model.Gift.ID,
		ExpirationDate: model.ExpirationDate,
	}

	return entity
}

func (m *AuthorizedGiftMapper) ToPersistence(entity domain.AuthorizedGift) *postgres_models.AuthorizedGiftModel {
	return &postgres_models.AuthorizedGiftModel{
		BaseModel:      *m.BaseMapper.toPersistence(entity.Base),
		GiftID:         string(entity.GiftID),
		ExpirationDate: entity.ExpirationDate,
	}
}
