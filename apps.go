package roku

const TVInput = "tvinput."

type apps struct {
	All []*App `xml:"app"`
}

// Apps is a collection of one or more App(s)
type Apps = []*App

// Found from querying App IDs once installed on a TV
// https://channelstore.roku.com/details/{AppID} will take you to the channel's page
// Note that this will redirect you to a URL of the form https://channelstore.roku.com/details/{WebID}/{AppName}
// where WebID also identifies the app but only via Roku's website (and not on your TV)
var (
	YouTubeAppID    = "837"
	NetflixAppID    = "12"
	PrimeVideoAppID = "13"
	HuluAppID       = "2285"
	PandoraAppID    = "28"
	SpotifyAppID    = "19977"
	ESPNAppID       = "34376"
	PlayOnRokuID    = "15985"
	HDMI1           = TVInput + "hdmi1"
	HDMI2           = TVInput + "hdmi2"
	HDMI3           = TVInput + "hdmi3"
	AV              = TVInput + "cvbs"
	LiveTV          = TVInput + "dtv"
	AppleTV         = "551012"
	Peacock         = "593099"
	HBOMax          = "61322"
	CBS             = "619667"
	DisneyPlus      = "291097"
	ParamountPlus   = "31440"
	Spotify         = "22297"
	CW              = "111255"
	ABC             = "73376"
	FoxSports       = "95307"
)
