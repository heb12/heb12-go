// Package manage lists and manages available Bible translations
package manage

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

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

// ListAvailable lists the available translations stored according to language
func (c *Config) ListAvailable() (map[string][]string, error) {
	languages, err := c.ListLanguages()
	if err != nil {
		return map[string][]string{}, err
	}

	versions := make(map[string][]string)
	for _, lang := range languages {
		files, err := filepath.Glob(c.BiblePath + "/" + lang + "/*")
		if err != nil {
			return map[string][]string{}, err
		}
		for _, file := range files {
			versionName := strings.Split(file, c.BiblePath+"/"+lang+"/")[1]
			versionName = strings.Split(versionName, ".xml")[0]
			versions[lang] = append(versions[lang], versionName)
		}
	}

	return versions, nil
}

// ListLanguages lists all the languages available in Config.BiblePath
func (c *Config) ListLanguages() ([]string, error) {
	languages, err := filepath.Glob(c.BiblePath + "/*")
	if err != nil {
		return []string{}, err
	}

	var langs []string
	for _, lang := range languages {
		// Split the file name to use just language code
		lang = strings.Split(lang, "/")[len(strings.Split(lang, "/"))-1]
		// If the file is not a directory with a language code (which is two characters)
		if len(lang) != 2 {
			continue
		}
		langs = append(langs, lang)

	}
	return langs, nil
}

// GetLanguage returns the language for a version ID
func (c *Config) GetLanguage(ver string) (string, error) {
	versions, err := c.ListAvailable()
	if err != nil {
		return "", nil
	}

	languages, err := c.ListLanguages()
	if err != nil {
		return "", nil
	}

	for _, lang := range languages {
		for _, version := range versions[lang] {
			if strings.ToLower(ver) == version {
				return lang, nil
			}
		}
	}

	return "", errors.New("Version " + ver + " not found")
}

// GetPath returns the full path of an OSIS Bible (when using normal GratisBible)
func (c *Config) GetPath(ver string, lang string) string {
	return c.BiblePath + "/" + lang + "/" + ver + ".xml"
}

// IsAvailable uses ListAvailable to determine is a certain language is available
func (c *Config) IsAvailable(ver string) bool {
	languages, err := c.ListAvailable()
	if err != nil {
		return false
	}
	for _, versions := range languages {
		for _, version := range versions {
			if version == ver {
				return true
			}
		}
	}
	return false
}

// Delete removes a version from the DocumentDir
func (c *Config) Delete(ver string, lang string) error {
	err := os.RemoveAll(c.GetPath(ver, lang))
	if err != nil {
		return err
	}
	return nil
}
