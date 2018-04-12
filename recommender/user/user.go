package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	utils "github.com/marshal003/exercises/recommender/utils"
)

// User data structure for User
type User struct {
	userID     int
	age        int
	gender     string
	occupation string
	zipCode    int
}

//ReadUsers implementing UserReader interface
func ReadUsers(dataFile utils.DataFile) map[int]User {
	channel := make(chan []string, 0)
	go dataFile.AsyncDataFileReader(channel)

	userToIDMap := make(map[int]User, 0)
	for tokens := range channel {
		user, err := parseUser(tokens)
		if err != nil {
			continue
		}
		userToIDMap[user.userID] = user
	}
	return userToIDMap
}

func parseUser(tokens []string) (User, error) {
	var user User
	if len(tokens) != 5 {
		return user, errors.New("Tokens count should be 5")
	}
	userID, err := strconv.Atoi(strings.Trim(tokens[0], " "))
	if err != nil {
		msg := fmt.Sprintf("Unable to parse userID: %s, as int" + tokens[0])
		return user, errors.New(msg)
	}
	age, err := strconv.Atoi(strings.Trim(tokens[1], " "))
	gender := strings.Trim(tokens[2], " ")
	occupation := strings.Trim(tokens[3], " ")
	zipCode, err := strconv.Atoi(strings.Trim(tokens[4], " "))
	return User{userID: userID, age: age, gender: gender, occupation: occupation, zipCode: zipCode}, nil
}
