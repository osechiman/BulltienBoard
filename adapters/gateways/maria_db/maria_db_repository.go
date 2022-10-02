package gateways

import (
	"bulltienboard/adapters/gateways/maria_db/models"
	"bulltienboard/entities"
	"bulltienboard/entities/errorobjects"
	"bulltienboard/entities/valueobjects"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const DRIVER_NAME = "mysql"

// MariaDBRepository は MariaDB上でデータ管理する為のStructです。
type MariaDBRepository struct {
	host     string
	port     int
	db_name  string
	user     string
	password string
	db       *sql.DB
	ctx      context.Context
}

// NewMariaDBRepository MariaDBRepositoryを初期化します。
func NewMariaDBRepository(host string, port int, db_name string, user string, password string) (*MariaDBRepository, error) {
	var mdr MariaDBRepository
	// 接続したいDBのhost名もしくはIPアドレス
	mdr.host = host
	// 接続時のport番号
	mdr.port = port
	// 接続時のデータベース名
	mdr.db_name = db_name
	// 接続時のユーザー名
	mdr.user = user
	// 接続時のパスワード
	mdr.password = password
	mdr.ctx = context.Background()

	db, err := sql.Open(DRIVER_NAME, mdr.connection())
	if err != nil {
		return nil, errorobjects.NewDatabaseConnectionError(err.Error())
	}
	// defer db.Close()
	mdr.db = db
	return &mdr, nil
}

// connection はDBへ接続するために必要な情報を文字列に結合して返却します。
func (mdr *MariaDBRepository) connection() string {
	cst := "%s:%s@tcp(%s:%d)/%s"
	// TODO:: mysqlのconfigセットが提供されているのでそれを使う
	return fmt.Sprintf(cst, mdr.user, mdr.password, mdr.host, mdr.port, mdr.db_name)
}

// GetBulletinBoardByID は指定されたvalueobjects.BulletinBoardIDを持つentities.BulletinBoardを取得します。
func (mdr *MariaDBRepository) GetBulletinBoardByID(ID valueobjects.BulletinBoardID) (entities.BulletinBoard, error) {
	bb, err := models.BulletinBoards().All(mdr.ctx, mdr.db)
	if err != nil {
		return entities.BulletinBoard{}, err
	}
	fmt.Printf("%v", bb)
	//rows, err := mdr.db.Query("SHOW DATABASES;")
	//if err != nil {
	//	fmt.Printf("%v", err)
	//	return entities.BulletinBoard{}, err
	//}
	//fmt.Printf("show databases: %v", rows)
	return entities.BulletinBoard{}, nil
}

// ListBulletinBoard はentities.BulletinBoardの一覧を取得します。
func (mdr *MariaDBRepository) ListBulletinBoard() ([]entities.BulletinBoard, error) {
	panic("not implemented") // TODO: Implement
}

// AddBulletinBoard はentities.BulletinBoardを追加します。
func (mdr *MariaDBRepository) AddBulletinBoard(bb entities.BulletinBoard) error {
	panic("not implemented") // TODO: Implement
}

// DeleteBulletinBoard はentities.BulletinBoardを全て削除します。
func (mdr *MariaDBRepository) DeleteBulletinBoard() error {
	panic("not implemented") // TODO: Implement
}

// GetThreadByID は指定されたvalueobjects.ThreadIDを持つentities.Threadを取得します。
func (mdr *MariaDBRepository) GetThreadByID(ID valueobjects.ThreadID) (entities.Thread, error) {
	panic("not implemented") // TODO: Implement
}

// ListThread はentities.Threadの一覧を取得します。
func (mdr *MariaDBRepository) ListThread() ([]entities.Thread, error) {
	panic("not implemented") // TODO: Implement
}

// ListThreadByBulletinBoardID は指定されたvalueobjects.BulletinBoardIDを持つentities.Threadの一覧を取得します。
func (mdr *MariaDBRepository) ListThreadByBulletinBoardID(bID valueobjects.BulletinBoardID) ([]entities.Thread, error) {
	panic("not implemented") // TODO: Implement
}

// AddThread はentities.Threadを追加します。
func (mdr *MariaDBRepository) AddThread(t entities.Thread) error {
	panic("not implemented") // TODO: Implement
}

// ListComment はentities.Commentの一覧を取得します。
func (mdr *MariaDBRepository) ListComment() ([]entities.Comment, error) {
	panic("not implemented") // TODO: Implement
}

// ListCommentByThreadID は指定されたvalueobjects.ThreadIDを持つentities.Commentの一覧を取得します。
func (mdr *MariaDBRepository) ListCommentByThreadID(tID valueobjects.ThreadID) ([]entities.Comment, error) {
	panic("not implemented") // TODO: Implement
}

// AddComment はentities.Comment を追加します。
func (mdr *MariaDBRepository) AddComment(c entities.Comment) error {
	panic("not implemented") // TODO: Implement
}
