# roku
> Roku External Control API package for Golang

```golang
package main

import "github.com/picatz/roku"

func main() {
    devices, err := roku.Find(roku.DefaultWaitTime)

    if err != nil {
        panic(err)
    }

    // rick-roll all devices on the network that have YouTube installed
    for _, device := range devices {
        apps, err := device.Apps()
        if err != nil {
            panic(err)
        }
        for _, app := range apps {
            if app.Name == "YouTube" {
                device.LaunchApp(app.ID, map[string]string{
                    // contentID of "Never Gonna Give You Up" on YouTube
                    "contentID": "dQw4w9WgXcQ",
                })
            }
        }
    }
}
```
