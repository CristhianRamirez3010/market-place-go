package usersEndpoints

import "fmt"

type UsersEndpoints struct {
	Create string
}

func BuildUserEndpoints() UsersEndpoints {
	user := "/user/"
	return UsersEndpoints{
		Create: fmt.Sprintf("%s%s", user, "create"),
	}
}
