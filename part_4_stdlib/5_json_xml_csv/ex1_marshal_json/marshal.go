package main

import (
	"encoding/json"
	"fmt"
	"time"
	"strings"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, '"')
	dur := time.Duration(d).String()
	if strings.HasSuffix(dur, "m0s") {
        dur = dur[:len(dur)-2]
    }
    if strings.HasSuffix(dur, "h0m") {
        dur = dur[:len(dur)-2]
    }
	dur = strings.TrimPrefix(dur, "0h")
	b = append(b, []byte(dur)...)
	b = append(b, '"')
	return b, nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0)
	b = append(b, '"')
	i := 0
	for i < int(r) {
		b = append(b, 226, 152, 133)
		i ++
	}
	for i < 5 {
		b = append(b, 226,152, 134)
		i++
	}
	b = append(b, '"')
	return b, nil
}

// Movie описывает фильм
type Movie struct {
	Title string
	Year int
	Director string
	Genres []string
	Duration Duration
	Rating Rating
}

// MarshalMovies кодирует фильмы в JSON.
// - если indent = 0 - использует json.Marshal
// - если indent > 0 - использует json.MarshalIndent
//   с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	var jsonb []byte
	var err error
	if indent == 0 {
		jsonb, err = json.Marshal(movies)
	} else {
		indentStr := strings.Repeat(" ", indent)
		jsonb, err = json.MarshalIndent(movies, "", indentStr)
	}
	return string(jsonb), err
}

// конец решения

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 36*time.Minute),
		Rating:   4,
	}

	s, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(s)
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
