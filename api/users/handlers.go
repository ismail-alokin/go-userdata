package users

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ismail-alokin/go-userdata/utils"
)

var users []UserData

func GetUserInformationList(c *gin.Context) {
	fmt.Println("len users", len(users))

	var res map[string]interface{}
	limit := 10
	limitStr := c.Query("limit")

	x, err := strconv.Atoi(limitStr)
	if err == nil && x < 100 {
		limit = x
	}

	if len(users) == limit {
		fmt.Println("From cache")
		res = map[string]interface{}{
			"success": true,
			"users":   users,
		}
	} else {
		users = []UserData{}
		fmt.Println("Hitting Github")
		var usersUrl = "https://api.github.com/users"
		usernames, err := fetchUsernameList(usersUrl)
		if err != nil {
			log.Println("Error occured while fetching usernames: ", err)
			utils.HandleServerError(err, c)
			return
		}

		for i := 0; i < limit; i++ {
			username := usernames[i].Login

			var userUrl = fmt.Sprintf("%v/%v", usersUrl, username)
			var user UserData

			err := fetchUserInfo(userUrl, &user)
			if err != nil {
				utils.HandleServerError(err, c)
				break
			}

			users = append(users, user)
		}

		res = map[string]interface{}{
			"success": true,
			"users":   users,
		}
	}

	utils.SendSuccessJSONResponse(res, c)
}
