package entities

import (
	"bulltienboard/entities/valueobjects"
	"reflect"
	"testing"
)

func TestNewBulletinBoard(t *testing.T) {
	bID, _ := valueobjects.NewBulletinBoardID("")
	type args struct {
		ID    BulletinBoardIDer
		title string
	}
	tests := []struct {
		name    string
		args    args
		want    BulletinBoard
		wantErr bool
	}{
		{
			name: "オブジェクトが正常に生成される",
			args: args{
				ID:    bID,
				title: "title",
			},
			want: BulletinBoard{
				ID:    bID,
				Title: "title",
			},
			wantErr: false,
		},
		{
			name: "textの値が1byte未満だった場合、エラーとなる",
			args: args{
				ID:    bID,
				title: "",
			},
			want:    BulletinBoard{},
			wantErr: true,
		},
		{
			name: "textの値が50byteより大きかった場合、エラーとなる",
			args: args{
				ID: bID,
				// 51byte
				title: "aabcdefjhijabcdefjhijabcdefjhijabcdefjhijabcdefjhij",
			},
			want:    BulletinBoard{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBulletinBoard(tt.args.ID, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBulletinBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBulletinBoard() got = %v, want %v", got, tt.want)
			}
		})
	}
}
