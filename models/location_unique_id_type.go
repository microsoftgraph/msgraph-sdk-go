package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
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
    result := UNKNOWN_LOCATIONUNIQUEIDTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_LOCATIONUNIQUEIDTYPE
        case "LOCATIONSTORE":
            result = LOCATIONSTORE_LOCATIONUNIQUEIDTYPE
        case "DIRECTORY":
            result = DIRECTORY_LOCATIONUNIQUEIDTYPE
        case "PRIVATE":
            result = PRIVATE_LOCATIONUNIQUEIDTYPE
        case "BING":
            result = BING_LOCATIONUNIQUEIDTYPE
        default:
            return 0, errors.New("Unknown LocationUniqueIdType value: " + v)
    }
    return &result, nil
}
func SerializeLocationUniqueIdType(values []LocationUniqueIdType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
