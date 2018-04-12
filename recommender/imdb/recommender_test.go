package imdb

import (
	"net/url"
	"testing"
	"time"

	"github.com/marshal003/exercises/recommender/rating"

	"github.com/marshal003/exercises/recommender/movie"
)

func TestMostWatchedMovie(t *testing.T) {
	url, _ := url.Parse("http://test.com")
	movieMap := map[int]movie.Movie{
		1001: movie.Movie{MovieID: 1001, Title: "Movie1", ReleaseDate: time.Now(), VideoReleaseDate: time.Now(), ImdbURL: *url, Genres: []movie.Genere{movie.ACTION, movie.ADVENTURE}},
		1002: movie.Movie{MovieID: 1002, Title: "Movie2", ReleaseDate: time.Now(), VideoReleaseDate: time.Now(), ImdbURL: *url, Genres: []movie.Genere{movie.Drama, movie.ACTION}},
		1003: movie.Movie{MovieID: 1004, Title: "Movie3", ReleaseDate: time.Now(), VideoReleaseDate: time.Now(), ImdbURL: *url, Genres: []movie.Genere{movie.ACTION, movie.ADVENTURE, movie.Comedy}},
	}
	ratings := []rating.UserRating{
		rating.UserRating{UserID: 201, MovieID: 1001, Rating: 4, RatedOn: time.Now()},
		rating.UserRating{UserID: 202, MovieID: 1003, Rating: 3, RatedOn: time.Now()},
		rating.UserRating{UserID: 204, MovieID: 1001, Rating: 2, RatedOn: time.Now()},
		rating.UserRating{UserID: 201, MovieID: 1003, Rating: 3, RatedOn: time.Now()},
		rating.UserRating{UserID: 202, MovieID: 1001, Rating: 4, RatedOn: time.Now()},
		rating.UserRating{UserID: 202, MovieID: 1003, Rating: 4, RatedOn: time.Now()},
		rating.UserRating{UserID: 205, MovieID: 1001, Rating: 4, RatedOn: time.Now()},
		rating.UserRating{UserID: 204, MovieID: 1001, Rating: 4, RatedOn: time.Now()},
	}

	imdb := CreateIMDB(movieMap, ratings, nil)
	recommender := Recommender{Imdb: imdb}
	result, err := recommender.MostWatchedMovie()
	if err != nil {
		t.Error("It shouldn't have thrown exception")
	}
	if result.MovieID != 1001 {
		t.Errorf("1001 should be the mostwatched movie, but recommender says : %d", result.MovieID)
	}
}

func TestMostWatchedGenere(t *testing.T) {
	url, _ := url.Parse("http://test.com")
	movieMap := map[int]movie.Movie{
		1001: movie.Movie{MovieID: 1001, Title: "Movie1", ReleaseDate: time.Now(), VideoReleaseDate: time.Now(), ImdbURL: *url, Genres: []movie.Genere{movie.ACTION, movie.ADVENTURE}},
		1002: movie.Movie{MovieID: 1002, Title: "Movie2", ReleaseDate: time.Now(), VideoReleaseDate: time.Now(), ImdbURL: *url, Genres: []movie.Genere{movie.Drama, movie.ACTION}},
		1003: movie.Movie{MovieID: 1004, Title: "Movie3", ReleaseDate: time.Now(), VideoReleaseDate: time.Now(), ImdbURL: *url, Genres: []movie.Genere{movie.ACTION, movie.ADVENTURE, movie.Comedy}},
	}
	ratings := []rating.UserRating{
		rating.UserRating{UserID: 201, MovieID: 1001, Rating: 4, RatedOn: time.Now()},
		rating.UserRating{UserID: 202, MovieID: 1003, Rating: 3, RatedOn: time.Now()},
		rating.UserRating{UserID: 204, MovieID: 1001, Rating: 2, RatedOn: time.Now()},
		rating.UserRating{UserID: 201, MovieID: 1003, Rating: 3, RatedOn: time.Now()},
		rating.UserRating{UserID: 202, MovieID: 1001, Rating: 4, RatedOn: time.Now()},
		rating.UserRating{UserID: 202, MovieID: 1003, Rating: 4, RatedOn: time.Now()},
		rating.UserRating{UserID: 205, MovieID: 1001, Rating: 4, RatedOn: time.Now()},
		rating.UserRating{UserID: 204, MovieID: 1001, Rating: 4, RatedOn: time.Now()},
	}

	imdb := CreateIMDB(movieMap, ratings, nil)
	recommender := Recommender{Imdb: imdb}
	result, err := recommender.MostWatchedGenere()
	if err != nil {
		t.Error("It shouldn't have thrown exception")
	}
	if result != movie.ACTION {
		t.Errorf("It should have returned ACTION but has returned: %v", result)
	}
}
