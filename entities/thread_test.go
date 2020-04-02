package entities

import (
	"reflect"
	"testing"
	"vspro/entities/valueobjects"
)

func TestNewThread(t *testing.T) {
	tID, _ := valueobjects.NewThreadID("")
	bID, _ := valueobjects.NewBulletinBoardID("")
	type args struct {
		ID    ThreadIDer
		bID   BulletinBoardIDer
		title string
	}
	tests := []struct {
		name    string
		args    args
		want    Thread
		wantErr bool
	}{
		{
			name: "オブジェクトが正常に生成される",
			args: args{
				ID:    tID,
				bID:   bID,
				title: "title",
			},
			want: Thread{
				ID:              tID,
				BulletinBoardID: bID,
				Title:           "title",
			},
			wantErr: false,
		},
		{
			name: "textの値が1byte未満だった場合、エラーとなる",
			args: args{
				ID:    tID,
				bID:   bID,
				title: "",
			},
			want:    Thread{},
			wantErr: true,
		},
		{
			name: "textの値が50byteより大きかった場合、エラーとなる",
			args: args{
				ID:  tID,
				bID: bID,
				// 51byte
				title: "aabcdefjhijabcdefjhijabcdefjhijabcdefjhijabcdefjhij",
			},
			want:    Thread{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewThread(tt.args.ID, tt.args.bID, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewThread() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewThread() got = %v, want %v", got, tt.want)
			}
		})
	}
}
