package auth

type UserLoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type PasswordUpdatePayload struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type UsersListItem struct {
	Id    uint   `json:"user_id"`
	Email string `json:"email"`
}

type UsersListPayload struct {
	Users []UsersListItem
}

type UserUpdatePayload struct {
	Email         string `json:"email" validate:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
	MaritalStatus string `json:"marital_status"`
	// ResidentialDetails *ResidentialDetails `validate:"json"`
	// OfficeDetails *OfficeDetails `json:"office_details"`
}

type UserDetails struct {
	Id                 uint               `json:"id"`
	Email              string             `json:"email"`
	FirstName          string             `json:"first_name"`
	LastName           string             `json:"last_name"`
	Gender             string             `json:"gender"`
	MaritalStatus      string             `json:"marital_status"`
	ResidentialDetails ResidentialDetails `json:"residential_details"`
	OfficeDetails      OfficeDetails      `json:"office_details"`
}

type ResidentialDetails struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Contact string `json:"contact"`
}

type OfficeDetails struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	EmployeeCode string `json:"employee_code"`
	Address      string
	City         string
	State        string
	Country      string
	Contact      string
}
