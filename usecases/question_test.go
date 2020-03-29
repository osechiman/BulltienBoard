package usecases

import (
	"reflect"
	"testing"
	"vspro/adapters/gateways"
	"vspro/entities"
	"vspro/entities/valueobjects"
)

func TestNewQuestionUsecase(t *testing.T) {
	type args struct {
		Repository QuestionRepositorer
	}
	qr := gateways.NewInMemoryRepository()
	qu := NewQuestionUsecase(qr)
	tests := []struct {
		name string
		args args
		want *QuestionUsecase
	}{
		{
			name: "初期化の確認",
			args: args{
				Repository: qr,
			},
			want: qu,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuestionUsecase(tt.args.Repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuestionUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestionUsecase_GetByID(t *testing.T) {
	type fields struct {
		Repository QuestionRepositorer
	}
	type args struct {
		ID entities.QuestionID
	}
	qr := gateways.NewInMemoryRepository()
	qid, _ := valueobjects.NewQuestionID("")
	q := entities.NewQuestion(qid, 1, "answer", "text")
	qu := &QuestionUsecase{
		Repository: qr,
	}
	qu.AddQuestion(q)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Question
		wantErr bool
	}{
		{
			name: "TrueCase",
			fields: fields{
				Repository: qr,
			},
			args: args{
				ID: qid,
			},
			want:    &q,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := &QuestionUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := qu.GetQuestionByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQuestionByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuestionByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestionUsecase_List(t *testing.T) {
	type fields struct {
		Repository QuestionRepositorer
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entities.Question
		wantErr bool
	}{
		// TODO: AddQuestion test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := &QuestionUsecase{
				Repository: tt.fields.Repository,
			}
			got, err := qu.ListQuestion()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListQuestion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestionUsecase_AddQuestion(t *testing.T) {
	type fields struct {
		Repository QuestionRepositorer
	}
	type args struct {
		q entities.Question
	}
	qr := gateways.NewInMemoryRepository()
	qid, _ := valueobjects.NewQuestionID("")
	q := entities.NewQuestion(qid, 1, "answer", "text")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TrueCase",
			fields: fields{
				Repository: qr,
			},
			args: args{
				q: q,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := &QuestionUsecase{
				Repository: tt.fields.Repository,
			}
			if err := qu.AddQuestion(tt.args.q); (err != nil) != tt.wantErr {
				t.Errorf("AddQuestion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
