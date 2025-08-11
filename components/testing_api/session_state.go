package testingapi

import (
	"crypto/sha512"
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var sessionStore map[string]string

func initSessions() {
	sessionStore = make(map[string]string)
}

func generateSession(username string, password string) (string, error) {
	passwordHash, err := getUserPassword(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	if err != nil {
		return "", errors.New("incorrect password")
	}

	sessionIdBytes := sha512.Sum512([]byte(username + password + strconv.FormatInt(time.Now().UnixMicro(), 27)))
	sessionId := string(sessionIdBytes[:])
	sessionStore[sessionId] = username

	return sessionId, nil
}

func getSessionUsername(sessionId string) (string, error) {
	username, ok := sessionStore[sessionId]
	if !ok {
		return "", errors.New("User not found")
	}
	return username, nil
}
