// Package bible uses the modules bref, heb12/manage, and heb12/osis to get Bible verses
package bible

import (
	"code.heb12.com/heb12/bref"
	"code.heb12.com/heb12/heb12/config"
	"code.heb12.com/heb12/heb12/manage"
	"code.heb12.com/heb12/heb12/osis"

	"errors"
)

func getManager() (*manage.Config, error) {
	gratisDir, err := config.GratisDir()
	if err != nil {
		return &manage.Config{}, err
	}
	manager, err := manage.New(gratisDir)
	if err != nil {
		return &manage.Config{}, err
	}
	return manager, err
}

// Get returns Bible text from a string reference with bref, manage, and osis
func Get(reference string, version string) ([]string, error) {
	// Prepare all of the different data from the input

	manager, err := getManager()
	if err != nil {
		return []string{}, err
	}

	ref, err := bref.Parse(reference)
	if err != nil {
		return []string{}, err
	}

	err = bref.Check(ref)
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

	osisData, err := osis.Load(manager.GetPath(version, language))
	if err != nil {
		return []string{}, err
	}

	text, err := osisData.GetVerses(osis.Reference{
		ID:      ref.ID,
		Chapter: ref.Chapter,
		From:    ref.From,
		To:      ref.To,
	})
	if err != nil {
		return []string{}, err
	}

	return text, nil
}

func List() (map[string][]string, error) {
	manager, err := getManager()
	if err != nil {
		return map[string][]string{}, err
	}

	versions, err := manager.ListAvailable()
	return versions, err
}

func ListLanguages() ([]string, error) {
	manager, err := getManager()
	if err != nil {
		return []string{}, err
	}

	langs, err := manager.ListLanguages()
	return langs, err
}
