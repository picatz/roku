package roku

type apps struct {
	All []*App `xml:"app"`
}

// Apps is a collection of one or more App(s)
type Apps = []*App

// Found using https://channelstore.roku.com/browse
var (
	// YouTubeAppID is the channel app ID number for Google's YouTube
	YouTubeAppID = 837
	// NetflixAppID is the channel app ID number for Netflix
	NetflixAppID = 12
	// PrimeVideoAppID is the channel app ID number for Amazon Prime Video
	PrimeVideoAppID = 13
	// FireFoxAppID is the channel app ID number for FireFox
	FireFoxAppID = 47545
	// HuluAppID is the channel app ID number for Hulu
	HuluAppID = 2285
	// PandoraAppID is the channel app ID number for Pandora
	PandoraAppID = 28
	// SpotifyAppID is the channel app ID number for Spotify
	SpotifyAppID = 19977
	// ESPNAppID is the channel app ID number for ESPN
	ESPNAppID = 34376
	// PlayOnRokuID is the channel app ID number for PlayOnRoku
	PlayOnRokuID = 15985
)
