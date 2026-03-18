package main

// sceneToObject maps scene identifiers to OSS object keys.
var sceneToObject = map[string]string{
	"app-download":     "releases/app-latest.apk",
	"desktop-download": "releases/desktop-latest.dmg",
	"sdk-download":     "releases/sdk-latest.zip",
}

// GetObjectKey returns the OSS object key for the given scene.
// Returns empty string if the scene is not found.
func GetObjectKey(scene string) string {
	return sceneToObject[scene]
}
