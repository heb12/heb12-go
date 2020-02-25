// Package osistool manages, parses, and otherwise deals with Bibles in the OSIS format (specifically the ones by gratis.bible)
package osistool

import (
	"os"
	"path/filepath"
	"strings"
)

// DocumentDir contains the location of the documents in relation to the DataDir
//const DocumentDir string = "/bibles/osis"

// Config contains the configuration needed for most functions
type Config struct {
	// BiblePath path to the directory of Gratis Bibles
	BiblePath string
}

// New generates a new Config (but only if the path exists)
func New(path string) (*Config, error) {
	_, err := filepath.Glob(path)
	return &Config{
		BiblePath: path,
	}, err
}

// GetDocumentPath returns the path where documents are stored
/*
func GetDocumentPath() (string, error) {
	scope := gap.NewScope(gap.User, "heb12.com", "heb12")
	dataDir, err := scope.DataDir(docDir)
	if err != nil {
		return "", err
	}
	return dataDir + DocumentDir, nil
}*/

// ListAvailable lists the available translations stored according to language
func (c *Config) ListAvailable() (map[string][]string, error) {
	languages, err := filepath.Glob(c.BiblePath + "/*")
	if err != nil {
		return map[string][]string{}, err
	}
	var versions = map[string][]string{}
	for _, lang := range languages {
		// If the file is not a directory with a language code (which is two characters)
		/*if len(lang) != 2 {
			continue
		}*/
		lang = strings.Split(lang, "/")[len(strings.Split(lang, "/"))-1]
		files, err := filepath.Glob(c.BiblePath + "/" + lang + "/*")
		if err != nil {
			return map[string][]string{}, err
		}
		for _, file := range files {
			versionName := strings.Split(file, c.BiblePath+"/"+lang+"/")[1]
			versions[lang] = append(versions[lang], versionName)
		}
	}
	return versions, nil
}

// Delete removes a version from the DocumentDir
func (c *Config) Delete(ver string, lang string) error {
	err := os.RemoveAll(c.BiblePath + "/" + lang + "/" + ver)
	if err != nil {
		return err
	}
	return nil
}
