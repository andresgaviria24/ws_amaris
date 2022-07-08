package dto

import "fmt"

type FriendlyChainsDto struct {
	TextOne string `json:"textOne"`
	TextTwo string `json:"textTwo"`
}

func (g FriendlyChainsDto) Message(message bool) string {

	if len(g.TextOne) == 0 || len(g.TextTwo) == 0 {
		return "No text can be empty"
	}

	if message {
		return fmt.Sprintf("%s y %s son amigas", g.TextOne, g.TextTwo)
	}

	return fmt.Sprintf("%s y %s NO son amigas", g.TextOne, g.TextTwo)
}
