package osistool

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

// Osis contains all of the XML data in an OSIS work
type Osis struct {
	XMLName  xml.Name `xml:"osis"`
	OsisText osisText `xml:"osisText"`
}

type osisText struct {
	XMLName xml.Name `xml:"osisText"`
	Header  header   `xml:"header"`
	Div     []div    `xml:"div"`
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
func (osisData *Osis) Info() *Work {
	return &osisData.OsisText.Header.Work
}

// LoadOsis loads the OSIS data from a file specified by filepath
func LoadOsis(filename string) (*Osis, error) {
	byteValue, err := ioutil.ReadFile(filename)
	var osisData Osis
	err = xml.Unmarshal(byteValue, &osisData)
	if err != nil {
		return &Osis{}, err
	}
	return &osisData, err
}

// A BookInfo contains information about a book in an OSIS work
type BookInfo struct {
	ID               string
	Num              int
	Chapters         int
	VersesPerChapter []int
}

// BooksInfo gets information about the books available in a certain OSIS work
func (osisData *Osis) BooksInfo() ([]BookInfo, error) {
	var books []BookInfo
	i := 0
	for _, div := range osisData.OsisText.Div {
		if div.Type != "book" {
			continue
		}
		books = append(books, BookInfo{
			ID:               div.OsisID, // for books this should be just the ID (unlike chapters and verses that have a longer OsisID)
			Num:              i + 1,
			Chapters:         len(div.Chapters),
			VersesPerChapter: []int{},
		})
		for _, chapter := range div.Chapters {
			books[i].VersesPerChapter = append(books[i].VersesPerChapter, len(chapter.Verses))
		}
		i++
	}
	if books == nil {
		return books, nil
	}
	return books, nil
}

// GetBookInfo returns the BookInfo for a specific book
func (osisData *Osis) GetBookInfo(id string) (BookInfo, error) {
	booksInfo, err := osisData.BooksInfo()
	if err != nil {
		return BookInfo{}, err
	}
	for _, book := range booksInfo {
		if strings.ToLower(book.ID) == strings.ToLower(id) {
			return book, nil
		}
	}
	return BookInfo{}, errors.New("Book " + id + " not found in OSIS work " + osisData.Info().Identifier)
}

// Check verifies that a Reference actually exists in an OSIS work
func (osisData *Osis) Check(ref Reference) error {
	bookInfo, err := osisData.GetBookInfo(ref.ID)
	if err != nil {
		return err
	}

	if ref.Chapter-1 > bookInfo.Chapters || ref.Chapter-1 < 0 {
		return fmt.Errorf("Chapter %d in book %s out of range", ref.Chapter, ref.ID)
	}

	if ref.From > bookInfo.VersesPerChapter[ref.Chapter-1] || ref.From < 1 {
		return fmt.Errorf("Start verse %d in book %s chapter %d out of range", ref.From, ref.ID, ref.Chapter)
	}

	if ref.To > bookInfo.VersesPerChapter[ref.Chapter-1] || ref.To < 1 {
		return fmt.Errorf("End verse %d in book %s chapter %d out of range", ref.To, ref.ID, ref.Chapter)
	}

	return nil
}

// isSplit determines if the OSIS version is split by books or by whole versions (like the gratis.bible standard vs split)
func (osisData *Osis) isSplit() bool {
	if len(osisData.OsisText.Div) == 0 {
		return true
	}
	return false
}

// GetVerses puts verses together based on the reference
func (osisData *Osis) GetVerses(ref Reference) ([]string, error) {
	err := osisData.Check(ref)
	if err != nil {
		return []string{}, fmt.Errorf("Invalid reference: %v. Because: %v", ref, err)
	}
	var verses []string
	for i := ref.From - 1; i < ref.To; i++ {
		bookInfo, err := osisData.GetBookInfo(ref.ID)
		if err != nil {
			return []string{}, err
		}
		// If the file has one book, handle it differently (to allow for the split version of gratis.bible)
		var verse string
		if !osisData.isSplit() {
			verse = osisData.OsisText.Div[bookInfo.Num-1].Chapters[ref.Chapter-1].Verses[i]
		} else {
			verse = osisData.OsisText.Div[0].Chapters[ref.Chapter-1].Verses[i]
		}
		// Remove duplicate spaces
		verse = strings.Join(strings.Fields(verse), " ")
		verses = append(verses, verse)
	}
	return verses, nil
}
