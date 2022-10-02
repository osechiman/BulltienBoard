package errorobjects

import (
	"reflect"
	"testing"
)

func TestNewDatabaseConnectionError(t *testing.T) {
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name string
		args args
		want *DatabaseConnectionError
	}{
		{
			name: "エラーオブジェクトが正常に生成される",
			args: args{
				msg: "database connection error object test",
			},
			want: &DatabaseConnectionError{
				msg:  "database connection error object test",
				code: ErrorCodeDatabaseConnectionError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDatabaseConnectionError(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatabaseConnectionError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseConnectionError_Error(t *testing.T) {
	type fields struct {
		msg  string
		code int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "msgの値とdefaultのエラーメッセージが結合されて出力されてくる",
			fields: fields{
				msg:  "database connection error object test",
				code: ErrorCodeDatabaseConnectionError,
			},
			want: "database connection error object test. error code is 6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dce := &DatabaseConnectionError{
				msg:  tt.fields.msg,
				code: tt.fields.code,
			}
			if got := dce.Error(); got != tt.want {
				t.Errorf("DatabaseConnectionError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
