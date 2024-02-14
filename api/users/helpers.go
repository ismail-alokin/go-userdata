package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func fetchUsernameList(usersUrl string) ([]Username, error) {
	resp, err := sendGetRequestToGitHub(usersUrl)

	if err != nil || resp.Status[:1] == "4" || resp.Status[:1] == "5" {
		log.Println("Error: ", err, "\n\nResponse: ", resp)
		var customErr error
		if err != nil {
			customErr = err
		} else {
			customErr = errors.New("server error")
		}
		return []Username{}, customErr
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []Username{}, err
	}

	var usernames []Username
	json.Unmarshal([]byte(body), &usernames)

	return usernames, nil
}

func fetchUserInfo(userUrl string, user *UserData) error {
	resp, err := sendGetRequestToGitHub(userUrl)

	if err != nil || resp.Status[:1] == "4" || resp.Status[:1] == "5" {
		log.Println("Error: ", err, "\n\nResponse: ", resp)
		var customErr error
		if err != nil {
			customErr = err
		} else {
			customErr = errors.New("server error")
		}
		return customErr
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	json.Unmarshal([]byte(body), user)
	return nil

}

func sendGetRequestToGitHub(url string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	authHeader := fmt.Sprintf("Bearer %v", githubToken)

	req.Header = http.Header{
		"Content-Type":  {"application/vnd.github+json"},
		"Authorization": {authHeader},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
