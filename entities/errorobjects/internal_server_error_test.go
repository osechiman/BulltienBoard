package errorobjects

import (
	"reflect"
	"testing"
)

func TestInternalServerError_Error(t *testing.T) {
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
				msg:  "internal server error object test",
				code: ErrorCodeInternalServerError,
			},
			want: "internal server error object test. error code is 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ise := &InternalServerError{
				msg:  tt.fields.msg,
				code: tt.fields.code,
			}
			if got := ise.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInternalServerError(t *testing.T) {
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name string
		args args
		want *InternalServerError
	}{
		{
			name: "エラーオブジェクトが正常に生成される",
			args: args{
				msg: "internal server error object test",
			},
			want: &InternalServerError{
				msg:  "internal server error object test",
				code: ErrorCodeInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInternalServerError(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInternalServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}
