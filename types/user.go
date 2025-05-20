package types

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Auth struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
