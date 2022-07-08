package service

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"ws_amaris/domain/dto"
)

func TestSplitter(t *testing.T) {

	service := InitServiceImpl()

	cases := []struct {
		name           string
		text           string
		expectedOutput dto.SplitterDto
	}{
		{
			name: "When entering names with a comma it returns an array of names separated and sorted ",
			text: "Laura,Andres,Daniela",
			expectedOutput: dto.SplitterDto{
				Split: []string{"Andres", "Daniela", "Laura"},
				Count: 3,
			},
		}, {
			name: "When entering one name it returns an array of one name sorted ",
			text: "Laura",
			expectedOutput: dto.SplitterDto{
				Split: []string{"Laura"},
				Count: 1,
			},
		},
	}

	for _, tt := range cases {

		output := service.Splitter(tt.text)

		assert.Equal(t, tt.expectedOutput, output)
	}
}

func TestFriendlyChains(t *testing.T) {

	service := InitServiceImpl()

	cases := []struct {
		name           string
		text           dto.FriendlyChainsDto
		expectedOutput string
	}{
		{
			name: "When you enter two words that are friends",
			text: dto.FriendlyChainsDto{
				TextOne: "tokyo",
				TextTwo: "kyoto",
			},
			expectedOutput: "tokyo y kyoto son amigas",
		}, {
			name: "When you enter two words that are friends",
			text: dto.FriendlyChainsDto{
				TextOne: "hola",
				TextTwo: "olah",
			},
			expectedOutput: "hola y olah son amigas",
		}, {
			name: "When you enter empty words that can be empty",
			text: dto.FriendlyChainsDto{
				TextOne: "",
				TextTwo: "",
			},
			expectedOutput: "No text can be empty",
		},
		{
			name: "When you enter two words that are not friends",
			text: dto.FriendlyChainsDto{
				TextOne: "hola",
				TextTwo: "amigos",
			},
			expectedOutput: "hola y amigos NO son amigas",
		},
	}

	for _, tt := range cases {

		output := service.FriendlyChains(tt.text)

		assert.Equal(t, tt.expectedOutput, output)
	}
}
