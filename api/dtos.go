package api

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type EventDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	DateTime    string `json:"date_time"`
	UserId      int    `json:"user_id"`
}
