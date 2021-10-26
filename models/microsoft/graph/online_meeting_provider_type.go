package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ONLINEMEETINGPROVIDERTYPE, nil
        case "SKYPEFORBUSINESS":
            return SKYPEFORBUSINESS_ONLINEMEETINGPROVIDERTYPE, nil
        case "SKYPEFORCONSUMER":
            return SKYPEFORCONSUMER_ONLINEMEETINGPROVIDERTYPE, nil
        case "TEAMSFORBUSINESS":
            return TEAMSFORBUSINESS_ONLINEMEETINGPROVIDERTYPE, nil
    }
    return 0, errors.New("Unknown OnlineMeetingProviderType value: " + v)
}
func SerializeOnlineMeetingProviderType(values []OnlineMeetingProviderType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
