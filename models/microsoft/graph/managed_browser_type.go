package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type ManagedBrowserType int

const (
    NOTCONFIGURED_MANAGEDBROWSERTYPE ManagedBrowserType = iota
    MICROSOFTEDGE_MANAGEDBROWSERTYPE
)

func (i ManagedBrowserType) String() string {
    return []string{"NOTCONFIGURED", "MICROSOFTEDGE"}[i]
}
func ParseManagedBrowserType(v string) (interface{}, error) {
    result := NOTCONFIGURED_MANAGEDBROWSERTYPE
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_MANAGEDBROWSERTYPE
        case "MICROSOFTEDGE":
            result = MICROSOFTEDGE_MANAGEDBROWSERTYPE
        default:
            return 0, errors.New("Unknown ManagedBrowserType value: " + v)
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
