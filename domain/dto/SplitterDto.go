package dto

import (
	"sort"
	"strings"
)

type SplitterDto struct {
	Split []string `json:"split"`
	Count int      `json:"count"`
}

func (g SplitterDto) SplitterOrdering(split string) SplitterDto {
	s := strings.Split(split, ",")
	sort.Strings(s)
	return SplitterDto{
		Split: s,
		Count: len(s),
	}
}
