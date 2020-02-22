package bver

import (
	"errors"
	"strings"
)

// Info returns information about a Bible version (and its language) when given an ID
func Info(ver string) (VersionInfo, string, error) {
	for _, lang := range AvailableLanguages {
		for _, version := range Versions[lang] {
			if strings.ToLower(ver) == strings.ToLower(version.ID) {
				return version, lang, nil
			}
		}
	}
	return VersionInfo{}, "", errors.New("Version " + ver + " not found.")
}

// Parse searches the names and alternative names of different Bible translations and returns the ID for it
func Parse(verName string) (string, error) {
	ver := strings.ToLower(verName)
	var result string
	for _, lang := range AvailableLanguages {
		for _, version := range Versions[lang] {
			name := strings.ToLower(version.Name)
			if ver == strings.ToLower(version.ID) {
				result = version.ID
			} else if ver == name {
				result = version.ID
			} else if ver == strings.Replace(name, " translation", "", 1) || ver == strings.Replace(name, " version", "", 1) || ver == strings.Replace(name, " bible", "", 1) {
				result = version.ID
			} else {
				for _, alias := range version.Aliases {
					if ver == strings.ToLower(alias) {
						result = version.ID
					}
				}
			}
			if result != "" {
				break
			}
		}
		if result != "" {
			break
		}
	}
	if result != "" {
		return result, nil
	}
	return result, errors.New("Version " + verName + " not found")
}
