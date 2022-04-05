package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type EndpointType int

const (
    DEFAULT_ESCAPED_ENDPOINTTYPE EndpointType = iota
    VOICEMAIL_ENDPOINTTYPE
    SKYPEFORBUSINESS_ENDPOINTTYPE
    SKYPEFORBUSINESSVOIPPHONE_ENDPOINTTYPE
    UNKNOWNFUTUREVALUE_ENDPOINTTYPE
)

func (i EndpointType) String() string {
    return []string{"DEFAULT_ESCAPED", "VOICEMAIL", "SKYPEFORBUSINESS", "SKYPEFORBUSINESSVOIPPHONE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEndpointType(v string) (interface{}, error) {
    result := DEFAULT_ESCAPED_ENDPOINTTYPE
    switch strings.ToUpper(v) {
        case "DEFAULT_ESCAPED":
            result = DEFAULT_ESCAPED_ENDPOINTTYPE
        case "VOICEMAIL":
            result = VOICEMAIL_ENDPOINTTYPE
        case "SKYPEFORBUSINESS":
            result = SKYPEFORBUSINESS_ENDPOINTTYPE
        case "SKYPEFORBUSINESSVOIPPHONE":
            result = SKYPEFORBUSINESSVOIPPHONE_ENDPOINTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ENDPOINTTYPE
        default:
            return 0, errors.New("Unknown EndpointType value: " + v)
    }
    return &result, nil
}
func SerializeEndpointType(values []EndpointType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
