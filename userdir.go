package userdir

import "os"
import "path"

func Home() string {
	home := os.Getenv("HOME")

	if home == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		home = os.Getenv("HOMEDRIVE")
	}
	if home == "" {
		home = os.Getenv("HOMEPATH")
	}

	return home
}

func AuthorizedKeys() string {
	return path.Join(Home(), ".ssh", "authorized_keys")
}
