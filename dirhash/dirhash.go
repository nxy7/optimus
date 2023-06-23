package dirhash

import (
	"bufio"
	"crypto"
	"crypto/md5"
	"io"
	"io/fs"
	"log"
	"os"
	"sort"
)

func HashDir(path string, ignore []string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", err
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
			}
		}

		if !isIgnored {
			filtered = append(filtered, de)
		}
	}
	files = filtered

	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Name() > files[j].Name()
	})

	hashArray := make([]string, 0)

	for _, de := range files {
		p := path + string(os.PathSeparator) + de.Name()
		if de.IsDir() {
			newIgnore := make([]string, 0)
			copy(newIgnore, ignore)
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

	hash := crypto.SHA256.New()
	hash.Write([]byte(path))
	for _, v := range hashArray {
		_, err := hash.Write([]byte(v))
		if err != nil {
			panic(err)
		}
	}
	res := string(hash.Sum(nil))

	// return it

	return res, nil
}

func HashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	sum := string(hash.Sum(nil))

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
