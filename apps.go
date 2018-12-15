package roku

type apps struct {
	All []*App `xml:"app"`
}

// Apps is a collection of one or more App(s)
type Apps = []*App
