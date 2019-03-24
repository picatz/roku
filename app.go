package roku

// App has all the information of an application.
// https://github.com/kinghrothgar/roku/blob/master/roku/roku.go#L63
type App struct {
	Name    string `xml:",chardata" json:"name,omitempty"`
	ID      int    `xml:"id,attr" json:"id,omitempty"`
	Type    string `xml:"type,attr" json:"type,omitempty"`
	SubType string `xml:"subtype,attr" json:"sub_type,omitempty"`
	Version string `xml:"version,attr" json:"version,omitempty"`
}
