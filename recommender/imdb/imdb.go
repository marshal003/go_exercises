package imdb

import (
	"github.com/marshal003/exercises/recommender/movie"
	"github.com/marshal003/exercises/recommender/rating"
	user "github.com/marshal003/exercises/recommender/user"
)

//IMDB data structure for IMDB
type IMDB struct {
	movieMap           map[int]movie.Movie
	ratings            []rating.UserRating
	users              map[int]user.User
	averageMovieRating map[int]AverageRating
}

// CreateIMDB Factory method to create initialize IMDB
func CreateIMDB(movieMap map[int]movie.Movie, ratings []rating.UserRating, users map[int]user.User) IMDB {
	imdb := IMDB{movieMap: movieMap, ratings: ratings, users: users}
	imdb.calculateAverageMovieRating()
	return imdb
}

// AverageRating datastructure to contain average movie rating
type AverageRating struct {
	ratingCount   int64
	averageRating float32
	movieID       int
}

func (mr *AverageRating) updateRatingAverage(rating float32) {
	mr.averageRating = (float32(mr.ratingCount)*mr.averageRating + rating) / float32(mr.ratingCount+1)
	mr.ratingCount++
}

// CalculateAverageMovieRating get average rating for each movie
func (imdb *IMDB) calculateAverageMovieRating() {
	movieRatings := make(map[int]AverageRating)
	for _, userRating := range imdb.ratings {
		movieRating, ok := movieRatings[userRating.MovieID]
		if !ok {
			movieRating = AverageRating{movieID: userRating.MovieID, ratingCount: 1, averageRating: 0}
		}
		movieRating.updateRatingAverage(userRating.Rating)
		movieRatings[userRating.MovieID] = movieRating
	}
	imdb.averageMovieRating = movieRatings
}
