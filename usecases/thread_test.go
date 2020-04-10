package usecases

import (
	"reflect"
	"testing"
	"vspro/entities"
	"vspro/entities/valueobjects"
)

func TestNewThreadUsecase(t *testing.T) {
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
				r: testThread.repository,
			},
			want: &ThreadUsecase{
				Repository: testThread.repository,
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
				Repository: testThread.repository,
			},
			args: args{
				t:                       entities.Thread{},
				bulletinBoardRepository: testBulletinBoard.repository,
			},
			wantErr: false,
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
				Repository: testThread.repository,
			},
			args: args{
				ID:                testThread.tid,
				commentRepository: testComment.repository,
			},
			want:    entities.Thread{},
			wantErr: false,
		},
		{
			name: "BulletinBoardIDが存在しない値だった場合、エラーが返却される",
			fields: fields{
				Repository: testThread.repository,
			},
			args: args{
				ID:                valueobjects.ThreadID{},
				commentRepository: testComment.repository,
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
				Repository: testThread.repository,
			},
			want:    []entities.Thread{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tu := &ThreadUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := tu.ListThread()
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
				Repository: testThread.repository,
			},
			args: args{
				bID: testBulletinBoard.bid,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tu := &ThreadUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := tu.ListThreadByBulletinBoardID(tt.args.bID)
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
