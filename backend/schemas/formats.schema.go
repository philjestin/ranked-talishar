package schemas

type CreateFormat struct {
	FormatName string `json:"format_name" binding:"required"`
	FormatDescription string `json:"format_description" binding:"required"`
}

type UpdateFormat struct {
	FormatName string `json:"format_name"`
	FormatDescription string `json:"format_description"`
}