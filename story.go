package main

import (
	"encoding/json"
	"io/ioutil"
)

type storyJson struct {
	Name string
	Pics []string
}

type Story struct {
	Name string
	Pics []*Pic
}

func NewStory(path string, pics []*Pic) (*Story, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	storyJson := new(storyJson)
	if err := json.Unmarshal(b, storyJson); err != nil {
		return nil, err
	}

	story := new(Story)
	story.Name = storyJson.Name
	story.Pics, err = namesToPics(storyJson.Pics, pics)

	return story, err
}
