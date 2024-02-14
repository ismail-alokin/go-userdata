package users

import "time"

type UserData struct {
	Name             string    `json:"name"`
	Company          string    `json:"company"`
	Location         string    `json:"location"`
	Email            string    `json:"email"`
	Bio              string    `json:"bio"`
	Twitter_username string    `json:"twitter_username"`
	Created_at       time.Time `json:"created_at"`
	Updated_at       time.Time `json:"updates_at"`
}

type Username struct {
	Login string `json:"login"`
}
