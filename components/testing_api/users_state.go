package testingapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type UserStore []User

type UserData struct {
	Password   string
	BaseCredit float64
}

type User struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Credit   float64 `json:"credit"`
}

var userMap map[string]UserData

func initUsers() bool {
	userMap = make(map[string]UserData)

	fileBytes, err := os.ReadFile("assets/json/users.json")

	if err != nil {
		fmt.Println("Error when reading users.json file")
		return false
	}

	var users UserStore
	err = json.Unmarshal(fileBytes, &users)
	if err != nil {
		fmt.Println("Error when parsing users.json file")
		return false
	}

	for _, user := range users {
		newUserData := UserData{Password: user.Password, BaseCredit: user.Credit}
		userMap[user.Username] = newUserData
	}

	return true
}

func getUserPassword(username string) (string, error) {
	userData, ok := userMap[username]
	if !ok {
		return "", errors.New("User not found")
	}
	return userData.Password, nil
}

func getUserBaseCredit(username string) (float64, error) {
	userData, ok := userMap[username]
	if !ok {
		return 0.0, errors.New("User not found")
	}
	return userData.BaseCredit, nil
}

func getAllUsers() []string {
	usernameArray := []string{}
	for username := range userMap {
		usernameArray = append(usernameArray, username)
	}
	return usernameArray
}
