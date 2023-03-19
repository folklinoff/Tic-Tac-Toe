package main

import (
	searchresults "github.com/folklinoff/tic-tac-toe/search-results"
)

func main() {
	sr := searchresults.SearchResults{}
	sr.UnmarshalJSON([]byte{})
}
