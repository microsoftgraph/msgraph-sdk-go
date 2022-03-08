package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of group entities.
type OnlineMeetingProviderType int

const (
    UNKNOWN_ONLINEMEETINGPROVIDERTYPE OnlineMeetingProviderType = iota
    SKYPEFORBUSINESS_ONLINEMEETINGPROVIDERTYPE
    SKYPEFORCONSUMER_ONLINEMEETINGPROVIDERTYPE
    TEAMSFORBUSINESS_ONLINEMEETINGPROVIDERTYPE
)

func (i OnlineMeetingProviderType) String() string {
    return []string{"UNKNOWN", "SKYPEFORBUSINESS", "SKYPEFORCONSUMER", "TEAMSFORBUSINESS"}[i]
}
func ParseOnlineMeetingProviderType(v string) (interface{}, error) {
    result := UNKNOWN_ONLINEMEETINGPROVIDERTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ONLINEMEETINGPROVIDERTYPE
        case "SKYPEFORBUSINESS":
            result = SKYPEFORBUSINESS_ONLINEMEETINGPROVIDERTYPE
        case "SKYPEFORCONSUMER":
            result = SKYPEFORCONSUMER_ONLINEMEETINGPROVIDERTYPE
        case "TEAMSFORBUSINESS":
            result = TEAMSFORBUSINESS_ONLINEMEETINGPROVIDERTYPE
        default:
            return 0, errors.New("Unknown OnlineMeetingProviderType value: " + v)
    }
    return &result, nil
}
func SerializeOnlineMeetingProviderType(values []OnlineMeetingProviderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
