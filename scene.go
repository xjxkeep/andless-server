package main

// sceneToObject maps scene identifiers to OSS object keys.
var sceneToObject = map[string]string{
	"linux-download":                    "release/latest/dashboard-linux-x64.deb",
	"linux-download-dev":                "dev/latest/dashboard-linux-x64.deb",
	"windows-download":                  "release/latest/dashboard-windows-x64.exe",
	"windows-download-dev":              "dev/latest/dashboard-windows-x64.exe",
	"macos-apple-silicon-download":      "release/latest/dashboard-macos-apple-silicon.dmg",
	"macos-apple-silicon-download-dev":  "dev/latest/dashboard-macos-apple-silicon.dmg",
	"macos-intel-download":              "release/latest/dashboard-macos-intel.dmg",
	"macos-intel-download-dev":          "dev/latest/dashboard-macos-intel.dmg",
	"android-download":                  "release/latest/app-release.apk",
}

// GetObjectKey returns the OSS object key for the given scene.
// Returns empty string if the scene is not found.
func GetObjectKey(scene string) string {
	return sceneToObject[scene]
}
