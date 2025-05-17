package entity

type User struct {
	UserID int32
	Name   string
	Email  string
}

func (u User) Validate() error {
	// .....
	return nil
}
