package main 

import (
 	"os"
	"os/exec"
	"path/filepath"
	"strconv"   
    "fmt"
)

func picPath(targetDir, id string) string {
	return filepath.Join(targetDir, "pics", id)
}

func hasPic(targetDir, id string) bool {
	_, err := os.Stat(picPath(targetDir, id))
	return !os.IsNotExist(err)
}

func writePic(targetDir string, alternates map[string]int, pic *Pic) error {
	path := picPath(targetDir, pic.id)
	if hasPic(targetDir, pic.id) {
		return nil
	}
	if err := os.MkdirAll(path, os.FileMode(0755)); err != nil {
		return err
	}
	if pic.isWide {
		if f, err := os.Create(path + "/isWide"); err != nil {
			return err
		} else {
			f.Close()
		}
	}
	return writeAlternates(path, alternates, pic)
}

func writeAlternates(picDir string, alternates map[string]int, pic *Pic) error {
	for key, width := range alternates {
		if pic.isWide {
			width *= 2
		}
		path := filepath.Join(picDir, key+".jpg")
		err := writeAlternate(pic.originalPath, path, width)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeAlternate(original, target string, width int) error {
	cmd := exec.Command("convert", original, "-strip", "-resize",
		strconv.Itoa(width), "-quality", "75", target)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error creating alternate %q: %s", err, out)
	}
	return nil
}

