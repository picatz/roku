# roku
> Roku External Control API package for Golang

## Rick Roll the Roku(s)

Rick Roll all roku devices on the network that have YouTube installed.
> `dQw4w9WgXcQ` is the contentID of "Never Gonna Give You Up" on YouTube

```golang
package main

import "github.com/picatz/roku"

func main() {
    devices, err := roku.Find(roku.DefaultWaitTime)

    if err != nil {
        panic(err)
    }

    for _, device := range devices {
        apps, err := device.Apps()
        if err != nil {
            panic(err)
        }
        for _, app := range apps {
            if app.Name == "YouTube" {
                device.LaunchApp(app.ID, map[string]string{
                    "contentID": "dQw4w9WgXcQ",
                })
            }
        }
    }
}
```

## Where's the remote(s)?

On supported devices, you can make the remote beep.

```golang
package main

import "github.com/picatz/roku"

func main() {
    devices, err := roku.Find(roku.DefaultWaitTime)

    if err != nil {
        panic(err)
    }

    for _, device := range devices {
        device.FindRemote()
     }
}
```
