package request

type CreateUser struct {
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
