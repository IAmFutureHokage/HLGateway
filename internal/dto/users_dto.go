package dto

type User struct {
	ID         string `json:"id"`
	Role       string `json:"role"`
	PostCode   string `json:"post_code"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type UserRequest struct {
	User User `json:"user"`
}

type UserResponse struct {
	User User `json:"user"`
}

type UserIDRequest struct {
	ID string `json:"id"`
}

type UsersResponse struct {
	Users []User `json:"users"`
}
