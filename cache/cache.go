package cache

import (
	"encoding/json"
	"io/ioutil"
	"optimus/dirhash"
	"optimus/utils"
	"os"
)

var CacheFilename = "optimus.cache"
var AppCache *Cache

type Cache struct {
	ProjectHash string          `json:"ProjectHash"`
	Commands    []*CommandCache `json:"Commands"`
}

type CommandCache struct {
	Name string `json:"Name"`
	// Should we save command output?
	Output          string `json:"Output"`
	ParentService   string `json:"Parent"`
	RanSuccessfully bool   `json:"RanSuccessfully"`
	Hash            string `json:"Checksum"`
	// DependsOn       []struct {
	// 	Name string
	// 	Hash string
	// }
}

func (cc *CommandCache) UpdateCache(c *Cache) {
	found := false
	for _, cc2 := range c.Commands {
		if cc2.Name == cc.Name && cc2.ParentService == cc.ParentService {
			cc2.RanSuccessfully = cc.RanSuccessfully
			cc2.Hash = cc.Hash
			found = true
		}
	}
	if !found {
		c.Commands = append(c.Commands, cc)
	}
}

func LoadCache() (*Cache, error) {
	if AppCache != nil {
		return AppCache, nil
	}
	p := utils.ProjectRoot()

	if _, err := os.Stat(p + string(os.PathSeparator) + CacheFilename); os.IsNotExist(err) {
		projectHash, err := dirhash.HashDir(p, make([]string, 0))
		if err != nil {
			return nil, err
		}
		c := Cache{
			ProjectHash: projectHash,
			Commands:    make([]*CommandCache, 0),
		}
		return &c, nil
	}

	file, err := os.Open(p + string(os.PathSeparator) + CacheFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteVal, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cache Cache
	err = json.Unmarshal(byteVal, &cache)
	if err != nil {
		return nil, err
	}
	AppCache = &cache

	return &cache, nil
}

func (c *Cache) SaveCache() error {
	projectHash, _ := dirhash.HashDir(utils.ProjectRoot(), make([]string, 0))
	c.ProjectHash = projectHash
	p := utils.ProjectRoot()
	filename := p + string(os.PathSeparator) + CacheFilename

	marchaled, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, marchaled, 0664)
	if err != nil {
		return err
	}

	return nil
}
