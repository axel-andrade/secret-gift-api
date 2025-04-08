package bootstrap

import (
	primary_handlers "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/handlers"
	authorized_gift_controllers "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/controllers/authorized_gifts"
	gift_controllers "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/controllers/gifts"
	gift_presenters "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/presenters/gifts"
	postgres_repositories "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/repositories"
	authorize_gift_impl "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/impl/authorized_gifts"
	create_gift_impl "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/impl/gifts"
	authorize_gift "github.com/axel-andrade/secret-gift-api/internal/core/usecases/gifts/authorize"
	create_gift "github.com/axel-andrade/secret-gift-api/internal/core/usecases/gifts/create"
)

type Dependencies struct {
	GiftPostgresRepository           *postgres_repositories.GiftPostgresRepository
	AuthorizedGiftPostgresRepository *postgres_repositories.AuthorizedGiftPostgresRepository

	JsonHandler *primary_handlers.JsonHandler

	CreateGiftPresenter    *gift_presenters.CreateGiftPresenter
	AuthorizeGiftPresenter *gift_presenters.AuthorizeGiftPresenter

	CreateGiftController    *gift_controllers.CreateGiftController
	AuthorizeGiftController *authorized_gift_controllers.AuthorizeGiftController

	CreateGiftUC    *create_gift.CreateGiftUC
	AuthorizeGiftUC *authorize_gift.AuthorizeGiftUC
}

func LoadDependencies() *Dependencies {
	d := &Dependencies{}

	loadRepositories(d)
	loadPrimaryHandlers(d)
	loadUseCases(d)
	loadPresenters(d)
	loadControllers(d)

	return d
}

func loadRepositories(d *Dependencies) {
	d.GiftPostgresRepository = postgres_repositories.BuildGiftPostgresRepository()
	d.AuthorizedGiftPostgresRepository = postgres_repositories.BuildAuthorizedGiftPostgresRepository()
}

func loadPrimaryHandlers(d *Dependencies) {
	d.JsonHandler = primary_handlers.BuildJsonHandler()
}

func loadPresenters(d *Dependencies) {
	d.CreateGiftPresenter = gift_presenters.BuildCreateGiftPresenter()
	d.AuthorizeGiftPresenter = gift_presenters.BuildAuthorizeGiftPresenter()
}

func loadUseCases(d *Dependencies) {
	d.CreateGiftUC = create_gift.BuildCreateGiftUC(create_gift_impl.BuildCreateGiftImpl(d.GiftPostgresRepository))
	d.AuthorizeGiftUC = authorize_gift.BuildAuthorizeGiftUC(authorize_gift_impl.BuildAuthorizeGiftImpl(d.GiftPostgresRepository, d.AuthorizedGiftPostgresRepository))
}

func loadControllers(d *Dependencies) {
	d.CreateGiftController = gift_controllers.BuildSignUpController(d.CreateGiftUC, d.CreateGiftPresenter)
	d.AuthorizeGiftController = authorized_gift_controllers.BuildAuthorizeGiftController(d.AuthorizeGiftUC, d.AuthorizeGiftPresenter)
}
