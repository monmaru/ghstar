package main

import (
	"strings"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		in   string
		want bool
	}{
		{"", true},
		{"a", false},
	}

	for _, tt := range tests {
		assert.Equal(t, isEmpty(tt.in), tt.want, "they should be equal")
	}
}

func TestShow(t *testing.T) {
	dummy := "dummy"
	repo := &github.Repository{HTMLURL: &dummy}
	show(repo)
	repo.Description = &dummy
	repo.Language = &dummy
	c := 1
	repo.StargazersCount = &c
	repo.ForksCount = &c
	repo.UpdatedAt = &github.Timestamp{}
	show(repo)
}

func TestIsTargetLang(t *testing.T) {
	lang := "Go"
	var tests = []struct {
		in1  string
		in2  *string
		want bool
	}{
		{"", nil, true},
		{"", &lang, true},
		{lang, &lang, true},
		{strings.ToLower(lang), &lang, true},
		{strings.ToUpper(lang), &lang, true},
		{"goo", &lang, false},
	}

	for _, tt := range tests {
		repo := &github.Repository{Language: tt.in2}
		assert.Equal(t, isTargetLang(repo, tt.in1), tt.want, "they should be equal")
	}
}
