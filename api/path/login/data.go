package login

type Datacookie struct {
	user   string
	passed string
	ver    bool
}
type GetDatacookie struct {
	Key []string
}
type ln struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
