package valueobjects

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestBulletinBoardID_Equals(t *testing.T) {
	uid, _ := uuid.NewRandom()
	bbid, _ := NewBulletinBoardID(uid.String())
	type fields struct {
		id  bulletinBoardID
		str string
	}
	type args struct {
		other BulletinBoardID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "引数に渡したBulletinBoardIDが自分自身のオブジェクトと同一の場合にtrueが返却される",
			fields: fields{
				id:  bulletinBoardID(uid),
				str: uid.String(),
			},
			args: args{
				other: bbid,
			},
			want: true,
		},
		{
			name: "引数に渡したBulletinBoardIDが自分自身のオブジェクトと異なる場合にfalseが返却される",
			fields: fields{
				id:  bulletinBoardID(uid),
				str: "",
			},
			args: args{
				other: bbid,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := BulletinBoardID{
				id:  tt.fields.id,
				str: tt.fields.str,
			}
			if got := bb.Equals(tt.args.other); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBulletinBoardID_Get(t *testing.T) {
	uid, _ := uuid.NewRandom()
	type fields struct {
		id  bulletinBoardID
		str string
	}
	tests := []struct {
		name   string
		fields fields
		want   BulletinBoardID
	}{
		{
			name: "作成したIDが取得出来る",
			fields: fields{
				id:  bulletinBoardID(uid),
				str: uid.String(),
			},
			want: BulletinBoardID{
				id:  bulletinBoardID(uid),
				str: uid.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := BulletinBoardID{
				id:  tt.fields.id,
				str: tt.fields.str,
			}
			if got := bb.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBulletinBoardID_String(t *testing.T) {
	uid, _ := uuid.NewRandom()
	type fields struct {
		id  bulletinBoardID
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
				id:  bulletinBoardID(uid),
				str: uid.String(),
			},
			want: uid.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bb := BulletinBoardID{
				id:  tt.fields.id,
				str: tt.fields.str,
			}
			if got := bb.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBulletinBoardID(t *testing.T) {
	uid, _ := uuid.NewRandom()
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		want    BulletinBoardID
		wantErr bool
	}{
		{
			name: "引数にUUIDの仕様を満たした文字列を渡した場合、BulletinBoardIDに変換される",
			args: args{
				ID: uid.String(),
			},
			want: BulletinBoardID{
				id:  bulletinBoardID(uid),
				str: uid.String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBulletinBoardID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBulletinBoardID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBulletinBoardID() got = %v, want %v", got, tt.want)
			}
		})
	}

	// uuidの性質上テーブルテストではカバー出来ないので独自実装しています。
	t.Run("引数がUUIDの仕様を満たしていない場合、引数の値は無視してBulletinBoardIDが返却される", func(t *testing.T) {
		bbid, err := NewBulletinBoardID("this is not uuid.")
		if err != nil {
			t.Errorf("NewBulletinBoardID() error = %v, wantErr false", err)
		}

		_, err = uuid.Parse(bbid.String())
		if err != nil {
			t.Errorf("NewBulletinBoardID parse error = %v, wantErr false", err)
		}
	})

	t.Run("引数が空文字の場合に新規IDが生成出来る", func(t *testing.T) {
		bbid, err := NewBulletinBoardID("")
		if err != nil {
			t.Errorf("NewBulletinBoardID() error = %v, wantErr false", err)
		}

		_, err = uuid.Parse(bbid.String())
		if err != nil {
			t.Errorf("NewBulletinBoardID parse error = %v, wantErr false", err)
		}
	})
}
