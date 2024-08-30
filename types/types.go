package types

type TokenClaims struct {
	Id    uint
	Name  string
	Email string
}
type CreateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginUser struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
