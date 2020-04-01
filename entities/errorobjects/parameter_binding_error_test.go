package errorobjects

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewParameterBindingError(t *testing.T) {
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name string
		args args
		want *ParameterBindingError
	}{
		{
			name: "エラーオブジェクトが正常に生成される",
			args: args{
				msg: "parameter binding error object test",
			},
			want: &ParameterBindingError{
				msg:            "parameter binding error object test",
				code:           ErrorCodeParameterBinding,
				HTTPStatusCode: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParameterBindingError(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParameterBindingError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameterBindingError_Error(t *testing.T) {
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
				msg:            "parameter binding error object test",
				code:           ErrorCodeParameterBinding,
				HTTPStatusCode: http.StatusBadRequest,
			},
			want: "parameter binding error object test. error code is 3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pbe := &ParameterBindingError{
				msg:            tt.fields.msg,
				code:           tt.fields.code,
				HTTPStatusCode: tt.fields.HTTPStatusCode,
			}
			if got := pbe.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
