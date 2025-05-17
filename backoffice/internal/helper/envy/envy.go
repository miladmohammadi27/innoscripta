// package envy is the concrete implementation of envy functionality
// it is use to parse environment variables into a struct using StructTags
package envy

// Envy is the fundamental interface for all envy operations.
type Envy interface {
	// Parse would take a struct and parse the environment variables into it
	// it would return an error if it fails to parse the environment variables
	Parse(v interface{}) error
}
