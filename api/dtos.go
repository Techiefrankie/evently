package api

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type EventDto struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description" validate:"required,min=3"`
	Location    string `json:"location" validate:"required,min=3"`
	DateTime    string `json:"date_time" validate:"datetime"`
	UserId      int    `json:"user_id"`
}

type UserDto struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name" validate:"required,min=2,max=100,msg=First name is required and must be between 2 and 100 characters"`
	LastName  string `json:"last_name" validate:"required,min=2,max=100,msg=Last name is required and must be between 2 and 100 characters"`
	Email     string `json:"email" validate:"required,email,msg=A valid email is required"`
	Password  string `json:"password" validate:"required,Password,msg=Password must be 8-20 characters long and contain at least one uppercase letter; one lowercase letter; one digit and one special character"`
}

type LoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Email       string   `json:"email"`
	UserId      int      `json:"user_id"`
	AccessToken string   `json:"access_token"`
	ExpiresIn   int      `json:"expires_in"`
	Roles       []string `json:"roles"`
}
