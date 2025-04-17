package dto

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=4"`
}

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type StartChatRequest struct {
	TargetID string `json:"target_id" binding:"required"`
}

type ChatResponse struct {
	ID      string   `json:"id"`
	Members []string `json:"members"`
	IsGroup bool     `json:"is_group"`
}
