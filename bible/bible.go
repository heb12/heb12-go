// Package bible uses the modules bref, heb12/manage, and heb12/osis to get Bible verses from Split Gratis Bible
package bible

import (
	"code.heb12.com/heb12/bref"
	"code.heb12.com/heb12/heb12/config"
	"code.heb12.com/heb12/heb12/manage"
	"code.heb12.com/heb12/heb12/osis"

	"errors"
	"strings"
)

// Config provides a manager for the rest of the program to depend upon
type Config struct {
	Manager *manage.Config
}

// New returns a basic configuration, a blank string uses heb12/config's value
func New(dir string) (*Config, error) {
	gratisDir, err := config.GratisDir()
	if err != nil {
		return &Config{}, err
	}

	if dir != "" {
		gratisDir = dir
	}

	manager, err := manage.New(
		manage.Config{
			BiblePath: gratisDir,
			Split:     true,
		},
	)
	return &Config{
		manager,
	}, err
}

// Get returns Bible text from a string reference with bref, manage, and osis
func (c *Config) Get(reference string, version string) ([]string, error) {
	// Prepare all of the different data from the input

	manager := c.Manager

	ref, err := bref.Process(reference)
	if err != nil {
		return []string{}, err
	}

	if !manager.IsAvailable(version) {
		return []string{}, errors.New("Version " + version + " not available")
	}

	language, err := manager.GetLanguage(version)
	if err != nil {
		return []string{}, err
	}

	// Get the actual Bible text with heb12/osis

	osisData, err := osis.Load(manager.GetPath(version, language) + "/" + strings.ToLower(ref.Book.ID) + ".xml")
	if err != nil {
		return []string{}, err
	}

	return osisData.GetVerses(osis.Reference{
		ID:      ref.Book.ID,
		Chapter: ref.Chapter,
		From:    ref.From,
		To:      ref.To,
	})
}

// List returns a list of all available versions in all available languages via the manager
func (c *Config) List() (map[string][]string, error) {
	return c.Manager.ListAvailable()
}

// ListLanguages returns a list of all available versions for a specific language via the manager
func (c *Config) ListLanguages() ([]string, error) {
	return c.Manager.ListLanguages()
}
