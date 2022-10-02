//go:build wireinject
// +build wireinject

package di

import (
	"bulltienboard/adapters/controllers"
	"bulltienboard/adapters/gateways"
	"bulltienboard/adapters/presenters"
	"bulltienboard/drivers/web/api"
	"bulltienboard/usecases"

	"github.com/google/wire"
)

func ProvideMariaDBRepository() *gateways.MariaDBRepository {
	md, _ := gateways.NewMariaDBRepository(
		"localhost",
		3306,
		"BulltienBoard",
		"root",
		"my-secret-pw",
	)
	return md
}

func InitializeRouter() *api.Router {
	// api.NewRouterで生成する値に必要なプロバイダ(コンストラクタ)を全て列挙します。
	// wire.Bindでは列挙したプロバイダ(コンストラクタ)がインターフェースを要求している時に実態として何を渡すか定義しています。
	// 第一引数はinterfaceです。
	// 第二引数は第一引数のinterfaceを満たした実態を渡しています。
	// この時第二引数の*gateways.InMemoryRepositoryはgateways.NewInMemoryRepositoryが実行されるように依存解決されます。
	wire.Build(
		api.NewRouter,
		controllers.NewBulletinBoardController,
		controllers.NewThreadController,
		controllers.NewCommentController,
		presenters.NewBulletinBoardPresenter,
		presenters.NewThreadPresenter,
		presenters.NewCommentPresenter,
		presenters.NewErrorPresenter,
		usecases.NewBulletinBoardUsecase,
		usecases.NewThreadUsecase,
		usecases.NewCommentUsecase,
		//gateways.NewInMemoryRepository,
		ProvideMariaDBRepository,
		// wire.Bind(new(usecases.BulletinBoardRepositorer), new(*gateways.InMemoryRepository)),
		wire.Bind(new(usecases.BulletinBoardRepositorer), ProvideMariaDBRepository),
		// wire.Bind(new(usecases.ThreadRepositorer), new(*gateways.InMemoryRepository)),
		wire.Bind(new(usecases.ThreadRepositorer), ProvideMariaDBRepository),
		// wire.Bind(new(usecases.CommentRepositorer), new(*gateways.InMemoryRepository)),
		wire.Bind(new(usecases.CommentRepositorer), ProvideMariaDBRepository),
	)
	return &api.Router{}
}
