package callrecords
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_WIFIRADIOTYPE, nil
        case "WIFI80211A":
            return WIFI80211A_WIFIRADIOTYPE, nil
        case "WIFI80211B":
            return WIFI80211B_WIFIRADIOTYPE, nil
        case "WIFI80211G":
            return WIFI80211G_WIFIRADIOTYPE, nil
        case "WIFI80211N":
            return WIFI80211N_WIFIRADIOTYPE, nil
        case "WIFI80211AC":
            return WIFI80211AC_WIFIRADIOTYPE, nil
        case "WIFI80211AX":
            return WIFI80211AX_WIFIRADIOTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WIFIRADIOTYPE, nil
    }
    return 0, errors.New("Unknown WifiRadioType value: " + v)
}
func SerializeWifiRadioType(values []WifiRadioType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
