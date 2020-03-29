package entities

type User struct {
	ID        string
	LifePoint int
}

func NewUser(lifePoint int) *User {
	return &User{LifePoint: lifePoint}
}
