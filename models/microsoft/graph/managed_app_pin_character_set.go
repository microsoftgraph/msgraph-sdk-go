package graph
import (
    "strings"
    "errors"
)
type ManagedAppPinCharacterSet int

const (
    NUMERIC_MANAGEDAPPPINCHARACTERSET ManagedAppPinCharacterSet = iota
    ALPHANUMERICANDSYMBOL_MANAGEDAPPPINCHARACTERSET
)

func (i ManagedAppPinCharacterSet) String() string {
    return []string{"NUMERIC", "ALPHANUMERICANDSYMBOL"}[i]
}
func ParseManagedAppPinCharacterSet(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NUMERIC":
            return NUMERIC_MANAGEDAPPPINCHARACTERSET, nil
        case "ALPHANUMERICANDSYMBOL":
            return ALPHANUMERICANDSYMBOL_MANAGEDAPPPINCHARACTERSET, nil
    }
    return 0, errors.New("Unknown ManagedAppPinCharacterSet value: " + v)
}
func SerializeManagedAppPinCharacterSet(values []ManagedAppPinCharacterSet) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
