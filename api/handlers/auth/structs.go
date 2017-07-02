package auth

type CredentialsData struct {
	Data *Credentials `json:"data"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
