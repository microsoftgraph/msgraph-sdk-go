package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type WellknownListName int

const (
    NONE_WELLKNOWNLISTNAME WellknownListName = iota
    DEFAULTLIST_WELLKNOWNLISTNAME
    FLAGGEDEMAILS_WELLKNOWNLISTNAME
    UNKNOWNFUTUREVALUE_WELLKNOWNLISTNAME
)

func (i WellknownListName) String() string {
    return []string{"NONE", "DEFAULTLIST", "FLAGGEDEMAILS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWellknownListName(v string) (interface{}, error) {
    result := NONE_WELLKNOWNLISTNAME
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_WELLKNOWNLISTNAME
        case "DEFAULTLIST":
            result = DEFAULTLIST_WELLKNOWNLISTNAME
        case "FLAGGEDEMAILS":
            result = FLAGGEDEMAILS_WELLKNOWNLISTNAME
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WELLKNOWNLISTNAME
        default:
            return 0, errors.New("Unknown WellknownListName value: " + v)
    }
    return &result, nil
}
func SerializeWellknownListName(values []WellknownListName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
