package errorobjects

import (
	"net/http"
	"reflect"
	"testing"
)

func TestMissingRequiredFieldsError_Error(t *testing.T) {
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
				msg:            "missing required field error object test",
				code:           ErrorCodeMissingRequiredFiled,
				HTTPStatusCode: http.StatusBadRequest,
			},
			want: "missing required field error object test. error code is 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mrfe := &MissingRequiredFieldsError{
				msg:            tt.fields.msg,
				code:           tt.fields.code,
				HTTPStatusCode: tt.fields.HTTPStatusCode,
			}
			if got := mrfe.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMissingRequiredFieldsError(t *testing.T) {
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name string
		args args
		want *MissingRequiredFieldsError
	}{
		{
			name: "エラーオブジェクトが正常に生成される",
			args: args{
				msg: "missing required field error object test",
			},
			want: &MissingRequiredFieldsError{
				msg:            "missing required field error object test",
				code:           ErrorCodeMissingRequiredFiled,
				HTTPStatusCode: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMissingRequiredFieldsError(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMissingRequiredFieldsError() = %v, want %v", got, tt.want)
			}
		})
	}
}
