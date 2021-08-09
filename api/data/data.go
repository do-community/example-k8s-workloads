package data

type User struct {
	Id                int    `json:"id"`
	UserName          string `json:"user_name"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	AvatarURL         string `json:"avatar_url"`
	Company           string `json:"company"`
	MostRecentComment string `json:"most_recent_comment"`
	LastLogin         string `json:"last_login"`
}
