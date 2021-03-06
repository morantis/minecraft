// minecraft project minecraft.go
package minecraft

import ()

const (
	// Proper Minecraft username regex
	ValidUsernameRegex = `[a-zA-Z0-9_]{1,16}`

	// Proper Minecraft UUID regex
	ValidUuidRegex = `[0-9a-f]{32}`

	// Minecraft username-or-UUID regex
	ValidUsernameOrUuidRegex = "(" + ValidUuidRegex + "|" + ValidUsernameRegex + ")"
)
