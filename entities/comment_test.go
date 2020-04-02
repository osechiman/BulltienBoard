package entities

import (
	"reflect"
	"testing"
	"vspro/entities/valueobjects"
)

func TestNewComment(t *testing.T) {
	cID, _ := valueobjects.NewCommentID("")
	tID, _ := valueobjects.NewThreadID("")
	ct, _ := valueobjects.NewCommentTime(-1)
	type args struct {
		ID    CommentIDer
		tID   ThreadIDer
		text  string
		cTime CommentTimer
	}
	tests := []struct {
		name    string
		args    args
		want    Comment
		wantErr bool
	}{
		{
			name: "オブジェクトが正常に生成される",
			args: args{
				ID:    cID,
				tID:   tID,
				text:  "comment",
				cTime: ct,
			},
			want: Comment{
				ID:       cID,
				ThreadID: tID,
				Text:     "comment",
				CreateAt: ct,
			},
			wantErr: false,
		},
		{
			name: "textの値が1byte未満だった場合、エラーとなる",
			args: args{
				ID:    cID,
				tID:   tID,
				text:  "",
				cTime: ct,
			},
			want:    Comment{},
			wantErr: true,
		},
		{
			name: "textの値が2048byteより大きかった場合、エラーとなる",
			args: args{
				ID:  cID,
				tID: tID,
				// 3000byte
				text:  "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghij",
				cTime: ct,
			},
			want:    Comment{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewComment(tt.args.ID, tt.args.tID, tt.args.text, tt.args.cTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComment() got = %v, want %v", got, tt.want)
			}
		})
	}
}
