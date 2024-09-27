package req

type BasicDemo struct {
	Username string `json:"username" validate:"required"`
	Gender   string `json:"gender"`
}
