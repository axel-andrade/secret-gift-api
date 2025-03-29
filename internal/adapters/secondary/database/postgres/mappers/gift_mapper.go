package postgres_mappers

import (
	postgres_models "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/models"
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type GiftMapper struct {
	BaseMapper BasePostgresMapper
}

func (m *GiftMapper) ToDomain(model postgres_models.GiftModel) *domain.Gift {
	entity := &domain.Gift{
		Base:       *m.BaseMapper.toDomain(model.BaseModel),
		Status:     domain.GiftStatus(model.Status),
		Type:       domain.GiftType(model.Type),
		CPF:        model.CPF,
		Amount:     model.Amount,
		PixKeyType: model.PixKeyType,
		PixKey:     model.PixKey,
		Message:    model.Message,
	}

	return entity
}

func (m *GiftMapper) ToPersistence(entity domain.Gift) *postgres_models.GiftModel {
	return &postgres_models.GiftModel{
		BaseModel:  *m.BaseMapper.toPersistence(entity.Base),
		Status:     string(entity.Status),
		Type:       string(entity.Type),
		CPF:        entity.CPF,
		Amount:     entity.Amount,
		PixKeyType: entity.PixKeyType,
		PixKey:     entity.PixKey,
		Message:    entity.Message,
	}
}
