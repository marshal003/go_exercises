package main

import (
	"fmt"

	"github.com/marshal003/exercises/recommender/imdb"
	"github.com/marshal003/exercises/recommender/movie"
	"github.com/marshal003/exercises/recommender/rating"
	"github.com/marshal003/exercises/recommender/user"
	"github.com/marshal003/exercises/recommender/utils"
)

func main() {

	dataFilePath := "/Users/300027725/go/src/github.com/marshal003/exercises/recommender/data"
	movieDataFile := dataFilePath + "/movie.data"
	movieMap := movie.ReadMovies(utils.DataFile{Filepath: movieDataFile, Separator: "|"})
	fmt.Printf("Total Movies Read: %d\n", len(movieMap))
	ratingDataFile := dataFilePath + "/ratings.data"
	ratings := rating.ReadUserRatings(utils.DataFile{Filepath: ratingDataFile, Separator: "\t"})
	fmt.Printf("Total Ratings Read: %d\n", len(ratings))

	userDataFile := dataFilePath + "/user.data"
	users := user.ReadUsers(utils.DataFile{Filepath: userDataFile, Separator: "|"})

	im := imdb.CreateIMDB(movieMap, ratings, users)
	recommender := imdb.Recommender{Imdb: im}

	topMovieByGenere, _ := recommender.TopMovieByGenere(movie.Horror)
	fmt.Printf("\nTopRatedMovieByGenere: %v\n", topMovieByGenere)

	mostWatchedMovie, _ := recommender.MostWatchedMovie()
	fmt.Printf("\nMostWatchedMovie: %v\n", mostWatchedMovie)

	mostWatchedGenere, _ := recommender.MostWatchedGenere()
	fmt.Printf("\nMostWatchedGenere: %v\n", mostWatchedGenere)

	mostActiveUser, _ := recommender.MostActiveUser()
	fmt.Printf("TopRatedMovie: %v\n", mostActiveUser)
}
