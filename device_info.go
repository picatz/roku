package roku

import "encoding/xml"

// DeviceInfo contains all the information about the roku device.
type DeviceInfo struct {
	XMLName                  xml.Name `xml:"device-info" json:"-"`
	Text                     string   `xml:",chardata" json:"text,omitempty"`
	Udn                      string   `xml:"udn" json:"udn,omitempty"`
	SerialNumber             string   `xml:"serial-number" json:"serial_number,omitempty"`
	DeviceID                 string   `xml:"device-id" json:"device_id,omitempty"`
	AdvertisingID            string   `xml:"advertising-id" json:"advertising_id,omitempty"`
	VendorName               string   `xml:"vendor-name" json:"vendor_name,omitempty"`
	ModelName                string   `xml:"model-name" json:"model_name,omitempty"`
	ModelNumber              string   `xml:"model-number" json:"model_number,omitempty"`
	ModelRegion              string   `xml:"model-region" json:"model_region,omitempty"`
	IsTv                     string   `xml:"is-tv" json:"is_tv,omitempty"`
	IsStick                  string   `xml:"is-stick" json:"is_stick,omitempty"`
	SupportsEthernet         string   `xml:"supports-ethernet" json:"supports_ethernet,omitempty"`
	WifiMac                  string   `xml:"wifi-mac" json:"wifi_mac,omitempty"`
	WifiDriver               string   `xml:"wifi-driver" json:"wifi_driver,omitempty"`
	EthernetMac              string   `xml:"ethernet-mac" json:"ethernet_mac,omitempty"`
	NetworkType              string   `xml:"network-type" json:"network_type,omitempty"`
	NetworkName              string   `xml:"network-name" json:"network_name,omitempty"`
	FriendlyDeviceName       string   `xml:"friendly-device-name" json:"friendly_device_name,omitempty"`
	FriendlyModelName        string   `xml:"friendly-model-name" json:"friendly_model_name,omitempty"`
	DefaultDeviceName        string   `xml:"default-device-name" json:"default_device_name,omitempty"`
	UserDeviceName           string   `xml:"user-device-name" json:"user_device_name,omitempty"`
	BuildNumber              string   `xml:"build-number" json:"build_number,omitempty"`
	SoftwareVersion          string   `xml:"software-version" json:"software_version,omitempty"`
	SoftwareBuild            string   `xml:"software-build" json:"software_build,omitempty"`
	SecureDevice             string   `xml:"secure-device" json:"secure_device,omitempty"`
	Language                 string   `xml:"language" json:"language,omitempty"`
	Country                  string   `xml:"country" json:"country,omitempty"`
	Locale                   string   `xml:"locale" json:"locale,omitempty"`
	TimeZoneAuto             string   `xml:"time-zone-auto" json:"time_zone_auto,omitempty"`
	TimeZone                 string   `xml:"time-zone" json:"time_zone,omitempty"`
	TimeZoneName             string   `xml:"time-zone-name" json:"time_zone_name,omitempty"`
	TimeZoneTz               string   `xml:"time-zone-tz" json:"time_zone_tz,omitempty"`
	TimeZoneOffset           string   `xml:"time-zone-offset" json:"time_zone_offset,omitempty"`
	ClockFormat              string   `xml:"clock-format" json:"clock_format,omitempty"`
	Uptime                   string   `xml:"uptime" json:"uptime,omitempty"`
	PowerMode                string   `xml:"power-mode" json:"power_mode,omitempty"`
	SupportsSuspend          string   `xml:"supports-suspend" json:"supports_suspend,omitempty"`
	SupportsFindRemote       string   `xml:"supports-find-remote" json:"supports_find_remote,omitempty"`
	FindRemoteIsPossible     string   `xml:"find-remote-is-possible" json:"find_remote_is_possible,omitempty"`
	SupportsAudioGuide       string   `xml:"supports-audio-guide" json:"supports_audio_guide,omitempty"`
	SupportsRva              string   `xml:"supports-rva" json:"supports_rva,omitempty"`
	DeveloperEnabled         string   `xml:"developer-enabled" json:"developer_enabled,omitempty"`
	KeyedDeveloperID         string   `xml:"keyed-developer-id" json:"keyed_developer_id,omitempty"`
	SearchEnabled            string   `xml:"search-enabled" json:"search_enabled,omitempty"`
	SearchChannelsEnabled    string   `xml:"search-channels-enabled" json:"search_channels_enabled,omitempty"`
	VoiceSearchEnabled       string   `xml:"voice-search-enabled" json:"voice_search_enabled,omitempty"`
	NotificationsEnabled     string   `xml:"notifications-enabled" json:"notifications_enabled,omitempty"`
	NotificationsFirstUse    string   `xml:"notifications-first-use" json:"notifications_first_use,omitempty"`
	SupportsPrivateListening string   `xml:"supports-private-listening" json:"supports_private_listening,omitempty"`
	HeadphonesConnected      string   `xml:"headphones-connected" json:"headphones_connected,omitempty"`
	SupportsEcsTextedit      string   `xml:"supports-ecs-textedit" json:"supports_ecs_textedit,omitempty"`
	SupportsEcsMicrophone    string   `xml:"supports-ecs-microphone" json:"supports_ecs_microphone,omitempty"`
	SupportsWakeOnWlan       string   `xml:"supports-wake-on-wlan" json:"supports_wake_on_wlan,omitempty"`
	HasPlayOnRoku            string   `xml:"has-play-on-roku" json:"has_play_on_roku,omitempty"`
	HasMobileScreensaver     string   `xml:"has-mobile-screensaver" json:"has_mobile_screensaver,omitempty"`
	SupportURL               string   `xml:"support-url" json:"support_url,omitempty"`
	GrandcentralVersion      string   `xml:"grandcentral-version" json:"grandcentral_version,omitempty"`
	TrcVersion               string   `xml:"trc-version" json:"trc_version,omitempty"`
	TrcChannelVersion        string   `xml:"trc-channel-version" json:"trc_channel_version,omitempty"`
	DavinciVersion           string   `xml:"davinci-version" json:"davinci_version,omitempty"`
}
