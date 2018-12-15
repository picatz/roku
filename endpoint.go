package roku

import "net/http"
import "net/url"
import "encoding/xml"
import "errors"
import "io/ioutil"

// Endpoint represents a roku device on the network.
type Endpoint struct {
	url string
}

var (
	// ErrNoAppFound is the error returned when no applicaiton is found.
	ErrNoAppFound = errors.New("no app found")
	// ErrNoAppsFound is the error returned when there are no applications found.
	ErrNoAppsFound = errors.New("no apps found")
	// ErrNoRespBody is the error returned when there was no response body from
	// the roku device.
	ErrNoRespBody = errors.New("no response body")

	pathToQueryActiveApp     = "/query/active-app"
	pathToQueryAvailableApps = "/query/apps"
	pathToQueryIcon          = "/query/icon/"
	pathToQueryDeviceInfo    = "/query/device-info"
	pathToKeypress           = "/keypress/"
	pathToKeyUp              = "/keyup/"
	pathToKeyDown            = "/keyup/"
	pathToLaunch             = "/launch/"
	pathToInstall            = "/install/"
	pathToSearch             = "/search/browse"
	pathToInput              = "/input"
)

// Apps returns the available applications for an Endpoint.
func (e *Endpoint) Apps() (*Apps, error) {
	apps := Apps{}

	resp, err := http.Get(e.url + pathToQueryAvailableApps)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Body == nil {
		return nil, ErrNoRespBody
	}

	decoder := xml.NewDecoder(resp.Body)

	err = decoder.Decode(&apps)

	if err != nil {
		return nil, err
	}

	if len(apps.All) == 0 {
		return nil, ErrNoAppsFound
	}

	return &apps, nil
}

// ActiveApp returns the currently active App for an Endpoint.
func (e *Endpoint) ActiveApp() (*App, error) {
	apps := Apps{}

	resp, err := http.Get(e.url + pathToQueryActiveApp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.Body == nil {
		return nil, ErrNoRespBody
	}

	decoder := xml.NewDecoder(resp.Body)

	err = decoder.Decode(&apps)

	if err != nil {
		return nil, err
	}

	if len(apps.All) == 0 {
		return nil, ErrNoAppFound
	}

	return &apps.All[0], nil
}

// Keypress allows simulating hitting a button on the roku remote.
func (e *Endpoint) Keypress(key string) error {
	resp, err := http.Post(e.url+pathToKeypress+key, "", nil)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

// KeyUp allows simulating releasing a button on the roku remote.
func (e *Endpoint) KeyUp(key string) error {
	resp, err := http.Post(e.url+pathToKeyUp+key, "", nil)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

// KeyDown allows simulating pressing down on a button on the roku remote.
func (e *Endpoint) KeyDown(key string) error {
	resp, err := http.Post(e.url+pathToKeyDown+key, "", nil)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

// Icon returns the image (PNG) for the given applicaton ID.
func (e *Endpoint) Icon(id string) ([]byte, error) {
	resp, err := http.Get(e.url + pathToQueryIcon + id)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// LaunchApp starts an application on the roku device.
func (e *Endpoint) LaunchApp(id string, params map[string]string) error {
	u, err := url.Parse(e.url + pathToLaunch + id)
	if err != nil {
		return err
	}

	query := u.Query()

	for key, value := range params {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()

	resp, err := http.Post(u.String(), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

// InstallApp attempts to install an application on the roku device.
func (e *Endpoint) InstallApp(id string, params map[string]string) error {
	u, err := url.Parse(e.url + pathToInstall + id)
	if err != nil {
		return err
	}

	query := u.Query()

	for key, value := range params {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()

	resp, err := http.Post(u.String(), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

// Search provides a search functionality for a roku device.
func (e *Endpoint) Search(params map[string]string) error {
	u, err := url.Parse(e.url + pathToSearch)
	if err != nil {
		return err
	}

	query := u.Query()

	for key, value := range params {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()

	resp, err := http.Post(u.String(), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

// Input provides input functionality for a roku device.
func (e *Endpoint) Input(params map[string]string) error {
	u, err := url.Parse(e.url + pathToInput)
	if err != nil {
		return err
	}

	query := u.Query()

	for key, value := range params {
		query.Set(key, value)
	}

	u.RawQuery = query.Encode()

	resp, err := http.Post(u.String(), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return nil
}

// DeviceInfo returns the DeviceInfo for a roku device endpoint.
func (e *Endpoint) DeviceInfo() (*DeviceInfo, error) {
	resp, err := http.Get(e.url + pathToQueryDeviceInfo)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	if resp.Body == nil {
		return nil, ErrNoRespBody
	}

	deviceInfo := &DeviceInfo{}

	decoder := xml.NewDecoder(resp.Body)

	err = decoder.Decode(deviceInfo)

	if err != nil {
		return nil, err
	}

	return deviceInfo, nil
}