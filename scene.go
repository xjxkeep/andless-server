package main

// sceneToObject maps scene identifiers to OSS object keys.
var sceneToObject = map[string]string{
	"linux-download":   "dev/latest/andless-console.deb",
	"windows-download": "dev/latest/AndlessConsole-Setup.exe",
	"macos-download":   "dev/latest/AndlessConsole.dmg",
}

// GetObjectKey returns the OSS object key for the given scene.
// Returns empty string if the scene is not found.
func GetObjectKey(scene string) string {
	return sceneToObject[scene]
}
