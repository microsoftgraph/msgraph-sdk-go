package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type ManagedAppPinCharacterSet int

const (
    NUMERIC_MANAGEDAPPPINCHARACTERSET ManagedAppPinCharacterSet = iota
    ALPHANUMERICANDSYMBOL_MANAGEDAPPPINCHARACTERSET
)

func (i ManagedAppPinCharacterSet) String() string {
    return []string{"NUMERIC", "ALPHANUMERICANDSYMBOL"}[i]
}
func ParseManagedAppPinCharacterSet(v string) (interface{}, error) {
    result := NUMERIC_MANAGEDAPPPINCHARACTERSET
    switch strings.ToUpper(v) {
        case "NUMERIC":
            result = NUMERIC_MANAGEDAPPPINCHARACTERSET
        case "ALPHANUMERICANDSYMBOL":
            result = ALPHANUMERICANDSYMBOL_MANAGEDAPPPINCHARACTERSET
        default:
            return 0, errors.New("Unknown ManagedAppPinCharacterSet value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppPinCharacterSet(values []ManagedAppPinCharacterSet) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
