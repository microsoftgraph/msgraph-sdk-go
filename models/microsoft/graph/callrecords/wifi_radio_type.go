package callrecords
import (
    "strings"
    "errors"
)
// 
type WifiRadioType int

const (
    UNKNOWN_WIFIRADIOTYPE WifiRadioType = iota
    WIFI80211A_WIFIRADIOTYPE
    WIFI80211B_WIFIRADIOTYPE
    WIFI80211G_WIFIRADIOTYPE
    WIFI80211N_WIFIRADIOTYPE
    WIFI80211AC_WIFIRADIOTYPE
    WIFI80211AX_WIFIRADIOTYPE
    UNKNOWNFUTUREVALUE_WIFIRADIOTYPE
)

func (i WifiRadioType) String() string {
    return []string{"UNKNOWN", "WIFI80211A", "WIFI80211B", "WIFI80211G", "WIFI80211N", "WIFI80211AC", "WIFI80211AX", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWifiRadioType(v string) (interface{}, error) {
    result := UNKNOWN_WIFIRADIOTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_WIFIRADIOTYPE
        case "WIFI80211A":
            result = WIFI80211A_WIFIRADIOTYPE
        case "WIFI80211B":
            result = WIFI80211B_WIFIRADIOTYPE
        case "WIFI80211G":
            result = WIFI80211G_WIFIRADIOTYPE
        case "WIFI80211N":
            result = WIFI80211N_WIFIRADIOTYPE
        case "WIFI80211AC":
            result = WIFI80211AC_WIFIRADIOTYPE
        case "WIFI80211AX":
            result = WIFI80211AX_WIFIRADIOTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WIFIRADIOTYPE
        default:
            return 0, errors.New("Unknown WifiRadioType value: " + v)
    }
    return &result, nil
}
func SerializeWifiRadioType(values []WifiRadioType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
