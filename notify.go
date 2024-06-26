package main

import "github.com/martinlindhe/notify"

func not() {
	// show a notification
	notify.Notify("ASKPASS", "SUDO", "Someone is trying to open askpass", "")

	// show a notification and play a alert sound
	// notify.Alert("app name", "alert", "some text", "path/to/icon.png")
}
