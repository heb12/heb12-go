// Package bible uses the modules bref, heb12/manage, and heb12/osis to get Bible verses from Split Gratis Bible
package bible

import (
	"reflect"
	"testing"

	"code.heb12.com/heb12/heb12/manage"
)

func TestNew(t *testing.T) {
	type args struct {
		gratisDir string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "Normal test",
			args: args{
				gratisDir: "../bibles_test",
			},
			want: &Config{
				Manager: &manage.Config{
					BiblePath: "../bibles_test",
					Split:     true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.gratisDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got.Manager, tt.want.Manager)
			}
		})
	}
}

var defaultConfig = Config{
	Manager: &manage.Config{
		BiblePath: "../bibles_test",
		Split:     true,
	},
}

func TestConfig_Get(t *testing.T) {
	type args struct {
		reference string
		version   string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "One verse",
			args: args{
				reference: "Hebrews 4 12",
				version:   "asv",
			},
			want:    []string{"For the word of God is living, and active, and sharper than any two-edged sword, and piercing even to the dividing of soul and spirit, of both joints and marrow, and quick to discern the thoughts and intents of the heart."},
			wantErr: false,
		},
		{
			name: "Single chapter book",
			args: args{
				reference: "2 John 1 2",
				version:   "asv",
			},
			want:    []string{"for the truth`s sake which abideth in us, and it shall be with us for ever:"},
			wantErr: false,
		},
		{
			name: "Multiple verses",
			args: args{
				reference: "John 3:16-17",
				version:   "asv",
			},
			want:    []string{"For God so loved the world, that he gave his only begotten Son, that whosoever believeth on him should not perish, but have eternal life.", "For God sent not the Son into the world to judge the world; but that the world should be saved through him."},
			wantErr: false,
		},
		{
			name: "Invalid reference",
			args: args{
				reference: "Random",
				version:   "asv",
			},
			want:    []string{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &defaultConfig
			got, err := c.Get(tt.args.reference, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// The following don't really need test cases since the tests are available in bible/manage and this is just an interface for them

func TestConfig_List(t *testing.T) {
	type fields struct {
		Manager *manage.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Manager: tt.fields.Manager,
			}
			got, err := c.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_ListLanguages(t *testing.T) {
	type fields struct {
		Manager *manage.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Manager: tt.fields.Manager,
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
