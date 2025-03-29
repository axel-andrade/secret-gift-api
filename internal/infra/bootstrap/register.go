package bootstrap

import (
	primary_handlers "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/handlers"
	gift_controllers "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/controllers/gifts"
	gift_presenters "github.com/axel-andrade/secret-gift-api/internal/adapters/primary/http/presenters/gifts"
	postgres_repositories "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/database/postgres/repositories"
	create_gift_impl "github.com/axel-andrade/secret-gift-api/internal/adapters/secondary/impl/gifts"
	create_gift "github.com/axel-andrade/secret-gift-api/internal/core/usecases/gifts/create"
)

type Dependencies struct {
	GiftPostgresRepository *postgres_repositories.GiftPostgresRepository

	JsonHandler *primary_handlers.JsonHandler

	CreateGiftPresenter *gift_presenters.CreateGiftPresenter

	CreateGiftController *gift_controllers.CreateGiftController

	CreateGiftUC *create_gift.CreateGiftUC
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
}

func loadPrimaryHandlers(d *Dependencies) {
	d.JsonHandler = primary_handlers.BuildJsonHandler()
}

func loadPresenters(d *Dependencies) {
	d.CreateGiftPresenter = gift_presenters.BuildCreateGiftPresenter()
}

func loadUseCases(d *Dependencies) {
	d.CreateGiftUC = create_gift.BuildCreateGiftUC(create_gift_impl.BuildCreateGiftImpl(d.GiftPostgresRepository))
}

func loadControllers(d *Dependencies) {

	d.CreateGiftController = gift_controllers.BuildSignUpController(d.CreateGiftUC, d.CreateGiftPresenter)
}
