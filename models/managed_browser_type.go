package models
import (
    "errors"
    "strings"
)
// Type of managed browser
type ManagedBrowserType int

const (
    // Not configured
    NOTCONFIGURED_MANAGEDBROWSERTYPE ManagedBrowserType = iota
    // Microsoft Edge
    MICROSOFTEDGE_MANAGEDBROWSERTYPE
)

func (i ManagedBrowserType) String() string {
    var values []string
    for p := ManagedBrowserType(1); p <= MICROSOFTEDGE_MANAGEDBROWSERTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"notConfigured", "microsoftEdge"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseManagedBrowserType(v string) (any, error) {
    var result ManagedBrowserType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "notConfigured":
                result |= NOTCONFIGURED_MANAGEDBROWSERTYPE
            case "microsoftEdge":
                result |= MICROSOFTEDGE_MANAGEDBROWSERTYPE
            default:
                return 0, errors.New("Unknown ManagedBrowserType value: " + v)
        }
    }
    return &result, nil
}
func SerializeManagedBrowserType(values []ManagedBrowserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ManagedBrowserType) isMultiValue() bool {
    return true
}
