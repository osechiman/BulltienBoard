package valueobjects

import (
	"reflect"
	"testing"
	"time"
)

func TestCommentTime_Equals(t *testing.T) {
	ut := time.Now().Unix()
	ct, _ := NewCommentTime(ut)
	type fields struct {
		unixTime int64
	}
	type args struct {
		other CommentTime
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "引数に渡したCommentTimeが自分自身のオブジェクトと同一の場合にtrueが返却される",
			fields: fields{
				unixTime: ut,
			},
			args: args{
				other: ct,
			},
			want: true,
		},
		{
			name: "引数に渡したCommentTimeが自分自身のオブジェクトと異なる場合にfalseが返却される",
			fields: fields{
				unixTime: 0,
			},
			args: args{
				other: ct,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommentTime{
				unixTime: tt.fields.unixTime,
			}
			if got := c.Equals(tt.args.other); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentTime_Get(t *testing.T) {
	ut := time.Now().Unix()
	type fields struct {
		unixTime int64
	}
	tests := []struct {
		name   string
		fields fields
		want   CommentTime
	}{
		{
			name: "作成したCommentTimeが取得出来る",
			fields: fields{
				unixTime: ut,
			},
			want: CommentTime{
				unixTime: ut,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommentTime{
				unixTime: tt.fields.unixTime,
			}
			if got := c.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommentTime_ToUnixTime(t *testing.T) {
	ut := time.Now().Unix()
	type fields struct {
		unixTime int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "CommentTimeがunixTimeに変換されて返却される",
			fields: fields{
				unixTime: ut,
			},
			want: ut,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CommentTime{
				unixTime: tt.fields.unixTime,
			}
			if got := c.ToUnixTime(); got != tt.want {
				t.Errorf("ToUnixTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCommentTime(t *testing.T) {
	ut := time.Now().Unix()
	type args struct {
		unixTime int64
	}
	tests := []struct {
		name    string
		args    args
		want    CommentTime
		wantErr bool
	}{
		{
			name: "引数にマイナスの値を渡した場合、現在時刻が返却される",
			args: args{
				unixTime: -1,
			},
			want: CommentTime{
				unixTime: ut,
			},
		},
		{
			name: "引数にプラスの値を渡した場合、その値がunixTimeとして採用され返却される",
			args: args{
				unixTime: ut,
			},
			want: CommentTime{
				unixTime: ut,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCommentTime(tt.args.unixTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCommentTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommentTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
