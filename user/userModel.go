package user

// User Base Model
// https://github.com/OpenSlides/OpenSlides/blob/master/client/src/app/shared/models/users/user.ts
type User struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Title           string `json:"title"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Gender          string `json:"gender"`
	StructureLevel  string `json:"structure_level"`
	Number          string `json:"number"`
	AboutMe         string `json:"about_me"`
	GroupsID        []int  `json:"groups_id"`
	IsPresent       bool   `json:"is_present"`
	IsCommittee     bool   `json:"is_committee"`
	Email           string `json:"email"`
	LastEmailSend   string `json:"last_email_send"`
	Comment         string `json:"comment"`
	IsActive        bool   `json:"is_active"`
	DefaultPassword string `json:"comdefault_passwordments"`
}
