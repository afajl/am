package main

type Representer interface {
	AddPic(pic *Pic) error
	AddGallery(gallery *Gallery) error
	AddStory(story *Story) error
	Reprezent() error
}
