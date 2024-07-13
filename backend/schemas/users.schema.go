package schemas

type CreateUser struct {
    UserName     string `json:"user_name" binding:"required,alphanum"`
    UserEmail    string `json:"user_email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type UpdateUser struct {
    UserName     string `json:"user_name"`
    UserEmail    string `json:"user_email"`
    Password string `json:"password"`
}