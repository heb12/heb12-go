package osistool

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// Osis contains all of the XML data in an OSIS work
type Osis struct {
	XMLName  xml.Name `xml:"osis"`
	OsisText osisText `xml:"osisText"`
}

type osisText struct {
	XMLName xml.Name `xml:"osisText"`
	Header  header   `xml:"header"`
	Div     div      `xml:"div"`
}

type header struct {
	XMLName      xml.Name     `xml:"header"`
	RevisionDesc revisionDesc `xml:"revisionDesc"`
	Work         Work         `xml:"work"`
}

type revisionDesc struct {
	XMLName xml.Name `xml:"revisionDesc"`
	Date    string   `xml:"date"`
	P       string   `xml:"p"`
}

// Work contains information about a specific work
type Work struct {
	XMLName     xml.Name `xml:"work"`
	Title       string   `xml:"title"`
	Contributor string   `xml:"contributor"`
	Subject     string   `xml:"subject"`
	Date        string   `xml:"date"`
	Description string   `xml:"description"`
	Publisher   string   `xml:"publisher"`
	Type        string   `xml:"type"`
	Identifier  string   `xml:"identifier"`
	Source      string   `xml:"source"`
	Language    string   `xml:"language"`
	Coverage    string   `xml:"coverage"`
	Rights      string   `xml:"rights"`
}

type div struct {
	XMLName  xml.Name  `xml:"div"`
	Type     string    `xml:"type,attr"`
	OsisID   string    `xml:"osisID,attr"`
	Chapters []chapter `xml:"chapter"`
}

type chapter struct {
	XMLName xml.Name `xml:"chapter"`
	OsisID  string   `xml:"osisID,attr"`
	Verses  []string `xml:"verse"`
}

// Reference contains basic information for a Bible reference
type Reference struct {
	ID      string
	Chapter int
	From    int
	To      int
}

// Info returns the Work information about a Bible version
func (osisData *Osis) Info() (Work, error) {
	return osisData.OsisText.Header.Work, nil
}

// LoadOsis loads the OSIS data from a file specified by filepath
func LoadOsis(filename string) (*Osis, error) {
	osisFile, err := os.Open(filename)
	if err != nil {
		return &Osis{}, err
	}
	defer osisFile.Close()
	byteValue, err := ioutil.ReadFile(filename)
	var osisData Osis
	err = xml.Unmarshal(byteValue, &osisData)
	if err != nil {
		return &Osis{}, err
	}
	return &osisData, err
}

// GetVerses processes XML and puts verses together based on the reference
func (osisData *Osis) GetVerses(ref Reference) []string {
	var verses []string
	for i := ref.From - 1; i < ref.To; i++ {
		verse := osisData.OsisText.Div.Chapters[ref.Chapter-1].Verses[i]
		verses = append(verses, verse)
	}
	return verses
}
