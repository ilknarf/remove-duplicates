package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	cPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	folder, err := ioutil.ReadDir(cPath)

	if err != nil {
		log.Fatal(err)
	}

	set := NewHashset()
	c := 0

	for _, f := range folder {
		name := f.Name()
		fPath := path.Join(cPath, name)

		file, err := os.Open(fPath)
		hash := md5.New()

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		io.Copy(hash, file)

		h := hash.Sum(nil)
		hString := string(h)

		if set.Contains(hString) {
			os.Remove(name)
			c++
		} else {
			set.Add(hString)
		}
	}

	fmt.Printf("%d files removed\n", c)
}
