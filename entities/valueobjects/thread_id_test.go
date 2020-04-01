package valueobjects

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestThreadID_Equals(t *testing.T) {
	uid, _ := uuid.NewRandom()
	tid, _ := NewThreadID(uid.String())
	type fields struct {
		id  threadID
		str string
	}
	type args struct {
		other ThreadID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "引数に渡したThreadIDが自分自身のオブジェクトと同一の場合にtrueが返却される",
			fields: fields{
				id:  threadID(uid),
				str: uid.String(),
			},
			args: args{
				other: tid,
			},
			want: true,
		},
		{
			name: "引数に渡したThreadIDが自分自身のオブジェクトと異なる場合にfalseが返却される",
			fields: fields{
				id:  threadID(uid),
				str: "",
			},
			args: args{
				other: tid,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ThreadID{
				id:  tt.fields.id,
				str: tt.fields.str,
			}
			if got := c.Equals(tt.args.other); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreadID_Get(t *testing.T) {
	uid, _ := uuid.NewRandom()
	type fields struct {
		id  threadID
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   ThreadID
	}{
		{
			name: "作成したIDが取得出来る",
			fields: fields{
				id:  threadID(uid),
				str: uid.String(),
			},
			want: ThreadID{
				id:  threadID(uid),
				str: uid.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ThreadID{
				id:  tt.fields.id,
				str: tt.fields.str,
			}
			if got := c.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThreadID_String(t *testing.T) {
	uid, _ := uuid.NewRandom()
	type fields struct {
		id  threadID
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
				id:  threadID(uid),
				str: uid.String(),
			},
			want: uid.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ThreadID{
				id:  tt.fields.id,
				str: tt.fields.str,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewThreadID(t *testing.T) {
	uid, _ := uuid.NewRandom()
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    ThreadID
		wantErr bool
	}{
		{
			name: "引数にUUIDの仕様を満たした文字列を渡した場合、ThreadIDに変換される",
			args: args{
				ID: uid.String(),
			},
			want: ThreadID{
				id:  threadID(uid),
				str: uid.String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewThreadID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewThreadID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewThreadID() got = %v, want %v", got, tt.want)
			}
		})
	}

	// uuidの性質上テーブルテストではカバー出来ないので独自実装しています。
	t.Run("引数がUUIDの仕様を満たしていない場合、引数の値は無視してThreadIDが返却される", func(t *testing.T) {
		tid, err := NewThreadID("this is not uuid.")
		if err != nil {
			t.Errorf("NewThreadID() error = %v, wantErr false", err)
		}

		_, err = uuid.Parse(tid.String())
		if err != nil {
			t.Errorf("NewThreadID parse error = %v, wantErr false", err)
		}
	})

	t.Run("引数が空文字の場合に新規IDが生成出来る", func(t *testing.T) {
		tid, err := NewThreadID("")
		if err != nil {
			t.Errorf("NewThreadID() error = %v, wantErr false", err)
		}

		_, err = uuid.Parse(tid.String())
		if err != nil {
			t.Errorf("NewThreadID parse error = %v, wantErr false", err)
		}
	})

}
