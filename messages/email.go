package messages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func getSlackUserInfoUrl(userId string) string {
	apiToken := os.Getenv("SLACK_API_TOKEN")
	return fmt.Sprintf("https://slack.com/api/users.info?token=%s&user=%s", apiToken, userId)
}

type profile struct {
	Email string `json:"email"`
}

type user struct {
	Profile profile `json:"profile"`
}

type body struct {
	User user `json:"user"`
}

// per https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang
func getJson(url string, target interface{}) error {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := netClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func getUserEmail(userId string) (string, error) {
	b := &body{}
	err := getJson(getSlackUserInfoUrl(userId), b)

	if err != nil {
		return "", err
	}

	return b.User.Profile.Email, nil
}
