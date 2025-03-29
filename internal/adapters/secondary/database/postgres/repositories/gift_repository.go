package postgres_repositories

import (
	postgres_mappers "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/mappers"
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
)

type GiftPostgresRepository struct {
	*BasePostgresRepository
	GiftMapper postgres_mappers.GiftMapper
}

func BuildGiftPostgresRepository() *GiftPostgresRepository {
	return &GiftPostgresRepository{BasePostgresRepository: BuildBasePostgresRepository()}
}

func (r *GiftPostgresRepository) CreateGift(gift *domain.Gift) (*domain.Gift, error) {
	model := r.GiftMapper.ToPersistence(*gift)

	q := r.getQueryOrTx()

	err := q.Create(model).Error

	if err != nil {
		return nil, err
	}

	return r.GiftMapper.ToDomain(*model), nil
}
