package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the getRecentNotebooks method.
type OnenoteSourceService int

const (
    UNKNOWN_ONENOTESOURCESERVICE OnenoteSourceService = iota
    ONEDRIVE_ONENOTESOURCESERVICE
    ONEDRIVEFORBUSINESS_ONENOTESOURCESERVICE
    ONPREMONEDRIVEFORBUSINESS_ONENOTESOURCESERVICE
)

func (i OnenoteSourceService) String() string {
    return []string{"UNKNOWN", "ONEDRIVE", "ONEDRIVEFORBUSINESS", "ONPREMONEDRIVEFORBUSINESS"}[i]
}
func ParseOnenoteSourceService(v string) (interface{}, error) {
    result := UNKNOWN_ONENOTESOURCESERVICE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ONENOTESOURCESERVICE
        case "ONEDRIVE":
            result = ONEDRIVE_ONENOTESOURCESERVICE
        case "ONEDRIVEFORBUSINESS":
            result = ONEDRIVEFORBUSINESS_ONENOTESOURCESERVICE
        case "ONPREMONEDRIVEFORBUSINESS":
            result = ONPREMONEDRIVEFORBUSINESS_ONENOTESOURCESERVICE
        default:
            return 0, errors.New("Unknown OnenoteSourceService value: " + v)
    }
    return &result, nil
}
func SerializeOnenoteSourceService(values []OnenoteSourceService) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
