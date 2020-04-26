package usecases

import (
	"reflect"
	"sort"
	"testing"
	"vspro/adapters/gateways"
	"vspro/entities"
	"vspro/entities/valueobjects"
)

func TestNewThreadUsecase(t *testing.T) {
	repository := gateways.GetInMemoryRepositoryInstance()
	repository.DeleteAll()

	type args struct {
		r ThreadRepositorer
	}
	tests := []struct {
		name string
		args args
		want *ThreadUsecase
	}{
		{
			name: "オブジェクトが正常に生成される",
			args: args{
				r: repository,
			},
			want: &ThreadUsecase{
				Repository: repository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewThreadUsecase(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewThreadUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreadUsecase_AddThread(t *testing.T) {
	repository := gateways.GetInMemoryRepositoryInstance()
	repository.DeleteAll()

	bid, _ := valueobjects.NewBulletinBoardID("")
	b, _ := entities.NewBulletinBoard(bid, "bulletin board title")

	bid1, _ := valueobjects.NewBulletinBoardID("")

	repository.AddBulletinBoard(b)

	tid, _ := valueobjects.NewThreadID("")

	type fields struct {
		Repository ThreadRepositorer
	}
	type args struct {
		t                       entities.Thread
		bulletinBoardRepository BulletinBoardRepositorer
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
				Repository: repository,
			},
			args: args{
				t: entities.Thread{
					ID:              tid,
					BulletinBoardID: bid,
					Title:           "thread",
				},
				bulletinBoardRepository: repository,
			},
			wantErr: false,
		},
		{
			name: "entities.Threadに指定するBulletinBoardIDがRepositoryに存在しない値だった場合、エラーが返却される",
			fields: fields{
				Repository: repository,
			},
			args: args{
				t: entities.Thread{
					ID:              tid,
					BulletinBoardID: bid1,
					Title:           "thread",
				},
				bulletinBoardRepository: repository,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tu := &ThreadUsecase{
				Repository: tt.fields.Repository,
			}
			if err := tu.AddThread(tt.args.t, tt.args.bulletinBoardRepository); (err != nil) != tt.wantErr {
				t.Errorf("AddThread() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestThreadUsecase_GetThreadByID(t *testing.T) {
	repository := gateways.GetInMemoryRepositoryInstance()
	repository.DeleteAll()

	bid, _ := valueobjects.NewBulletinBoardID("")
	b, _ := entities.NewBulletinBoard(bid, "bulletin board title")

	repository.AddBulletinBoard(b)

	tid, _ := valueobjects.NewThreadID("")
	th, _ := entities.NewThread(tid, bid, "thread")

	repository.AddThread(th)

	tid1, _ := valueobjects.NewThreadID("")
	th1, _ := entities.NewThread(tid1, bid, "thread2")

	repository.AddThread(th1)

	type fields struct {
		Repository ThreadRepositorer
	}
	type args struct {
		ID                valueobjects.ThreadID
		commentRepository CommentRepositorer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.Thread
		wantErr bool
	}{
		{
			name: "ThreadIDからentities.Threadが取得出来る",
			fields: fields{
				Repository: repository,
			},
			args: args{
				ID:                tid,
				commentRepository: repository,
			},
			want: entities.Thread{
				ID:              tid,
				BulletinBoardID: bid,
				Title:           "thread",
				Comments:        []entities.Comment{},
			},
			wantErr: false,
		},
		{
			name: "ThreadIDが存在しない値だった場合、エラーが返却される",
			fields: fields{
				Repository: repository,
			},
			args: args{
				ID:                valueobjects.ThreadID{},
				commentRepository: repository,
			},
			want:    entities.Thread{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tu := &ThreadUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := tu.GetThreadByID(tt.args.ID, tt.args.commentRepository)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetThreadByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetThreadByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreadUsecase_ListThread(t *testing.T) {
	repository := gateways.GetInMemoryRepositoryInstance()
	repository.DeleteAll()

	bid, _ := valueobjects.NewBulletinBoardID("")
	b, _ := entities.NewBulletinBoard(bid, "bulletin board title")

	repository.AddBulletinBoard(b)

	tid, _ := valueobjects.NewThreadID("")
	th, _ := entities.NewThread(tid, bid, "thread")

	repository.AddThread(th)

	tid1, _ := valueobjects.NewThreadID("")
	th1, _ := entities.NewThread(tid1, bid, "thread1")

	repository.AddThread(th1)

	ths := append([]entities.Thread{}, th, th1)

	type fields struct {
		Repository ThreadRepositorer
	}
	tests := []struct {
		name    string
		fields  fields
		want    []entities.Thread
		wantErr bool
	}{
		{
			name: "[]entities.Threadが取得出来る",
			fields: fields{
				Repository: repository,
			},
			want:    ths,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tu := &ThreadUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := tu.ListThread()

			// Sliceの順序はソートせずに返却する仕様なので、テスト時には一度ソートをして値が等価であるかを検証します。
			sort.Slice(got, func(i, j int) bool {
				return got[i].ID.String() < got[j].ID.String()
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].ID.String() < tt.want[j].ID.String()
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("ListThread() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListThread() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreadUsecase_ListThreadByBulletinBoardID(t *testing.T) {
	repository := gateways.GetInMemoryRepositoryInstance()
	repository.DeleteAll()

	bid, _ := valueobjects.NewBulletinBoardID("")
	b, _ := entities.NewBulletinBoard(bid, "bulletin board title")

	repository.AddBulletinBoard(b)

	bid1, _ := valueobjects.NewBulletinBoardID("")
	b1, _ := entities.NewBulletinBoard(bid1, "bulletin board1 title")

	repository.AddBulletinBoard(b1)

	bid2, _ := valueobjects.NewBulletinBoardID("")
	b2, _ := entities.NewBulletinBoard(bid2, "bulletin board2 title")

	repository.AddBulletinBoard(b2)

	tid, _ := valueobjects.NewThreadID("")
	th, _ := entities.NewThread(tid, bid2, "thread")

	repository.AddThread(th)

	tid1, _ := valueobjects.NewThreadID("")
	th1, _ := entities.NewThread(tid1, bid, "thread")

	repository.AddThread(th1)

	tid2, _ := valueobjects.NewThreadID("")
	th2, _ := entities.NewThread(tid2, bid, "thread2")

	repository.AddThread(th2)

	ths := append([]entities.Thread{}, th1, th2)

	type fields struct {
		Repository ThreadRepositorer
	}
	type args struct {
		bID valueobjects.BulletinBoardID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Thread
		wantErr bool
	}{
		{
			name: "BulletinBoardIDから[]entities.Threadが取得出来る",
			fields: fields{
				Repository: repository,
			},
			args: args{
				bID: bid,
			},
			want:    ths,
			wantErr: false,
		},
		{
			name: "指定するBulletinBoardIDに紐づくThreadが存在しない場合、エラーが返却される",
			fields: fields{
				Repository: repository,
			},
			args: args{
				bID: bid1,
			},
			want:    []entities.Thread{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tu := &ThreadUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := tu.ListThreadByBulletinBoardID(tt.args.bID)

			// Sliceの順序はソートせずに返却する仕様なので、テスト時には一度ソートをして値が等価であるかを検証します。
			sort.Slice(got, func(i, j int) bool {
				return got[i].ID.String() < got[j].ID.String()
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].ID.String() < tt.want[j].ID.String()
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("ListThreadByBulletinBoardID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListThreadByBulletinBoardID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
