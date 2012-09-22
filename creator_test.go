package main

import (
	"github.com/afajl/log"
	"github.com/bmizerany/assert"
	"testing"
)

func init() {
	log.OutLevel = log.TRACE
}

func TestParsers(t *testing.T) {
	pics, err := parsePics("testdata")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(pics), 2, pics)
}

func TestParseGalleries(t *testing.T) {
	pics := []*Pic{&Pic{originalPath: "/knas/1.jpg", isWide: true},
		&Pic{originalPath: "/hoppsan/2.jpg", isWide: false}}
	gals, err := parseGalleries("testdata", pics)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(gals), 2, gals)
}

func TestParseStories(t *testing.T) {
	pics := []*Pic{&Pic{originalPath: "/knas/1.jpg"},
		&Pic{originalPath: "/hoppsan/2.jpg"}}
	stories, err := parseStories("testdata", pics)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(stories), 2, stories)
}
