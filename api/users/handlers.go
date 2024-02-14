package users

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

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
	if err == nil && x < 30 {
		limit = x
	}

	if len(users) == limit {
		fmt.Println("From cache")
		res = map[string]interface{}{
			"success": true,
			"length":  len(users),
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

		var wg sync.WaitGroup
		var m sync.Mutex

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		for i := 0; i < limit; i++ {
			wg.Add(1)
			username := usernames[i].Login
			if i == 5 {
				usersUrl = "https://api.github.com/usersjjjjjjjj"
			}

			var userUrl = fmt.Sprintf("%v/%v", usersUrl, username)

			go fetchUsersInfo(userUrl, &users, &wg, &m, &ctx)
			if err != nil {
				utils.HandleServerError(err, c)
				break
			}

		}
		wg.Wait()
		res = map[string]interface{}{
			"success": true,
			"length":  len(users),
			"users":   users,
		}
	}

	utils.SendSuccessJSONResponse(res, c)
}
