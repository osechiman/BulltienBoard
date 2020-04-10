package errorobjects

import (
	"reflect"
	"testing"
)

func TestResourceLimitedError_Error(t *testing.T) {
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
				msg:  "resource limited error object test",
				code: ErrorCodeResourceLimitedError,
			},
			want: "resource limited error object test. error code is 5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ise := &ResourceLimitedError{
				msg:  tt.fields.msg,
				code: tt.fields.code,
			}
			if got := ise.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResourceLimitedError(t *testing.T) {
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name string
		args args
		want *ResourceLimitedError
	}{
		{
			name: "エラーオブジェクトが正常に生成される",
			args: args{
				msg: "resource limited error object test",
			},
			want: &ResourceLimitedError{
				msg:  "resource limited error object test",
				code: ErrorCodeResourceLimitedError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResourceLimitedError(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourceLimitedError() = %v, want %v", got, tt.want)
			}
		})
	}
}
