package player

type CrateRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Score     string `json:"score" form:"score" binding:"required"`
}
