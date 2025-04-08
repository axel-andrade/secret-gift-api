package postgres_repositories

import (
	postgres_mappers "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/mappers"
	postgres_models "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/models"
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
	"gorm.io/gorm"
)

type GiftPostgresRepository struct {
	*BasePostgresRepository
	GiftMapper postgres_mappers.GiftMapper
}

func BuildGiftPostgresRepository() *GiftPostgresRepository {
	return &GiftPostgresRepository{BasePostgresRepository: BuildBasePostgresRepository()}
}

func (r *GiftPostgresRepository) GetGift(id string) (*domain.Gift, error) {
	q := r.getQueryOrTx()

	var model postgres_models.GiftModel
	err := q.Where("id = ?", id).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return r.GiftMapper.ToDomain(model), nil
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
