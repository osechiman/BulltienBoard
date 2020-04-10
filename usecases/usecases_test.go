package usecases

import (
	"os"
	"testing"
	"vspro/adapters/gateways"
	"vspro/entities"
	"vspro/entities/valueobjects"
)

var testBulletinBoard TestBulletinBoard

type TestBulletinBoard struct {
	repository BulletinBoardRepositorer
	bid        valueobjects.BulletinBoardID
	title      string
}

var testThread TestThread

type TestThread struct {
	repository ThreadRepositorer
	tid        valueobjects.ThreadID
	bid        valueobjects.BulletinBoardID
	title      string
}

var testComment TestComment

type TestComment struct {
	repository  CommentRepositorer
	tRepository ThreadRepositorer
	cid         valueobjects.CommentID
	tid         valueobjects.ThreadID
	ct          valueobjects.CommentTime
	text        string
}

//TODO:: 事前にデータを用意するのは無理だと判断したのでデータ登録部分は各テストケース無いで行う
func TestMain(m *testing.M) {
	SetUp()
	ec := m.Run()
	os.Exit(ec)
}

// SetUp はテスト実行時にパッケージ全体で利用したい値を準備します。
func SetUp() {
	SetUpBulletinBoardTest()
	SetUpThreadTest()
	SetUpCommentTest()
}

// SetUpBulletinBoardTest はBulletinBoardのテストに関連する値を準備します。
func SetUpBulletinBoardTest() {
	r := gateways.GetInMemoryRepositoryInstance()
	bid, _ := valueobjects.NewBulletinBoardID("")
	bt := "bulletin board title"
	testBulletinBoard = TestBulletinBoard{
		repository: r,
		bid:        bid,
		title:      bt,
	}
}

// SetUpThreadTest はThreadTestのテストに関連する値を準備します。
func SetUpThreadTest() {
	r := gateways.GetInMemoryRepositoryInstance()
	tid, _ := valueobjects.NewThreadID("")
	title := "thread title"
	t, _ := entities.NewThread(tid, testBulletinBoard.bid, title)
	r.AddThread(t)

	tid2, _ := valueobjects.NewThreadID("")
	title2 := "thread title2"
	t2, _ := entities.NewThread(tid2, testBulletinBoard.bid, title2)
	r.AddThread(t2)

	testThread = TestThread{
		repository: r,
		tid:        tid,
		bid:        testBulletinBoard.bid,
		title:      title,
	}
}

// SetUpCommentTest はCommentのテストに関連する値を準備します。
func SetUpCommentTest() {
	r := gateways.GetInMemoryRepositoryInstance()
	cid, _ := valueobjects.NewCommentID("")
	ct, _ := valueobjects.NewCommentTime(-1)
	text := "comment"

	testComment = TestComment{
		repository:  r,
		tRepository: r,
		cid:         cid,
		tid:         testThread.tid,
		ct:          ct,
		text:        text,
	}
}
