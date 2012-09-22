package main

import (
	"fmt"
	"github.com/afajl/log"
	"path/filepath"
)

type Creator interface {
	Create() error
}

type FsCreator struct {
	sourceDir string
	repr      Representer
}

func NewFsCreator(conf *Config, repr Representer) Creator {
	c := new(FsCreator)
	c.sourceDir = conf.SourceDir
	c.repr = repr
	return c
}

func (fc *FsCreator) Create() error {
	var err error
	var pics []*Pic
	var galleries []*Gallery
	var stories []*Story

	pics, err = parsePics(fc.sourceDir)
	if err != nil {
		return err
	}

	galleries, err = parseGalleries(fc.sourceDir, pics)
	if err != nil {
		return err
	}

	stories, err = parseStories(fc.sourceDir, pics)
	if err != nil {
		return err
	}

    for _, pic := range pics {
        if err := fc.repr.AddPic(pic); err != nil {
            return err
        }
    }

	for _, gallery := range galleries {
        if err := fc.repr.AddGallery(gallery); err != nil {
            return err
        }
        log.Trace("fc.repr", fc.repr)
    }
    for _, story := range stories {
        if err := fc.repr.AddStory(story); err != nil {
            return err
        }
    }

    log.Trace("fc.repr", fc.repr)

    return fc.repr.Reprezent()
}

func parsePics(sourceDir string) ([]*Pic, error) {
	paths, err := filepath.Glob(sourceDir + "/pics/*.png")
	if err != nil {
		return nil, err
	}

	var pics = make([]*Pic, 0, len(paths))
	for _, path := range paths {
		log.Debugf("parsing pic %v", path)
		pic, err := NewPic(path)
		if err != nil {
			return nil, err
		}
		pics = append(pics, pic)
	}
	return pics, nil
}

func parseGalleries(sourceDir string, pics []*Pic) ([]*Gallery, error) {
	paths, err := filepath.Glob(sourceDir + "/galleries/*.json")
	if err != nil {
		return nil, err
	}
	var galleries = make([]*Gallery, 0, len(paths))

	for _, path := range paths {
		log.Debugf("parsing gallery %v", path)
		gallery, err := NewGallery(path, pics)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", path, err)
		}
		galleries = append(galleries, gallery)
	}
	return galleries, nil
}

func parseStories(sourceDir string, pics []*Pic) ([]*Story, error) {
	paths, err := filepath.Glob(sourceDir + "/stories/*.json")
	if err != nil {
		return nil, err
	}
	var stories = make([]*Story, 0, len(paths))

	for _, path := range paths {
		log.Debugf("parsing story %v", path)
		story, err := NewStory(path, pics)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", path, err)
		}
		stories = append(stories, story)
	}
	return stories, nil
}
