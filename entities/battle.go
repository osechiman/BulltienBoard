package entities

type Battle struct {
	ID          string
	Participant map[string]User
}

func NewBattle(participant map[string]User) *Battle {
	return &Battle{Participant: participant}
}
