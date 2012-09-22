package main

import (
	"fmt"
	"github.com/afajl/log"
	"hash/adler32"
	"image"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
)

type Pic struct {
	originalPath  string
	id            string
	isWide        bool
	height, width int
}

func NewPic(path string) (*Pic, error) {
	pic := new(Pic)
	pic.originalPath = path

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	pic.id = picHash(f)
	f.Seek(0, 0)
	pic.height, pic.width, err = picSize(f)
	if err != nil {
		return nil, err
	}

	pic.isWide = pic.height < pic.width

	return pic, err
}

func picHash(picf io.Reader) string {
	hash := adler32.New()
	if _, err := io.Copy(hash, picf); err != nil {
		log.Panic(err)
	}
	return fmt.Sprintf("%x", hash.Sum32())
}

func picSize(picf io.Reader) (int, int, error) {
	m, _, err := image.Decode(picf)
	if err != nil {
		return 0, 0, err
	}
	r := m.Bounds()
	return r.Dy(), r.Dx(), nil
}

/*func picIsWide(pic *Pic, portraitRatio float64) (isWide bool, err error) {*/
	/*landscapeRatio := portraitRatio / 2*/
	/*ratio := float64(pic.height) / float64(pic.width)*/

	/*if math.Mod(ratio, portraitRatio) < 0.01 {*/
		/*isWide = false*/
	/*} else if math.Mod(ratio, landscapeRatio) < 0.01 {*/
		/*isWide = true*/
	/*} else {*/
		/*err = fmt.Errorf("pic ratio is wrong: w:%v h:%v, ratio:%v", pic.height, pic.width, ratio)*/
	/*}*/
	/*return*/
/*}*/

func findPic(name string, pics []*Pic) *Pic {
	for _, pic := range pics {
		if filepath.Base(pic.originalPath) == name {
			return pic
		}
	}
	return nil
}

func namesToPics(picnames []string, pics []*Pic) (res []*Pic, err error) {
	res = make([]*Pic, 0, len(picnames))

	for _, picname := range picnames {
		pic := findPic(picname, pics)
		if pic == nil {
			err = fmt.Errorf("no pic found named %q", picname)
			return
		}
		res = append(res, pic)
	}
	return
}
