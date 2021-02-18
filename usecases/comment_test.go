package usecases

import (
	"bulltienboard/adapters/gateways"
	"bulltienboard/entities"
	"bulltienboard/entities/valueobjects"
	"reflect"
	"sort"
	"testing"
)

func TestCommentUsecase_AddComment(t *testing.T) {
	repository := gateways.GetInMemoryRepositoryInstance()
	repository.DeleteAll()

	bid, _ := valueobjects.NewBulletinBoardID("")
	b, _ := entities.NewBulletinBoard(bid, "bulletin board title")

	repository.AddBulletinBoard(b)

	tid, _ := valueobjects.NewThreadID("")
	th, _ := entities.NewThread(tid, bid, "thread title")

	tid1, _ := valueobjects.NewThreadID("")

	repository.AddThread(th)

	cid, _ := valueobjects.NewCommentID("")
	ct, _ := valueobjects.NewCommentTime(-1)
	c := entities.Comment{
		ID:       cid,
		ThreadID: tid,
		Text:     "comment",
		CreateAt: ct,
	}

	type fields struct {
		CommentRepository CommentRepositorer
		ThreadRepository  ThreadRepositorer
	}
	type args struct {
		c entities.Comment
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "エンティティの登録が正常に出来る",
			fields: fields{
				CommentRepository: repository,
				ThreadRepository:  repository,
			},
			args: args{
				c: c,
			},
			wantErr: false,
		},
		{
			name: "entities.Commentに指定するThreadIDがRepositoryに存在しない値だった場合、エラーが返却される",
			fields: fields{
				CommentRepository: repository,
				ThreadRepository:  repository,
			},
			args: args{
				c: entities.Comment{
					ID:       cid,
					ThreadID: tid1,
					Text:     "comment",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CommentUsecase{
				CommentRepository: tt.fields.CommentRepository,
				ThreadRepository:  tt.fields.ThreadRepository,
			}
			if err := cc.AddComment(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AddComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	t.Run("Commentの登録数がCommentLimitを超えて登録された場合、エラーが返却される", func(t *testing.T) {
		repository := gateways.GetInMemoryRepositoryInstance()
		repository.DeleteAll()

		cc := &CommentUsecase{
			CommentRepository: repository,
			ThreadRepository:  repository,
		}
		bid, _ := valueobjects.NewBulletinBoardID("")
		b, _ := entities.NewBulletinBoard(bid, "bulletin board title")

		repository.AddBulletinBoard(b)

		tid, _ := valueobjects.NewThreadID("")
		th, _ := entities.NewThread(tid, bid, "thread title")

		repository.AddThread(th)

		// 上限値までCommentを登録する
		for i := 0; i < CommentLimit; i++ {
			cid, _ := valueobjects.NewCommentID("")
			ct, _ := valueobjects.NewCommentTime(-1)
			c, _ := entities.NewComment(cid, tid, "comment", ct)
			repository.AddComment(c)
		}

		// 上限値以上に登録する為のComment生成
		cid, _ := valueobjects.NewCommentID("")
		ct, _ := valueobjects.NewCommentTime(-1)
		c, _ := entities.NewComment(cid, tid, "last comment", ct)

		wantErr := true
		if err := cc.AddComment(c); (err != nil) != wantErr {
			t.Errorf("AddComment() error = %v, wantErr %v", err, wantErr)
		}
	})
}

func TestCommentUsecase_ListComment(t *testing.T) {
	repository := gateways.GetInMemoryRepositoryInstance()
	repository.DeleteAll()

	bid, _ := valueobjects.NewBulletinBoardID("")
	b, _ := entities.NewBulletinBoard(bid, "bulletin board title")

	repository.AddBulletinBoard(b)

	tid, _ := valueobjects.NewThreadID("")
	th, _ := entities.NewThread(tid, bid, "thread title")

	repository.AddThread(th)

	cid, _ := valueobjects.NewCommentID("")
	ct, _ := valueobjects.NewCommentTime(-1)
	c, _ := entities.NewComment(cid, tid, "comment", ct)

	repository.AddComment(c)

	cid1, _ := valueobjects.NewCommentID("")
	ct1, _ := valueobjects.NewCommentTime(-1)
	c1, _ := entities.NewComment(cid1, tid, "comment", ct1)

	repository.AddComment(c1)

	want := append([]entities.Comment{}, c, c1)

	type fields struct {
		CommentRepository CommentRepositorer
		ThreadRepository  ThreadRepositorer
	}
	tests := []struct {
		name    string
		fields  fields
		want    []entities.Comment
		wantErr bool
	}{
		{
			name: "[]entities.Commentが取得出来る",
			fields: fields{
				CommentRepository: repository,
				ThreadRepository:  repository,
			},
			want:    want,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CommentUsecase{
				CommentRepository: tt.fields.CommentRepository,
				ThreadRepository:  tt.fields.ThreadRepository,
			}
			got, err := cc.ListComment()

			// Sliceの順序はソートせずに返却する仕様なので、テスト時には一度ソートをして値が等価であるかを検証します。
			sort.Slice(got, func(i, j int) bool {
				return got[i].ID.String() < got[j].ID.String()
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].ID.String() < tt.want[j].ID.String()
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("ListComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListComment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentUsecase_ListCommentByThreadID(t *testing.T) {
	repository := gateways.GetInMemoryRepositoryInstance()
	repository.DeleteAll()

	bid, _ := valueobjects.NewBulletinBoardID("")
	b, _ := entities.NewBulletinBoard(bid, "bulletin board title")

	repository.AddBulletinBoard(b)

	tid, _ := valueobjects.NewThreadID("")
	th, _ := entities.NewThread(tid, bid, "thread title")

	repository.AddThread(th)

	tid1, _ := valueobjects.NewThreadID("")
	th1, _ := entities.NewThread(tid1, bid, "thread1 title")

	repository.AddThread(th1)

	cid, _ := valueobjects.NewCommentID("")
	ct, _ := valueobjects.NewCommentTime(-1)
	c, _ := entities.NewComment(cid, tid, "comment", ct)

	repository.AddComment(c)

	cs := append([]entities.Comment{}, c)

	type fields struct {
		CommentRepository CommentRepositorer
		ThreadRepository  ThreadRepositorer
	}
	type args struct {
		tID valueobjects.ThreadID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Comment
		wantErr bool
	}{
		{
			name: "指定するThreadIDに紐づくCommentのみが取得出来る",
			fields: fields{
				CommentRepository: repository,
				ThreadRepository:  repository,
			},
			args: args{
				tID: tid,
			},
			want:    cs,
			wantErr: false,
		},
		{
			name: "指定するThreadIDに紐づくCommentが存在しない場合、エラーが返却される",
			fields: fields{
				CommentRepository: repository,
				ThreadRepository:  repository,
			},
			args: args{
				tID: tid1,
			},
			want:    []entities.Comment{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CommentUsecase{
				CommentRepository: tt.fields.CommentRepository,
				ThreadRepository:  tt.fields.ThreadRepository,
			}
			got, err := cc.ListCommentByThreadID(tt.args.tID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListCommentByThreadID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListCommentByThreadID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCommentUsecase(t *testing.T) {
	repository := gateways.GetInMemoryRepositoryInstance()
	repository.DeleteAll()

	type args struct {
		CommentRepository CommentRepositorer
		ThreadRepository  ThreadRepositorer
	}
	tests := []struct {
		name string
		args args
		want *CommentUsecase
	}{
		{
			name: "オブジェクトが正常に生成される",
			args: args{
				CommentRepository: repository,
				ThreadRepository:  repository,
			},
			want: &CommentUsecase{
				CommentRepository: repository,
				ThreadRepository:  repository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommentUsecase(tt.args.CommentRepository, tt.args.ThreadRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
