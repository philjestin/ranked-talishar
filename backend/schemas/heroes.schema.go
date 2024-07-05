package schemas

import "github.com/google/uuid"

type CreateHero struct {
	HeroName string `json:"hero_name" binding:"required"`
	FormatID uuid.NullUUID `json:"format_id" binding:"required"`
}

type UpdateHero struct {
	HeroName string `json:"hero_name"`
	FormatID uuid.NullUUID `json:"format_id"`
}