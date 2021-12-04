package models

type User struct {
	Id        uint
	FirstName string
	LastName  string
	Password  string
	Age       uint
	Gender    string
}

func NewUser() *User {
	return &User{}
}

func (u *User) SetId(id uint) *User {
	u.Id = id
	return u
}

func (u *User) SetFirstName(firstName string) *User {
	u.FirstName = firstName
	return u
}

func (u *User) SetLastName(latName string) *User {
	u.LastName = latName
	return u
}

func (u *User) SetPassword(password string) *User {
	u.Password = password
	return u
}

func (u *User) SetAge(age uint) *User {
	u.Age = age
	return u
}

func (u *User) SetGender(gender string) *User {
	u.Gender = gender
	return u
}

func (u *User) Build() *User {
	return &User{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		Age:       u.Age,
		Gender:    u.Gender,
	}
}
