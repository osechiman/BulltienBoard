package errorobjects

import (
	"net/http"
	"reflect"
	"testing"
)

func TestInternalServerError_Error(t *testing.T) {
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
			name: "エラーメッセージ出力テスト",
			fields: fields{
				msg:            "internal server error object test",
				code:           ErrorCodeInternalServerError,
				HTTPStatusCode: http.StatusInternalServerError,
			},
			want: "internal server error object test. error code is 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ise := &InternalServerError{
				msg:            tt.fields.msg,
				code:           tt.fields.code,
				HTTPStatusCode: tt.fields.HTTPStatusCode,
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
			name: "オブジェクト生成テスト",
			args: args{
				msg: "internal server error object test",
			},
			want: &InternalServerError{
				msg:            "internal server error object test",
				code:           ErrorCodeInternalServerError,
				HTTPStatusCode: http.StatusInternalServerError,
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
