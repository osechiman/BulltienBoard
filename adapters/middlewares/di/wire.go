// +build wireinject

package di

import (
	"vspro/adapters/gateways"
	"vspro/usecases"

	"github.com/google/wire"
)

func GetBulletinBoardUsecase() *usecases.BulletinBoardUsecase {
	wire.Build(
		usecases.NewBulletinBoardUsecase,
		gateways.NewInMemoryRepository,
		wire.Bind(new(usecases.BulletinBoardRepositorer), new(*gateways.InMemoryRepository)),
		wire.Bind(new(usecases.ThreadRepositorer), new(*gateways.InMemoryRepository)),
	)
	return &usecases.BulletinBoardUsecase{}
}

func GetThreadUsecase() *usecases.ThreadUsecase {
	wire.Build(
		usecases.NewThreadUsecase,
		gateways.NewInMemoryRepository,
		wire.Bind(new(usecases.ThreadRepositorer), new(*gateways.InMemoryRepository)),
		wire.Bind(new(usecases.BulletinBoardRepositorer), new(*gateways.InMemoryRepository)),
		wire.Bind(new(usecases.CommentRepositorer), new(*gateways.InMemoryRepository)),
	)
	return &usecases.ThreadUsecase{}
}

func GetCommentUsecase() *usecases.CommentUsecase {
	wire.Build(
		usecases.NewCommentUsecase,
		gateways.NewInMemoryRepository,
		wire.Bind(new(usecases.CommentRepositorer), new(*gateways.InMemoryRepository)),
		wire.Bind(new(usecases.ThreadRepositorer), new(*gateways.InMemoryRepository)),
	)
	return &usecases.CommentUsecase{}
}
