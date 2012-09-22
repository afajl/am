package main

import (
	"fmt"
	"github.com/afajl/log"
	"path/filepath"

)

type StaticRepresenter struct {
	picAlternates map[string]int
	targetDir     string
	pics          []*Pic
	galleries     []*Gallery
	stories       []*Story
}

func NewStatic(conf *Config) Representer {
	repr := new(StaticRepresenter)
	repr.targetDir = conf.TargetDir
	repr.picAlternates = conf.PicAlternates

	repr.pics = []*Pic{}
	repr.galleries = []*Gallery{}
	repr.stories = []*Story{}
	return repr
}

func (r *StaticRepresenter) AddPic(pic *Pic) error {
	err := writePic(r.targetDir, r.picAlternates, pic)
	r.pics = append(r.pics, pic)
	return err
}

func (r *StaticRepresenter) AddGallery(gallery *Gallery) error {
    log.Tracef("AddGallery(%q)", gallery)
    for i, pic := range gallery.Pics {
		if pic.isWide && ( (i+1)%2 == 0 ) {
			return fmt.Errorf("pic #%d in gallery %q (%q) is wide and must be on an uneven position", i+1, pic.originalPath, gallery.Name)
		}
	}
    log.Trace("appending gallery")
	r.galleries = append(r.galleries, gallery)
    log.Trace("appended gallery", r.galleries)
	return nil
}

func (r *StaticRepresenter) AddStory(story *Story) error {
	r.stories = append(r.stories, story)
	return nil
}

func (r *StaticRepresenter) Reprezent() error {
    log.Trace("in Reprezent", r.galleries)
    for _, gallery := range r.galleries {
        log.Debugf("writing gallery %q", gallery) 
        if err := writeGallery(r.targetDir, gallery); err != nil {
            return err
        }
    }
	return nil
}

func writeGallery(targetDir string, gallery *Gallery) error {
    file := gallery.Path
    if file == "/" {
        file = "index"
    }
    return WriteTempl("gallery.html", filepath.Join(targetDir, file + ".html"), gallery)
}
