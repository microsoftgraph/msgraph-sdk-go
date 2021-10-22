package callrecords
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_NETWORKCONNECTIONTYPE, nil
        case "WIRED":
            return WIRED_NETWORKCONNECTIONTYPE, nil
        case "WIFI":
            return WIFI_NETWORKCONNECTIONTYPE, nil
        case "MOBILE":
            return MOBILE_NETWORKCONNECTIONTYPE, nil
        case "TUNNEL":
            return TUNNEL_NETWORKCONNECTIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_NETWORKCONNECTIONTYPE, nil
    }
    return 0, errors.New("Unknown NetworkConnectionType value: " + v)
}
func SerializeNetworkConnectionType(values []NetworkConnectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
