package chain_of_responsibility

type userInfo struct {
	Id      int
	Name    string
	IsAdmin bool
}

func NewUserInfo(id int, name string, isadmin bool) *userInfo {
	return &userInfo{Id: id, Name: name, IsAdmin: isadmin}
}
