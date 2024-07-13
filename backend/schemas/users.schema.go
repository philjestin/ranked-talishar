package schemas

type CreateUser struct {
    UserName     string `json:"user_name" binding:"required"`
    UserEmail    string `json:"user_email" binding:"required"`
    HashedPassword string `json:"hashed_password" binding:"required"`
}

type UpdateUser struct {
    UserName     string `json:"user_name"`
    UserEmail    string `json:"user_email"`
    HashedPassword string `json:"hashed_password"`
}