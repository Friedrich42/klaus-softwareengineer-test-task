package main

import (
	"cipactonal/internal/database"
	"cipactonal/internal/repository"
	"cipactonal/internal/router"
	"cipactonal/internal/usecase"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			fx.Annotate(database.NewSqlite, fx.As(new(database.Database))),
			fx.Annotate(repository.NewRepository, fx.As(new(usecase.Repository))),
			fx.Annotate(usecase.NewHandler, fx.As(new(router.UseCase))),
			router.NewHandler,
		),
		fx.Invoke(
			router.RegisterServer,
		),
	).Run()
}
