package config

import "testing"

// This is not very thorough, as it could change depending on the OS
func TestInitDirs(t *testing.T) {
	if err := InitDirs(); err != nil {
		t.Error(err)
	}
}

// testConfig is used for the next two tests
var testConfig = Config{
	GUI: ConfigGUI{
		Locale:      "en",
		Book:        "Heb",
		Chapter:     4,
		Translation: "web",
		Font: Font{
			Family: "Arial",
			Size:   18,
		},
	},
}

func TestGratisDir(t *testing.T) {
	dir, err := GratisDir()
	if err != nil {
		t.Error(err)
	}
	t.Log(dir)
}

func TestWriteConfig(t *testing.T) {
	err := WriteConfig(testConfig)
	if err != nil {
		t.Error(err)
	}
}

func TestReadConfig(t *testing.T) {
	_ = WriteConfig(testConfig)
	config, err := ReadConfig()
	if err != nil {
		t.Error(err)
	}
	if config != testConfig {
		t.Errorf("Expected config as %v but got %v instead.", testConfig, config)
	}
}
