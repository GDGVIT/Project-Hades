package views

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password`
}

type Role struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type AddMembers struct {
	Email  string `json:"email"`
	Org    string `json:"org"`
	Accept bool   `json:"accept"`
}
