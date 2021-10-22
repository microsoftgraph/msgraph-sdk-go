package graph
import (
    "strings"
    "errors"
)
type ManagedBrowserType int

const (
    NOTCONFIGURED_MANAGEDBROWSERTYPE ManagedBrowserType = iota
    MICROSOFTEDGE_MANAGEDBROWSERTYPE
)

func (i ManagedBrowserType) String() string {
    return []string{"NOTCONFIGURED", "MICROSOFTEDGE"}[i]
}
func ParseManagedBrowserType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTCONFIGURED":
            return NOTCONFIGURED_MANAGEDBROWSERTYPE, nil
        case "MICROSOFTEDGE":
            return MICROSOFTEDGE_MANAGEDBROWSERTYPE, nil
    }
    return 0, errors.New("Unknown ManagedBrowserType value: " + v)
}
func SerializeManagedBrowserType(values []ManagedBrowserType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
