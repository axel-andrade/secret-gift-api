package postgres_repositories

import (
	postgres_mappers "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/mappers"
	postgres_models "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/models"
	"github.com/axel-andrade/secret-gift-api/internal/core/domain"
	"gorm.io/gorm"
)

type AuthorizedGiftPostgresRepository struct {
	*BasePostgresRepository
	AuthorizedGiftMapper postgres_mappers.AuthorizedGiftMapper
}

func BuildAuthorizedGiftPostgresRepository() *AuthorizedGiftPostgresRepository {
	return &AuthorizedGiftPostgresRepository{BasePostgresRepository: BuildBasePostgresRepository()}
}

func (r *AuthorizedGiftPostgresRepository) GetAuthorizedGift(giftId string) (*domain.AuthorizedGift, error) {
	q := r.getQueryOrTx()

	var model postgres_models.AuthorizedGiftModel
	err := q.Where("gift_id = ?", giftId).First(&model).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return r.AuthorizedGiftMapper.ToDomain(model), nil
}

func (r *AuthorizedGiftPostgresRepository) CreateAuthorizedGift(ag *domain.AuthorizedGift) (*domain.AuthorizedGift, error) {
	model := r.AuthorizedGiftMapper.ToPersistence(*ag)

	q := r.getQueryOrTx()

	err := q.Create(model).Error

	if err != nil {
		return nil, err
	}

	return r.AuthorizedGiftMapper.ToDomain(*model), nil
}
