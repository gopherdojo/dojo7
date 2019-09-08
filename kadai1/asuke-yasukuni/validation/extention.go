// Validation package for commands.
package validation

// Returns true for allowed formats, false otherwise.
func Ext(ext string) bool {
	if ext == "jpg" || ext == "png" {
		return true
	}
	return false
}
