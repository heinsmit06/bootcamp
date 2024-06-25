package firststruct

type User struct {
	Name     string
	password string
}

func NewUser(name, password string) User {
	new_user := User{
		Name:     name,
		password: password,
	}
	return new_user
}

func (u User) GetPassword() string {
	return u.password
}
