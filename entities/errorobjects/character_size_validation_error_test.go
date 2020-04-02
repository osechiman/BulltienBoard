package errorobjects

import (
	"reflect"
	"testing"
)

func TestCharacterSizeValidationError_Error(t *testing.T) {
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
				msg:  "character size validation error object test",
				code: ErrorCodeCharacterSizeValidation,
			},
			want: "character size validation error object test. error code is 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ise := &CharacterSizeValidationError{
				msg:  tt.fields.msg,
				code: tt.fields.code,
			}
			if got := ise.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCharacterSizeValidationError(t *testing.T) {
	type args struct {
		msg interface{}
	}
	tests := []struct {
		name string
		args args
		want *CharacterSizeValidationError
	}{
		{
			name: "エラーオブジェクトが正常に生成される",
			args: args{
				msg: "character size validation error object test",
			},
			want: &CharacterSizeValidationError{
				msg:  "character size validation error object test",
				code: ErrorCodeCharacterSizeValidation,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCharacterSizeValidationError(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCharacterSizeValidationError() = %v, want %v", got, tt.want)
			}
		})
	}
}
