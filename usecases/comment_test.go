package usecases

import (
	"reflect"
	"testing"
	"vspro/entities"
	"vspro/entities/valueobjects"
)

func TestCommentUsecase_AddComment(t *testing.T) {
	type fields struct {
		Repository CommentRepositorer
	}
	type args struct {
		c                entities.Comment
		threadRepository ThreadRepositorer
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
				Repository: testComment.repository,
			},
			args: args{
				c:                testComment.c,
				threadRepository: testComment.tRepository,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CommentUsecase{
				Repository: tt.fields.Repository,
			}
			if err := cc.AddComment(tt.args.c, tt.args.threadRepository); (err != nil) != tt.wantErr {
				t.Errorf("AddComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommentUsecase_ListComment(t *testing.T) {
	type fields struct {
		Repository CommentRepositorer
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entities.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CommentUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := cc.ListComment()
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
	type fields struct {
		Repository CommentRepositorer
	}
	type args struct {
		tID valueobjects.ThreadID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entities.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := &CommentUsecase{
				Repository: tt.fields.Repository,
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
	type args struct {
		r CommentRepositorer
	}
	tests := []struct {
		name string
		args args
		want *CommentUsecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommentUsecase(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
