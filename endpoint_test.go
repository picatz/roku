package roku

import "testing"

func TestActiveApp(t *testing.T) {
	devices, err := Find(DefaultWaitTime)

	if err != nil {
		t.Error(err)
	}

	if len(devices) == 0 {
		t.Fatal("no roku devices found")
	}

	endpoint := devices[0]

	app, err := endpoint.ActiveApp()

	if err != nil {
		t.Error(err)
	}

	t.Error(app)
}

func TestApps(t *testing.T) {
	devices, err := Find(DefaultWaitTime)

	if err != nil {
		t.Error(err)
	}

	if len(devices) == 0 {
		t.Fatal("no roku devices found")
	}

	endpoint := devices[0]

	apps, err := endpoint.Apps()

	if err != nil {
		t.Error(err)
	}

	t.Error(apps)
}

func TestKeypress(t *testing.T) {
	devices, err := Find(DefaultWaitTime)

	if err != nil {
		t.Error(err)
	}

	if len(devices) == 0 {
		t.Fatal("no roku devices found")
	}

	endpoint := devices[0]

	//err = endpoint.Keypress("home")
	//err = endpoint.Keypress("FindRemote")
	err = endpoint.Keypress("wiggle")

	if err != nil {
		t.Error(err)
	}
}

func TestLaunchApp(t *testing.T) {
	devices, err := Find(DefaultWaitTime)

	if err != nil {
		t.Error(err)
	}

	if len(devices) == 0 {
		t.Fatal("no roku devices found")
	}

	endpoint := devices[0]

	apps, err := endpoint.Apps()

	if err != nil {
		t.Fatal(err)
	}

	err = endpoint.LaunchApp(apps[0].ID, nil)

	// Rick Roll:
	// err = endpoint.LaunchApp("837", map[string]string{"contentID": "dQw4w9WgXcQ"})

	if err != nil {
		t.Error(err)
	}
}

func TestInstallApp(t *testing.T) {
	devices, err := Find(DefaultWaitTime)

	if err != nil {
		t.Error(err)
	}

	if len(devices) == 0 {
		t.Fatal("no roku devices found")
	}

	endpoint := devices[0]

	apps, err := endpoint.Apps()

	if err != nil {
		t.Fatal(err)
	}

	err = endpoint.InstallApp(apps[0].ID, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestSearch(t *testing.T) {
	devices, err := Find(DefaultWaitTime)

	if err != nil {
		t.Error(err)
	}

	if len(devices) == 0 {
		t.Fatal("no roku devices found")
	}

	endpoint := devices[0]

	err = endpoint.Search(Params{"keyword": "film"})

	if err != nil {
		t.Error(err)
	}
}

func TestIcon(t *testing.T) {
	devices, err := Find(DefaultWaitTime)

	if err != nil {
		t.Error(err)
	}

	if len(devices) == 0 {
		t.Fatal("no roku devices found")
	}

	endpoint := devices[0]

	apps, err := endpoint.Apps()

	if err != nil {
		t.Fatal(err)
	}

	bytes, err := endpoint.Icon(apps[0].ID)

	if err != nil {
		t.Error(err)
	}

	if len(bytes) <= 0 {
		t.Error("no icon found")
	}
}

func TestInput(t *testing.T) {
	devices, err := Find(DefaultWaitTime)

	if err != nil {
		t.Error(err)
	}

	if len(devices) == 0 {
		t.Fatal("no roku devices found")
	}

	endpoint := devices[0]

	if err != nil {
		t.Fatal(err)
	}

	err = endpoint.Input(Params{
		"acceleration.x": "0.0",
		"acceleration.y": "0.0",
		"acceleration.z": "9.8",
	})

	if err != nil {
		t.Error(err)
	}
}

func TestDeviceInfo(t *testing.T) {
	devices, err := Find(DefaultWaitTime)

	if err != nil {
		t.Error(err)
	}

	if len(devices) == 0 {
		t.Fatal("no roku devices found")
	}

	endpoint := devices[0]

	if err != nil {
		t.Fatal(err)
	}

	info, err := endpoint.DeviceInfo()

	if err != nil {
		t.Fatal(err)
	}

	t.Log(info.Uptime)
}
