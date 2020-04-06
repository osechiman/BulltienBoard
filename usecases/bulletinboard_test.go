package usecases

import (
	"reflect"
	"testing"
	"vspro/entities"
	"vspro/entities/valueobjects"
)

func TestBulletinBoardUsecase_AddBulletinBoard(t *testing.T) {
	type fields struct {
		Repository BulletinBoardRepositorer
	}
	type args struct {
		bb entities.BulletinBoard
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
				Repository: testBulletinBoard.repository,
			},
			args: args{
				bb: testBulletinBoard.bb,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bbu := &BulletinBoardUsecase{
				Repository: tt.fields.Repository,
			}
			if err := bbu.AddBulletinBoard(tt.args.bb); (err != nil) != tt.wantErr {
				t.Errorf("AddBulletinBoard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBulletinBoardUsecase_GetBulletinBoardByID(t *testing.T) {
	bb := testBulletinBoard.bb
	bb.Threads = append(bb.Threads, testThread.t, testThread.t2)
	bid, _ := valueobjects.NewBulletinBoardID("")

	type fields struct {
		Repository BulletinBoardRepositorer
	}
	type args struct {
		ID               valueobjects.BulletinBoardID
		threadRepository ThreadRepositorer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.BulletinBoard
		wantErr bool
	}{
		{
			name: "BulletinBoardIDからentities.BulletinBoardが取得出来る",
			fields: fields{
				Repository: testBulletinBoard.repository,
			},
			args: args{
				ID:               testBulletinBoard.bid,
				threadRepository: testThread.repository,
			},
			want:    bb,
			wantErr: false,
		},
		{
			name: "BulletinBoardIDが存在しない値だった場合、エラーが返却される",
			fields: fields{
				Repository: testBulletinBoard.repository,
			},
			args: args{
				ID:               bid,
				threadRepository: testThread.repository,
			},
			want:    entities.BulletinBoard{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bbu := &BulletinBoardUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := bbu.GetBulletinBoardByID(tt.args.ID, tt.args.threadRepository)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBulletinBoardByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBulletinBoardByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBulletinBoardUsecase_ListBulletinBoard(t *testing.T) {
	bbs := make([]entities.BulletinBoard, 0)
	bbs = append(bbs, testBulletinBoard.bb, testBulletinBoard.bb2)
	type fields struct {
		Repository BulletinBoardRepositorer
	}
	tests := []struct {
		name    string
		fields  fields
		want    []entities.BulletinBoard
		wantErr bool
	}{
		{
			name: "[]entities.BulletinBoardが取得出来る",
			fields: fields{
				Repository: testBulletinBoard.repository,
			},
			want:    bbs,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bbu := &BulletinBoardUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := bbu.ListBulletinBoard()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListBulletinBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListBulletinBoard() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBulletinBoardUsecase(t *testing.T) {
	type args struct {
		r BulletinBoardRepositorer
	}
	tests := []struct {
		name string
		args args
		want *BulletinBoardUsecase
	}{
		{
			name: "オブジェクトが正常に生成される",
			args: args{
				r: testBulletinBoard.repository,
			},
			want: &BulletinBoardUsecase{
				Repository: testBulletinBoard.repository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBulletinBoardUsecase(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBulletinBoardUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
