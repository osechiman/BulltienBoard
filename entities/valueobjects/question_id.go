package valueobjects

// ValueObjectは外部ライブラリに依存して良いというポリシーで運用
import (
	"github.com/google/uuid"
)

type questionID uuid.UUID

type QuestionID struct {
	id  questionID
	str string
}

func NewQuestionID(ID string) (QuestionID, error) {
	uid, err := uuid.Parse(ID)
	if err != nil {
		uid, err = uuid.NewRandom()
		if err != nil {
			return QuestionID{}, err
		}
	}

	qid := QuestionID{id: questionID(uid), str: uid.String()}
	return qid, nil
}

func (q QuestionID) Get() QuestionID {
	return q
}

func (q QuestionID) String() string {
	return q.str
}

func (q QuestionID) Equals(other QuestionID) bool {
	if (q.Get() == other.Get()) && (q.String() == other.String()) {
		return true
	}
	return false
}
