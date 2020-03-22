// Package manage lists and manages available Bible translations
package manage

import (
	"reflect"
	"testing"
)

var defaultConfig = Config{
	BiblePath: "../bibles_test",
	Split:     false,
}
var splitConfig = Config{
	BiblePath: "../bibles_test",
	Split:     true,
}

func TestNew(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "Normal Gratis",
			args: args{
				config: defaultConfig,
			},
			want:    &defaultConfig,
			wantErr: false,
		},
		{
			name: "Split Gratis",
			args: args{
				config: splitConfig,
			},
			want:    &splitConfig,
			wantErr: false,
		},
		{
			name: "Without Split specified",
			args: args{
				config: Config{
					BiblePath: "../bibles_test",
				},
			},
			want: &Config{
				BiblePath: "../bibles_test",
				Split:     false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_ListAvailable(t *testing.T) {
	tests := []struct {
		name    string
		fields  Config
		want    map[string][]string
		wantErr bool
	}{
		{
			name:   "Normal Gratis",
			fields: defaultConfig,
			want: map[string][]string{
				"en": {
					"asv",
				},
			},
			wantErr: false,
		},
		{
			name:   "Split Gratis",
			fields: splitConfig,
			want: map[string][]string{
				"en": {
					"asv",
				},
			},
			wantErr: false,
		},
		{
			name: "No versions",
			fields: Config{
				BiblePath: "..",
			},
			want:    map[string][]string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BiblePath: tt.fields.BiblePath,
				Split:     tt.fields.Split,
			}
			got, err := c.ListAvailable()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.ListAvailable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.ListAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_ListLanguages(t *testing.T) {
	tests := []struct {
		name    string
		fields  Config
		want    []string
		wantErr bool
	}{
		{
			name:    "English",
			fields:  defaultConfig,
			want:    []string{"en"},
			wantErr: false,
		},
		{
			name: "Test none",
			fields: Config{
				BiblePath: "..",
			},
			want:    []string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BiblePath: tt.fields.BiblePath,
				Split:     tt.fields.Split,
			}
			got, err := c.ListLanguages()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.ListLanguages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.ListLanguages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetLanguage(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		fields  Config
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "English version",
			fields: defaultConfig,
			args: args{
				ver: "asv",
			},
			want:    "en",
			wantErr: false,
		},
		{
			name:   "Non-existent version",
			fields: defaultConfig,
			args: args{
				ver: "random",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BiblePath: tt.fields.BiblePath,
				Split:     tt.fields.Split,
			}
			got, err := c.GetLanguage(tt.args.ver)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.GetLanguage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Config.GetLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetPath(t *testing.T) {
	type args struct {
		ver  string
		lang string
	}
	tests := []struct {
		name   string
		fields Config
		args   args
		want   string
	}{
		{
			name:   "Normal",
			fields: defaultConfig,
			args: args{
				ver:  "asv",
				lang: "en",
			},
			want: defaultConfig.BiblePath + "/en/asv.xml",
		},
		{
			name:   "Split",
			fields: splitConfig,
			args: args{
				ver:  "asv",
				lang: "en",
			},
			want: defaultConfig.BiblePath + "/en/asv/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BiblePath: tt.fields.BiblePath,
				Split:     tt.fields.Split,
			}
			if got := c.GetPath(tt.args.ver, tt.args.lang); got != tt.want {
				t.Errorf("Config.GetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_GetPathShort(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		fields  Config
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "Normal",
			fields: defaultConfig,
			args: args{
				ver: "asv",
			},
			want:    defaultConfig.BiblePath + "/en/asv.xml",
			wantErr: false,
		},
		{
			name:   "Split",
			fields: splitConfig,
			args: args{
				ver: "asv",
			},
			want:    defaultConfig.BiblePath + "/en/asv/",
			wantErr: false,
		},
		{
			name:   "Non-existing normal",
			fields: defaultConfig,
			args: args{
				ver: "random",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "Non-existing split",
			fields: splitConfig,
			args: args{
				ver: "random",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BiblePath: tt.fields.BiblePath,
				Split:     tt.fields.Split,
			}
			got, err := c.GetPathShort(tt.args.ver)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.GetPathShort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Config.GetPathShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_ListSplitBooks(t *testing.T) {
	type args struct {
		ver  string
		lang string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Split books list",
			args: args{
				ver:  "asv",
				lang: "en",
			},
			want:    []string{"2john", "gen", "heb", "john"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &splitConfig
			got, err := c.ListSplitBooks(tt.args.ver, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.ListSplitBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.ListSplitBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_IsAvailable(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name   string
		fields Config
		args   args
		want   bool
	}{
		{
			name:   "Normal true",
			fields: defaultConfig,
			args: args{
				ver: "asv",
			},
			want: true,
		},
		{
			name:   "Normal false",
			fields: defaultConfig,
			args: args{
				ver: "random",
			},
			want: false,
		},
		{
			name:   "Split true",
			fields: splitConfig,
			args: args{
				ver: "asv",
			},
			want: true,
		},
		{
			name:   "Split false",
			fields: splitConfig,
			args: args{
				ver: "random",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BiblePath: tt.fields.BiblePath,
				Split:     tt.fields.Split,
			}
			if got := c.IsAvailable(tt.args.ver); got != tt.want {
				t.Errorf("Config.IsAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

// I am still figuring out the best way to handle this one, since I don't want to delete the test files
func TestConfig_Delete(t *testing.T) {
	type fields struct {
		BiblePath string
		Split     bool
	}
	type args struct {
		ver  string
		lang string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BiblePath: tt.fields.BiblePath,
				Split:     tt.fields.Split,
			}
			if err := c.Delete(tt.args.ver, tt.args.lang); (err != nil) != tt.wantErr {
				t.Errorf("Config.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
