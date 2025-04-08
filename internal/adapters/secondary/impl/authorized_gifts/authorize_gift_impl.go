package auth_gift_impl

import (
	postgres_repositories "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/repositories"
	authorize_gift "github.com/axel-andrade/secret-gift-api/internal/core/usecases/gifts/authorize"
)

func BuildAuthorizeGiftImpl(gr *postgres_repositories.GiftPostgresRepository, ar *postgres_repositories.AuthorizedGiftPostgresRepository) authorize_gift.AuthorizeGiftGateway {
	return &struct {
		*postgres_repositories.GiftPostgresRepository
		*postgres_repositories.AuthorizedGiftPostgresRepository
	}{
		GiftPostgresRepository:           gr,
		AuthorizedGiftPostgresRepository: ar,
	}
}
