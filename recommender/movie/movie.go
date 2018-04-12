package movie

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/marshal003/exercises/recommender/utils"
)

// Genere custom type for Genere
type Genere int

// Genre Constants
const (
	UNKNOWN Genere = iota
	ACTION
	ADVENTURE
	ANIMATION
	CHILDRENS
	Comedy
	Crime
	Documentary
	Drama
	Fantasy
	FilmNoir
	Horror
	Musical
	Mystery
	Romance
	SciFi
	Thriller
	War
	Western
)

func (g Genere) String() string {
	generes := []string{
		"UNKNOWN",
		"ACTION",
		"DVENTURE",
		"ANIMATION",
		"CHILDRENS",
		"Comedy",
		"Crime",
		"Documentary",
		"Drama",
		"Fantasy",
		"FilmNoir",
		"Horror",
		"Musical",
		"Mystery",
		"Romance",
		"SciFi",
		"Thriller",
		"War",
		"Western",
	}
	if g < UNKNOWN || g > Western {
		return ""
	}
	return generes[g]
}

//GenereFromIndex Utility method to return genere from its index
func GenereFromIndex(index int) (Genere, error) {
	generes := []Genere{
		UNKNOWN,
		ACTION,
		ADVENTURE,
		ANIMATION,
		CHILDRENS,
		Comedy,
		Crime,
		Documentary,
		Drama,
		Fantasy,
		FilmNoir,
		Horror,
		Musical,
		Mystery,
		Romance,
		SciFi,
		Thriller,
		War,
		Western,
	}
	if index < int(UNKNOWN) || index > int(Western) {
		return UNKNOWN, errors.New("Unknown Genere Index")
	}
	return generes[index], nil
}

//Movie data structure for Movie
type Movie struct {
	MovieID          int
	Title            string
	ReleaseDate      time.Time
	VideoReleaseDate time.Time
	ImdbURL          url.URL
	Genres           []Genere
}

//ReadMovies reads movie.data file and returns map of movie and its Ids
func ReadMovies(dataFile utils.DataFile) map[int]Movie {
	channel := make(chan []string, 0)
	go dataFile.AsyncDataFileReader(channel)

	movies := make(map[int]Movie)
	for tokens := range channel {
		movie, err := parseMovie(tokens)
		if err != nil {
			continue
		}
		movies[movie.MovieID] = movie
	}
	return movies
}

func parseMovie(tokens []string) (Movie, error) {
	var movie Movie
	if len(tokens) != 24 {
		return movie, errors.New("Tokens count should be 24")
	}
	movieID, err := strconv.Atoi(strings.Trim(tokens[0], " "))
	if err != nil {
		return movie, errors.New("Unable to parse movieId")
	}
	title := strings.Trim(tokens[1], " ")
	imdbURL, _ := url.Parse(strings.Trim(tokens[4], " "))
	releaseDate, _ := parseDate(tokens[2])
	videoReleaseDate, _ := parseDate(tokens[2])
	generes := parseGenere(tokens[5:23])
	return Movie{
		MovieID:          movieID,
		Title:            title,
		ReleaseDate:      releaseDate,
		VideoReleaseDate: videoReleaseDate,
		ImdbURL:          *imdbURL,
		Genres:           generes,
	}, nil
}

func parseGenere(tokens []string) []Genere {
	generes := make([]Genere, 0)
	for index, token := range tokens {
		val, err := strconv.Atoi(strings.Trim(token, " "))
		if err != nil || val != 1 {
			continue
		}

		if genere, err := GenereFromIndex(index); err == nil {
			generes = append(generes, genere)
		}
	}
	return generes
}

func parseDate(date string) (time.Time, error) {
	layout := "01-Jan-2006"
	return time.Parse(layout, strings.Trim(date, " "))
}
