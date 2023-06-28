package dirhash

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"io"
	"io/fs"
	"log"
	"os"
	"sort"
)

func DefaultIgnoredPaths() []string {
	return []string{"node_modules", "optimus.cache"}
}

func HashDir(path string, ignore []string) ([]byte, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(path + "/.gitignore"); !os.IsNotExist(err) {
		ReadIgnore(path+"/.gitignore", &ignore)
	}

	filtered := make([]fs.DirEntry, 0)
	for _, de := range files {
		deName := de.Name()
		isIgnored := false
		for _, v := range ignore {
			if deName == v {
				isIgnored = true
				break
			}
		}

		if !isIgnored {
			filtered = append(filtered, de)
		}
	}

	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].Name() > filtered[j].Name()
	})

	hashArray := make([][]byte, 0)

	for _, de := range filtered {
		p := path + string(os.PathSeparator) + de.Name()
		if de.IsDir() {
			newIgnore := make([]string, 0)

			for _, ignoredString := range ignore {
				newIgnore = append(newIgnore, ignoredString)
			}

			hash, err := HashDir(p, newIgnore)
			if err != nil {
				panic(err)
			}
			hashArray = append(hashArray, hash)
		} else {
			hash, err := HashFile(p)
			if err != nil {
				panic(err)
			}
			hashArray = append(hashArray, hash)
		}
	}

	hash := sha1.New()
	hash.Write([]byte(path))
	for _, v := range hashArray {
		_, err := hash.Write([]byte(v))
		if err != nil {
			panic(err)
		}
	}
	res := hash.Sum(nil)

	return res, nil
}

func HashFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return nil, err
	}
	sum := hash.Sum(nil)

	return sum, nil
}

func ReadIgnore(path string, ignore *[]string) error {
	if ignore == nil {
		panic("ignore slice is empty")
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*ignore = append(*ignore, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
