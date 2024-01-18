package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetAge() int8
	GetName() string
	GetID() string
	//GetJSONValue() (string, error)

	SetID(string)

	EncryptPassword()
}

// NewUserDomain constructor
func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		name:     name,
		password: password,
		age:      age,
	}
}

func NewUpdateUserDomain(name string, age int8) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}
