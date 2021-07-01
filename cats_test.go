package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
Тесты получения картинки по заданным параметрам
*/
func TestGetCats(t *testing.T) {
	tests := []struct {
		name string
		page int
		post int
		url  string
		err  error
	}{
		{
			name: "success",
			page: 3199,
			post: 7,
			url:  "http://img1.joyreactor.cc/pics/post/%D0%9A%D0%BE%D0%BC%D0%B8%D0%BA%D1%81%D1%8B-%D0%BA%D0%BE%D1%82-%D0%91%D1%8D%D1%82%D0%BC%D0%B5%D0%BD-DC-Comics-5039897.jpeg",
			err:  nil,
		},
		{
			name: "pageError",
			page: 0,
			post: 10,
			url:  "",
			err:  ErrInvalidPageNumber,
		},
		{
			name: "postError",
			page: 1,
			post: 11,
			url:  "",
			err:  ErrInvalidPostNumber,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, err := GetCats(tt.page, tt.post, "%D0%BA%D0%BE%D1%82%D1%8D", 4000)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.url, gotURL)
		})
	}
}

/**
Тесты получения документа по URL
*/
func TestGetDocumentFromURL(t *testing.T) {
	tests := []struct {
		name   string
		joyUrl string
		err    error
	}{
		{
			name:   "success",
			joyUrl: "http://joyreactor.cc/tag/котэ",
			err:    nil,
		},
		{
			name:   "invalidUrl",
			joyUrl: "http://test",
			err:    ErrInvalidUrl,
		},
		{
			name:   "pageNotFound",
			joyUrl: "http://joyreactor.cc/test",
			err:    ErrPageNotFound,
		},
	}

	for _, getDocumentTest := range tests {
		t.Run(getDocumentTest.name, func(t *testing.T) {
			_, err := GetDocumentFromURL(getDocumentTest.joyUrl)
			assert.Equal(t, getDocumentTest.err, err)
		})
	}
}

/**
Тесты получения случайного поста с котами
*/
func TestGetRandomCats(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		err  error
	}{
		{
			name: "success",
			tag:  "котэ",
			err:  nil,
		},
		{
			name: "noMaxPage",
			tag:  "%!invalidTag%!",
			err:  ErrNoMaxPages,
		},
	}

	for _, getRandomCats := range tests {
		t.Run(getRandomCats.name, func(t *testing.T) {
			_, err := GetRandomCats(getRandomCats.tag)
			assert.Equal(t, getRandomCats.err, err)
		})
	}
}

/**
Тесты получения количества страниц
*/
func TestGetPagesCount(t *testing.T) {
	tests := []struct {
		name string
		tag  string
		err  error
	}{
		{
			name: "success",
			tag:  "котэ",
			err:  nil,
		},
		{
			name: "connectionFailed",
			tag:  "%!invalidTag%!",
			err:  ConnectionFailed,
		},
		{
			name: "noMaxPage",
			tag:  "joyre",
			err:  ErrNoMaxPages,
		},
	}

	for _, getPagesCount := range tests {
		t.Run(getPagesCount.name, func(t *testing.T) {
			_, err := GetPagesCount(getPagesCount.tag)
			assert.Equal(t, getPagesCount.err, err)
		})
	}
}
