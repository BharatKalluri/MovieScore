package getratings

import (
	"encoding/json"
	"fmt"
	"github.com/ttacon/chalk"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

//PrettyPrinter Print function which prints all the information from all the modules
func PrettyPrinter(MovieName string, year string) {
	RtRating := RtScraper(MovieName, year)
	ImdbRatings := GetImdbRatings(MovieName)
	IntRtRatings, err := strconv.Atoi(RtRating)
	if IntRtRatings == -1 && len(ImdbRatings.Title) == 0 {
		fmt.Println("The Movie Does not seem to exist!")
		fmt.Println("Tip: If you are using spaces in your film name, enclose the movie name in double quotes!")
	} else {
		fmt.Println(chalk.Magenta, "Movie Name: "+ImdbRatings.Title)
		fmt.Println(chalk.Magenta, "Director: "+ImdbRatings.Director)
		fmt.Println(chalk.Magenta, "Cast: "+ImdbRatings.Actors)
		fmt.Println(chalk.Magenta, "Year: "+ImdbRatings.Year)
		fmt.Println(chalk.Magenta, "Released: "+ImdbRatings.Released)
		fmt.Println(chalk.Magenta, "Rated: "+ImdbRatings.Rated)
		fmt.Println(chalk.Magenta, "Genre: "+ImdbRatings.Genre)
		fmt.Println(chalk.Magenta, "Poster: "+ImdbRatings.Poster)
		fmt.Println(chalk.Magenta, "Metascore Rated: "+ImdbRatings.Metascore)
		fmt.Println(chalk.Magenta, "Awards: "+ImdbRatings.Awards)
		fmt.Println(chalk.Magenta, "Plot: "+ImdbRatings.Plot)
		fmt.Println(" Ratings from IMDB and Rotten Tomatoes---")
		fmt.Println(chalk.Magenta, chalk.Underline.TextStyle("IMDB Rating: "+ImdbRatings.ImdbRating))
		if IntRtRatings == -1 && err == nil {
			fmt.Println(chalk.Red, "There seems to be a problem with rt, try with the year argument!")
		} else if IntRtRatings > 60 && err == nil {
			fmt.Println(chalk.Red, chalk.Underline.TextStyle("Rotten Tomatoes Rating: "+RtRating+"% (Certified Fresh!)"))
		} else {
			fmt.Println(chalk.Green, chalk.Underline.TextStyle("Rotten Tomatoes Rating: "+RtRating+"% (Rotten!)"))
		}
	}
}

//ASCIIPoster generates the ASCII POSTER!!
func ASCIIPoster() {
	fmt.Println(chalk.Cyan, `
------------------------------------------------------
  __  __            _         _____                    
 |  \/  |          (_)       / ____|                   
 | \  / | _____   ___  ___  | (___   ___ ___  _ __ ___ 
 | |\/| |/ _ \ \ / / |/ _ \  \___ \ / __/ _ \| '__/ _ \
 | |  | | (_) \ V /| |  __/  ____) | (_| (_) | | |  __/
 |_|  |_|\___/ \_/ |_|\___| |_____/ \___\___/|_|  \___|
------------------------------------------------------
	`)
}

//GetJSON Function which takes the url and the target as arguments for parsing json
func GetJSON(url string, target interface{}) error {
	if strings.Contains(url, "omdb") {
		// This is a free key, and throttled to 10000 req per day
		url = url + "&apikey=e389610c"
	}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
