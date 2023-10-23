package User

type userInf interface {
	UserNew(name string, password string, email string)
	GetNameMthd() string
	GetEmail() string
	GetUserInfo() string
}

type UserStruct struct {
	name     string
	password string
	email    string
}

func GetName() string {
	return "hello"
}

func UserNew(name string, password string, email string) UserStruct {
	var instUserStruct UserStruct
	instUserStruct.name = name
	instUserStruct.email = email
	instUserStruct.password = password
	return instUserStruct
}

func (instUserStruct UserStruct) GetNameMthd() string {
	return instUserStruct.name
}
