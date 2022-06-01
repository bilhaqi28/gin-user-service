package response

type UserLogin struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}
