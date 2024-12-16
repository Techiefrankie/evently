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
	UserId      int    `json:"user_id" binding:"required"`
}

type UserDto struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
