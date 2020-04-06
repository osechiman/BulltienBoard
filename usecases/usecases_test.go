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
	bb         entities.BulletinBoard
	bb2        entities.BulletinBoard
}

var testThread TestThread

type TestThread struct {
	repository ThreadRepositorer
	tid        valueobjects.ThreadID
	bid        valueobjects.BulletinBoardID
	title      string
	t          entities.Thread
	t2         entities.Thread
}

var testComment TestComment

type TestComment struct {
	repository  CommentRepositorer
	tRepository ThreadRepositorer
	cid         valueobjects.CommentID
	tid         valueobjects.ThreadID
	ct          valueobjects.CommentTime
	text        string
	c           entities.Comment
	c2          entities.Comment
}

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
	bb, _ := entities.NewBulletinBoard(bid, bt)
	r.AddBulletinBoard(bb)

	bid2, _ := valueobjects.NewBulletinBoardID("")
	bt2 := "bulletin board title2"
	bb2, _ := entities.NewBulletinBoard(bid2, bt2)
	r.AddBulletinBoard(bb2)

	testBulletinBoard = TestBulletinBoard{
		repository: r,
		bid:        bid,
		title:      bt,
		bb:         bb,
		bb2:        bb2,
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
		t:          t,
		t2:         t2,
	}
}

// SetUpCommentTest はCommentのテストに関連する値を準備します。
func SetUpCommentTest() {
	r := gateways.GetInMemoryRepositoryInstance()
	cid, _ := valueobjects.NewCommentID("")
	ct, _ := valueobjects.NewCommentTime(-1)
	text := "comment"
	c, _ := entities.NewComment(cid, testThread.tid, text, ct)
	r.AddComment(c)

	cid2, _ := valueobjects.NewCommentID("")
	ct2, _ := valueobjects.NewCommentTime(-1)
	text2 := "comment2"
	c2, _ := entities.NewComment(cid2, testThread.tid, text2, ct2)
	r.AddComment(c2)

	testComment = TestComment{
		repository:  r,
		tRepository: r,
		cid:         cid,
		tid:         testThread.tid,
		ct:          ct,
		text:        text,
		c:           c,
	}
}
