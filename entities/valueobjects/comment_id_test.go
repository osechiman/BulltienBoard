package valueobjects

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestCommentID_Equals(t *testing.T) {
	uid, _ := uuid.NewRandom()
	cid, _ := NewCommentID(uid.String())
	type fields struct {
		id  commentID
		str string
	}
	type args struct {
		other CommentID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "引数に渡したCommentIDが自分自身のオブジェクトと同一の場合にtrueが返却される",
			fields: fields{
				id:  commentID(uid),
				str: uid.String(),
			},
			args: args{
				other: cid,
			},
			want: true,
		},
		{
			name: "引数に渡したCommentIDが自分自身のオブジェクトと異なる場合にfalseが返却される",
			fields: fields{
				id:  commentID(uid),
				str: "",
			},
			args: args{
				other: cid,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommentID{
				id:  tt.fields.id,
				str: tt.fields.str,
			}
			if got := c.Equals(tt.args.other); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentID_Get(t *testing.T) {
	uid, _ := uuid.NewRandom()
	type fields struct {
		id  commentID
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   CommentID
	}{
		{
			name: "作成したIDが取得出来る",
			fields: fields{
				id:  commentID(uid),
				str: uid.String(),
			},
			want: CommentID{
				id:  commentID(uid),
				str: uid.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommentID{
				id:  tt.fields.id,
				str: tt.fields.str,
			}
			if got := c.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentID_String(t *testing.T) {
	uid, _ := uuid.NewRandom()
	type fields struct {
		id  commentID
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "作成したIDが文字列として取得出来る",
			fields: fields{
				id:  commentID(uid),
				str: uid.String(),
			},
			want: uid.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommentID{
				id:  tt.fields.id,
				str: tt.fields.str,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCommentID(t *testing.T) {
	uid, _ := uuid.NewRandom()
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    CommentID
		wantErr bool
	}{
		{
			name: "引数にUUIDの仕様を満たした文字列を渡した場合、CommentIDに変換される",
			args: args{
				ID: uid.String(),
			},
			want: CommentID{
				id:  commentID(uid),
				str: uid.String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCommentID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCommentID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentID() got = %v, want %v", got, tt.want)
			}
		})
	}

	// uuidの性質上テーブルテストではカバー出来ないので独自実装しています。
	t.Run("引数がUUIDの仕様を満たしていない場合、引数の値は無視してCommentIDが返却される", func(t *testing.T) {
		cid, err := NewCommentID("this is not uuid.")
		if err != nil {
			t.Errorf("NewCommentID() error = %v, wantErr false", err)
		}

		_, err = uuid.Parse(cid.String())
		if err != nil {
			t.Errorf("NewCommentID parse error = %v, wantErr false", err)
		}
	})

	t.Run("引数が空文字の場合に新規IDが生成出来る", func(t *testing.T) {
		cid, err := NewCommentID("")
		if err != nil {
			t.Errorf("NewCommentID() error = %v, wantErr false", err)
		}

		_, err = uuid.Parse(cid.String())
		if err != nil {
			t.Errorf("NewCommentID parse error = %v, wantErr false", err)
		}
	})

}
