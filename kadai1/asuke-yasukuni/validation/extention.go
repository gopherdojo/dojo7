// Validation package for commands.
package validation

// Returns true for allowed formats, false otherwise.
func Ext(ext string) bool {
	// Because there are few correspondence formats, we do not make map.
	if ext == "jpg" || ext == "png" {
		return true
	}
	return false
}
