package schemas

type CreateGame struct {
	GameName  string `json:"game_name" binding:"required"`
}

type UpdateGame struct {
	GameName string `json:"game_name"`
}