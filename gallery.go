package main

import (
	"encoding/json"
	"io/ioutil"
)

type galleryJson struct {
	Name, Path string
	Pics       []string
}

type Gallery struct {
	Name, Path string
	Pics       []*Pic
}

func NewGallery(path string, pics []*Pic) (*Gallery, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	galleryJson := new(galleryJson)
	if err := json.Unmarshal(b, galleryJson); err != nil {
		return nil, err
	}

	gallery := new(Gallery)
	gallery.Name = galleryJson.Name
	gallery.Path = galleryJson.Path
	gallery.Pics, err = namesToPics(galleryJson.Pics, pics)

	return gallery, err
}
