package presenters

import (
	"reflect"
	"testing"
	"vspro/entities"
)

func TestAnswerQuestion(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestListQuestion(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestNewQuestionPresenter(t *testing.T) {
	tests := []struct {
		name string
		want *QuestionPresenter
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuestionPresenter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuestionPresenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestionPresenter_ConvertToHttpErrorResponse(t *testing.T) {
	type fields struct {
		HTTPQuestionResponse HTTPQuestionResponse
	}
	type args struct {
		httpStatusCode int
		err            error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *HTTPQuestionResponse
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := &QuestionPresenter{
				HTTPQuestionResponse: tt.fields.HTTPQuestionResponse,
			}
			if got := qp.ConvertToHttpErrorResponse(tt.args.httpStatusCode, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToHttpErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestionPresenter_ConvertToHttpQuestionListResponse(t *testing.T) {
	type fields struct {
		HTTPQuestionResponse HTTPQuestionResponse
	}
	type args struct {
		ql []*entities.Question
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *HTTPQuestionResponse
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := &QuestionPresenter{
				HTTPQuestionResponse: tt.fields.HTTPQuestionResponse,
			}
			if got := qp.ConvertToHttpQuestionListResponse(tt.args.ql); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToHttpQuestionListResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestionPresenter_ConvertToHttpQuestionResponse(t *testing.T) {
	type fields struct {
		HTTPQuestionResponse HTTPQuestionResponse
	}
	type args struct {
		q *entities.Question
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *HTTPQuestionResponse
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := &QuestionPresenter{
				HTTPQuestionResponse: tt.fields.HTTPQuestionResponse,
			}
			if got := qp.ConvertToHttpQuestionResponse(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToHttpQuestionResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertEntitiesQuestionToQuestion(t *testing.T) {
	type args struct {
		q *entities.Question
	}
	tests := []struct {
		name string
		args args
		want *Question
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertEntitiesQuestionToQuestion(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertEntitiesQuestionToQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newHTTPErrorResponse(t *testing.T) {
	type args struct {
		status  int
		message string
		err     error
	}
	tests := []struct {
		name string
		args args
		want *HTTPQuestionResponse
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHTTPErrorResponse(tt.args.status, tt.args.message, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHTTPErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newHTTPQuestionResponse(t *testing.T) {
	type args struct {
		status  int
		message string
		data    Questions
	}
	tests := []struct {
		name string
		args args
		want *HTTPQuestionResponse
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHTTPQuestionResponse(tt.args.status, tt.args.message, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHTTPQuestionResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
