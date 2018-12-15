package roku

import "encoding/xml"

// DeviceInfo contains all the information about the roku device.
type DeviceInfo struct {
	XMLName                  xml.Name `xml:"device-info"`
	Text                     string   `xml:",chardata"`
	Udn                      string   `xml:"udn"`
	SerialNumber             string   `xml:"serial-number"`
	DeviceID                 string   `xml:"device-id"`
	AdvertisingID            string   `xml:"advertising-id"`
	VendorName               string   `xml:"vendor-name"`
	ModelName                string   `xml:"model-name"`
	ModelNumber              string   `xml:"model-number"`
	ModelRegion              string   `xml:"model-region"`
	IsTv                     string   `xml:"is-tv"`
	IsStick                  string   `xml:"is-stick"`
	SupportsEthernet         string   `xml:"supports-ethernet"`
	WifiMac                  string   `xml:"wifi-mac"`
	WifiDriver               string   `xml:"wifi-driver"`
	EthernetMac              string   `xml:"ethernet-mac"`
	NetworkType              string   `xml:"network-type"`
	NetworkName              string   `xml:"network-name"`
	FriendlyDeviceName       string   `xml:"friendly-device-name"`
	FriendlyModelName        string   `xml:"friendly-model-name"`
	DefaultDeviceName        string   `xml:"default-device-name"`
	UserDeviceName           string   `xml:"user-device-name"`
	BuildNumber              string   `xml:"build-number"`
	SoftwareVersion          string   `xml:"software-version"`
	SoftwareBuild            string   `xml:"software-build"`
	SecureDevice             string   `xml:"secure-device"`
	Language                 string   `xml:"language"`
	Country                  string   `xml:"country"`
	Locale                   string   `xml:"locale"`
	TimeZoneAuto             string   `xml:"time-zone-auto"`
	TimeZone                 string   `xml:"time-zone"`
	TimeZoneName             string   `xml:"time-zone-name"`
	TimeZoneTz               string   `xml:"time-zone-tz"`
	TimeZoneOffset           string   `xml:"time-zone-offset"`
	ClockFormat              string   `xml:"clock-format"`
	Uptime                   string   `xml:"uptime"`
	PowerMode                string   `xml:"power-mode"`
	SupportsSuspend          string   `xml:"supports-suspend"`
	SupportsFindRemote       string   `xml:"supports-find-remote"`
	FindRemoteIsPossible     string   `xml:"find-remote-is-possible"`
	SupportsAudioGuide       string   `xml:"supports-audio-guide"`
	SupportsRva              string   `xml:"supports-rva"`
	DeveloperEnabled         string   `xml:"developer-enabled"`
	KeyedDeveloperID         string   `xml:"keyed-developer-id"`
	SearchEnabled            string   `xml:"search-enabled"`
	SearchChannelsEnabled    string   `xml:"search-channels-enabled"`
	VoiceSearchEnabled       string   `xml:"voice-search-enabled"`
	NotificationsEnabled     string   `xml:"notifications-enabled"`
	NotificationsFirstUse    string   `xml:"notifications-first-use"`
	SupportsPrivateListening string   `xml:"supports-private-listening"`
	HeadphonesConnected      string   `xml:"headphones-connected"`
	SupportsEcsTextedit      string   `xml:"supports-ecs-textedit"`
	SupportsEcsMicrophone    string   `xml:"supports-ecs-microphone"`
	SupportsWakeOnWlan       string   `xml:"supports-wake-on-wlan"`
	HasPlayOnRoku            string   `xml:"has-play-on-roku"`
	HasMobileScreensaver     string   `xml:"has-mobile-screensaver"`
	SupportURL               string   `xml:"support-url"`
	GrandcentralVersion      string   `xml:"grandcentral-version"`
	TrcVersion               string   `xml:"trc-version"`
	TrcChannelVersion        string   `xml:"trc-channel-version"`
	DavinciVersion           string   `xml:"davinci-version"`
}
