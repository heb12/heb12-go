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
	// Split determines whether or not to use the split version of Gratis Bibles
	Split bool
}

// New generates a new Config (but only if the path exists)
func New(config Config) (*Config, error) {
	_, err := filepath.Glob(config.BiblePath)
	return &Config{
		BiblePath: config.BiblePath,
		Split:     config.Split,
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

			// a Gratis Bible should be in the form of asv.xml
			// but there are also directories in the case of Gratis Split
			var split bool
			if len(strings.Split(versionName, ".")) == 2 {
				split = false
			} else if len(strings.Split(versionName, ".")) == 1 {
				split = true
			}

			if split != false {
				continue
			}

			versionName = strings.Split(versionName, ".xml")[0]

			versions[lang] = append(versions[lang], versionName)
		}
	}

	if len(versions) == 0 {
		return map[string][]string{}, errors.New("No versions available")
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

	if len(langs) == 0 {
		return []string{}, errors.New("No language dirs")
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
	if c.Split {
		return c.BiblePath + "/" + lang + "/" + ver + "/"
	}
	return c.BiblePath + "/" + lang + "/" + ver + ".xml"
}

// GetPathShort is identical to GetPath, except it only requires the version but not the language (requires that the language exists)
func (c *Config) GetPathShort(ver string) (string, error) {
	lang, err := c.GetLanguage(ver)
	if err != nil {
		return "", err
	}

	if c.Split {
		return c.BiblePath + "/" + lang + "/" + ver + "/", nil
	}
	return c.BiblePath + "/" + lang + "/" + ver + ".xml", nil
}

// ListSplitBooks returns the books available when using Gratis Split
func (c *Config) ListSplitBooks(ver string, lang string) ([]string, error) {
	if !c.Split {
		return []string{}, errors.New("ListSplit() is only for Gratis Split")
	}

	rawBooks, err := filepath.Glob(c.GetPath(ver, lang) + "/*")
	if err != nil {
		return []string{}, err
	}

	var books []string
	for _, book := range rawBooks {
		bookName := strings.Split(book, "/")[len(strings.Split(book, "/"))-1]
		books = append(books, strings.Split(bookName, ".xml")[0])
	}

	return books, nil
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
