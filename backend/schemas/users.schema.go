package schemas

type CreateUser struct {
    UserName     string `json:"user_name" binding:"required"`
    UserEmail    string `json:"user_email" binding:"required"`
}

type UpdateUser struct {
    UserName     string `json:"user_name"`
    UserEmail    string `json:"user_email"`
}