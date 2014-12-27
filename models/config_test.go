package models

import (
	"reflect"
	"testing"
)

func TestConfiguration(t *testing.T) {
	got := getConfig()
	want := config{
		Title: "Model Configuration",
		Server: serverConfig{
			Host: "localhost",
			Port: "3306",
		},
		DB: dbConfig{
			Name: "mydb",
			User: "root",
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Configuration == %v, want %v", got, want)
	}
}
