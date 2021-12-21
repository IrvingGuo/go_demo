package entity

type Signin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AutoSignin struct {
	Token string `header:"Authorization"`
}

type StatusPayload struct {
	Base
	Status uint `json:"status"`
}
