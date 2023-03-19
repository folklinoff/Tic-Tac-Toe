package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	searchresults "github.com/folklinoff/tic-tac-toe/search-results"
)

func tic_tac_toe_API(request string) (answer []string) {
	s := make([]string, 3)

	if response, err := http.Get(request); err != nil {
		s[0] = "Tic-tac-toe bot does not respond"
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		sr := &searchresults.SearchResults{}
		if err = json.Unmarshal(contents, sr); err != nil {
			s[0] = "Something going wrong, try to change your question"
		}

		if !sr.Ready {
			s[0] = "Something going wrong, try to change your question"
		}

		for i := range sr.Results {
			s[i] = sr.Results[i].URL
		}
	}

	return s
}
