package models

import (
	//"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	_, err := Connect()
	if err != nil {
		t.Errorf("DB connect Error = %v", err)
	}
}

func TestSearchSQL(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "SELECT author, title, beginning FROM aozora WHERE CONCAT(author, title, beginning) LIKE ?;"},
		{2, "SELECT author, title, beginning FROM aozora WHERE CONCAT(author, title, beginning) LIKE ? AND CONCAT(author, title, beginning) LIKE ?;"},
	}
	for _, c := range cases {
		got := getSearchSQL(c.in)
		if got != c.want {
			t.Errorf("getSearchSQL(%v) == \n %v, \n want %v", c.in, got, c.want)
		}
	}
}

func TestSearch(t *testing.T) {
	d, _ := Connect()
	d.Search("は")
	d.Search("ある　は")
}