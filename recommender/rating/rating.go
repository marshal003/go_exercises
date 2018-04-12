package rating

import (
	"errors"
	"strconv"
	"strings"
	"time"

	utils "github.com/marshal003/exercises/recommender/utils"
)

// UserRating data structure for MovieRating given by the user
type UserRating struct {
	UserID  int
	MovieID int
	Rating  float32
	RatedOn time.Time
}

func parseRating(tokens []string) (UserRating, error) {
	var userRating UserRating
	if len(tokens) != 4 {
		return userRating, errors.New("Token count should be 4")
	}
	userID, err := strconv.Atoi(strings.Trim(tokens[0], " "))
	if err != nil {
		return userRating, errors.New("Unable to parse userID")
	}
	movieID, err := strconv.Atoi(strings.Trim(tokens[1], " "))
	if err != nil {
		return userRating, errors.New("Unable to parse MovieID")
	}
	rating, err := strconv.ParseFloat(strings.Trim(tokens[2], " "), 32)
	if err != nil {
		return userRating, errors.New("Unable to parse rating")
	}
	timestamp, err := strconv.ParseInt(strings.Trim(tokens[3], " "), 10, 64)
	if err != nil {
		return userRating, errors.New("Unable to parse timestamp")
	}
	ratedOn := time.Unix(timestamp, 0)

	userRating = UserRating{
		UserID:  userID,
		MovieID: movieID,
		Rating:  float32(rating),
		RatedOn: ratedOn,
	}
	return userRating, nil
}

//ReadUserRatings readUserRatings
func ReadUserRatings(dataFile utils.DataFile) []UserRating {
	channel := make(chan []string, 0)
	go dataFile.AsyncDataFileReader(channel)

	ratings := make([]UserRating, 0)
	for tokens := range channel {
		rating, err := parseRating(tokens)
		if err != nil {
			continue
		}
		ratings = append(ratings, rating)
	}
	return ratings
}
