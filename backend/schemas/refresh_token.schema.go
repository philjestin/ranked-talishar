package schemas

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
	UserID       string `json:"userId"`
}
