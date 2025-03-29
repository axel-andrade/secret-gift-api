package create_gift_impl

import (
	postgres_repositories "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/repositories"
	create_gift "github.com/axel-andrade/secret-gift-api/internal/core/usecases/gifts/create"
)

var BuildCreateGiftImpl = func(r *postgres_repositories.GiftPostgresRepository) create_gift.CreateGiftGateway {
	return &struct {
		*postgres_repositories.GiftPostgresRepository
	}{r}
}
