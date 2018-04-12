package imdb

import (
	"errors"
	"fmt"

	"github.com/marshal003/exercises/recommender/movie"
	user "github.com/marshal003/exercises/recommender/user"
)

//Recommender Struct that implements algorithm on IMDB to provide recommendations
type Recommender struct {
	Imdb IMDB
}

//IRecommender recommender for movies in IMDB database
type IRecommender interface {
	TopMovieByGenere(genere movie.Genere) MovieResult
	TopMovieByYear(year int) MovieResult
	MostWatchedMovie() movie.Movie
	MostWatchedGenere() movie.Genere
	HighestRatedGenere() movie.Genere
	MostActiveUser() user.User
}

//MovieResult Struct which contains result of queried movies
type MovieResult struct {
	Movie       movie.Movie
	MovieRating AverageRating
}

func (m MovieResult) String() string {
	return fmt.Sprintf("Movie: %s, ReleaseYear: %d, AverageRating: %f",
		m.Movie.Title, m.Movie.ReleaseDate.Year(), m.MovieRating.averageRating)
}

// TopMovieByGenere func to get map of movies by genere
func (rec Recommender) TopMovieByGenere(forGenere movie.Genere) (MovieResult, error) {
	var topRatedMovie movie.Movie

	if rec.Imdb.averageMovieRating == nil {
		return MovieResult{}, errors.New("IMDB is not initialized, Call Initialize before using this method")
	}
	var highestAverageRating AverageRating
	for movieID, movieRating := range rec.Imdb.averageMovieRating {
		movie := rec.Imdb.movieMap[movieID]
		for _, genere := range movie.Genres {
			if genere != forGenere {
				continue
			}
			if (highestAverageRating == AverageRating{}) || (highestAverageRating.averageRating < movieRating.averageRating) {
				highestAverageRating = movieRating
				topRatedMovie = movie
			}
		}
	}
	return MovieResult{Movie: topRatedMovie, MovieRating: highestAverageRating}, nil
}

//TopMovieByYear Return top movie of the Year
func (rec Recommender) TopMovieByYear(forYear int) (MovieResult, error) {
	var topRatedMovie movie.Movie
	imdb := rec.Imdb

	if imdb.averageMovieRating == nil {
		return MovieResult{}, errors.New("IMDB is not initialized, Call Initialize before using this method")
	}
	var highestAverageRating AverageRating
	for movieID, movieRating := range imdb.averageMovieRating {
		movie := imdb.movieMap[movieID]
		if forYear != movie.ReleaseDate.Year() {
			continue
		}
		if (highestAverageRating == AverageRating{}) || (highestAverageRating.averageRating < movieRating.averageRating) {
			highestAverageRating = movieRating
			topRatedMovie = movie
		}
	}
	return MovieResult{Movie: topRatedMovie, MovieRating: highestAverageRating}, nil
}

// MostWatchedGenere ...
func (rec Recommender) MostWatchedGenere() (movie.Genere, error) {
	var mostWatchedGenere movie.Genere
	imdb := rec.Imdb

	if imdb.averageMovieRating == nil {
		return mostWatchedGenere, errors.New("IMDB is not initialized, Call Initialize before using this method")
	}

	genereUserCountMap := make(map[movie.Genere]int64, 0)
	for movieID, movieRating := range imdb.averageMovieRating {
		movie := imdb.movieMap[movieID]
		for _, genere := range movie.Genres {
			if genereRateCount, ok := genereUserCountMap[genere]; !ok {
				genereUserCountMap[genere] = movieRating.ratingCount
			} else {
				genereRateCount = genereRateCount + movieRating.ratingCount
				genereUserCountMap[genere] = genereRateCount
			}
		}
	}

	var highestViewCount = int64(-1)
	for genere, viewCount := range genereUserCountMap {
		if highestViewCount < viewCount {
			mostWatchedGenere = genere
			highestViewCount = viewCount
		}
	}
	return mostWatchedGenere, nil
}

//MostWatchedMovie ...
func (rec Recommender) MostWatchedMovie() (movie.Movie, error) {
	var topRatedMovie int
	imdb := rec.Imdb

	if imdb.averageMovieRating == nil {
		return movie.Movie{}, errors.New("IMDB is not initialized, Call Initialize before using this method")
	}
	var highestViewCount = int64(-1)
	for movieID, movieRating := range imdb.averageMovieRating {
		if highestViewCount < movieRating.ratingCount {
			highestViewCount = movieRating.ratingCount
			topRatedMovie = movieID
		}
	}
	return imdb.movieMap[topRatedMovie], nil
}

//MostActiveUser ...
func (rec Recommender) MostActiveUser() (user.User, error) {
	var activeUser int
	imdb := rec.Imdb

	if imdb.averageMovieRating == nil {
		return user.User{}, errors.New("IMDB is not initialized, Call Initialize before using this method")
	}
	userViewCountMap := make(map[int]int64, 0)
	for _, r := range imdb.ratings {
		if _, ok := userViewCountMap[r.UserID]; !ok {
			userViewCountMap[r.UserID] = 0
		}
		userViewCountMap[r.UserID]++
	}
	var highestViewCount = int64(-1)
	for userID, viewCount := range userViewCountMap {
		if highestViewCount < viewCount {
			highestViewCount = viewCount
			activeUser = userID
		}
	}
	return imdb.users[activeUser], nil
}
