package main

// sceneToObject maps scene identifiers to OSS object keys.
var sceneToObject = map[string]string{
	"linux-download":       "release/latest/andless-console.deb",
	"linux-download-dev":   "dev/latest/andless-console.deb",
	"windows-download":     "release/latest/AndlessConsole-Setup.exe",
	"windows-download-dev": "dev/latest/AndlessConsole-Setup.exe",
	"macos-download":       "release/latest/AndlessConsole.dmg",
	"macos-download-dev":   "dev/latest/AndlessConsole.dmg",
	"android-download":     "release/latest/app-release.apk",
}

// GetObjectKey returns the OSS object key for the given scene.
// Returns empty string if the scene is not found.
func GetObjectKey(scene string) string {
	return sceneToObject[scene]
}
