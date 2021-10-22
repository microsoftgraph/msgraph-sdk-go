package graph
import (
    "strings"
    "errors"
)
type EndpointType int

const (
    DEFAULT_ENDPOINTTYPE EndpointType = iota
    VOICEMAIL_ENDPOINTTYPE
    SKYPEFORBUSINESS_ENDPOINTTYPE
    SKYPEFORBUSINESSVOIPPHONE_ENDPOINTTYPE
    UNKNOWNFUTUREVALUE_ENDPOINTTYPE
)

func (i EndpointType) String() string {
    return []string{"DEFAULT", "VOICEMAIL", "SKYPEFORBUSINESS", "SKYPEFORBUSINESSVOIPPHONE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEndpointType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DEFAULT":
            return DEFAULT_ENDPOINTTYPE, nil
        case "VOICEMAIL":
            return VOICEMAIL_ENDPOINTTYPE, nil
        case "SKYPEFORBUSINESS":
            return SKYPEFORBUSINESS_ENDPOINTTYPE, nil
        case "SKYPEFORBUSINESSVOIPPHONE":
            return SKYPEFORBUSINESSVOIPPHONE_ENDPOINTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ENDPOINTTYPE, nil
    }
    return 0, errors.New("Unknown EndpointType value: " + v)
}
func SerializeEndpointType(values []EndpointType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
