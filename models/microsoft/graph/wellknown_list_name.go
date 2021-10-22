package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_WELLKNOWNLISTNAME, nil
        case "DEFAULTLIST":
            return DEFAULTLIST_WELLKNOWNLISTNAME, nil
        case "FLAGGEDEMAILS":
            return FLAGGEDEMAILS_WELLKNOWNLISTNAME, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WELLKNOWNLISTNAME, nil
    }
    return 0, errors.New("Unknown WellknownListName value: " + v)
}
func SerializeWellknownListName(values []WellknownListName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
