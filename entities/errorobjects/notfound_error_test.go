package errorobjects

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewNotFoundError(t *testing.T) {
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name string
		args args
		want *NotFoundError
	}{
		{
			name: "エラーオブジェクトが正常に生成される",
			args: args{
				msg: "not found error object test",
			},
			want: &NotFoundError{
				msg:            "not found error object test",
				code:           ErrorCodeNotFound,
				HTTPStatusCode: http.StatusNotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotFoundError(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotFoundError_Error(t *testing.T) {
	type fields struct {
		msg            string
		code           int
		HTTPStatusCode int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "msgの値とdefaultのエラーメッセージが結合されて出力されてくる",
			fields: fields{
				msg:            "not found error object test",
				code:           ErrorCodeNotFound,
				HTTPStatusCode: http.StatusNotFound,
			},
			want: "not found error object test not found. error code is 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nfe := &NotFoundError{
				msg:            tt.fields.msg,
				code:           tt.fields.code,
				HTTPStatusCode: tt.fields.HTTPStatusCode,
			}
			if got := nfe.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
