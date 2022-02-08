package callrecords
import (
    "strings"
    "errors"
)
// 
type NetworkConnectionType int

const (
    UNKNOWN_NETWORKCONNECTIONTYPE NetworkConnectionType = iota
    WIRED_NETWORKCONNECTIONTYPE
    WIFI_NETWORKCONNECTIONTYPE
    MOBILE_NETWORKCONNECTIONTYPE
    TUNNEL_NETWORKCONNECTIONTYPE
    UNKNOWNFUTUREVALUE_NETWORKCONNECTIONTYPE
)

func (i NetworkConnectionType) String() string {
    return []string{"UNKNOWN", "WIRED", "WIFI", "MOBILE", "TUNNEL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseNetworkConnectionType(v string) (interface{}, error) {
    result := UNKNOWN_NETWORKCONNECTIONTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_NETWORKCONNECTIONTYPE
        case "WIRED":
            result = WIRED_NETWORKCONNECTIONTYPE
        case "WIFI":
            result = WIFI_NETWORKCONNECTIONTYPE
        case "MOBILE":
            result = MOBILE_NETWORKCONNECTIONTYPE
        case "TUNNEL":
            result = TUNNEL_NETWORKCONNECTIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_NETWORKCONNECTIONTYPE
        default:
            return 0, errors.New("Unknown NetworkConnectionType value: " + v)
    }
    return &result, nil
}
func SerializeNetworkConnectionType(values []NetworkConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
