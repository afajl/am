package main

import (
	"github.com/afajl/log"
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

const buildDir = "testdata/build"

func init() {
	log.OutLevel = log.TRACE
    conf = NewConfig()
}

func TestWritePic(t *testing.T) {
	picId := "testid"

	os.MkdirAll(buildDir, os.FileMode(0755))
	defer os.RemoveAll(buildDir)

	pic := &Pic{originalPath: "testdata/pics/1.png", id: picId}

	assert.Equal(t, hasPic(buildDir, picId), false)
	err := writePic(buildDir, map[string]int{"a": 10, "b": 20}, pic)
	assert.Equal(t, err, nil)
	assert.Equal(t, hasPic(buildDir, picId), true)

	_, err = os.Stat(picPath(buildDir, picId) + "/a.jpg")
	assert.Equal(t, os.IsNotExist(err), false)

	_, err = os.Stat(picPath(buildDir, picId) + "/b.jpg")
	assert.Equal(t, os.IsNotExist(err), false)
}


func TestWriteGallery(t *testing.T) {
	os.MkdirAll(buildDir, os.FileMode(0755))
	defer os.RemoveAll(buildDir)

    gallery := &Gallery{Name: "/", Pics: []*Pic{}}
    if err := writeGallery(buildDir, gallery); err != nil {
        t.Fatal(err)
    }
    if _, err := os.Stat(buildDir + "/index.html"); !os.IsNotExist(err) {
        t.Fatal("gallery path of / should create index.html") 
    }

}
