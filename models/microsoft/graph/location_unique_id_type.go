package graph
import (
    "strings"
    "errors"
)
// 
type LocationUniqueIdType int

const (
    UNKNOWN_LOCATIONUNIQUEIDTYPE LocationUniqueIdType = iota
    LOCATIONSTORE_LOCATIONUNIQUEIDTYPE
    DIRECTORY_LOCATIONUNIQUEIDTYPE
    PRIVATE_LOCATIONUNIQUEIDTYPE
    BING_LOCATIONUNIQUEIDTYPE
)

func (i LocationUniqueIdType) String() string {
    return []string{"UNKNOWN", "LOCATIONSTORE", "DIRECTORY", "PRIVATE", "BING"}[i]
}
func ParseLocationUniqueIdType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_LOCATIONUNIQUEIDTYPE, nil
        case "LOCATIONSTORE":
            return LOCATIONSTORE_LOCATIONUNIQUEIDTYPE, nil
        case "DIRECTORY":
            return DIRECTORY_LOCATIONUNIQUEIDTYPE, nil
        case "PRIVATE":
            return PRIVATE_LOCATIONUNIQUEIDTYPE, nil
        case "BING":
            return BING_LOCATIONUNIQUEIDTYPE, nil
    }
    return 0, errors.New("Unknown LocationUniqueIdType value: " + v)
}
func SerializeLocationUniqueIdType(values []LocationUniqueIdType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
