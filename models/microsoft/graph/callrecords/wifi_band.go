package callrecords
import (
    "strings"
    "errors"
)
// 
type WifiBand int

const (
    UNKNOWN_WIFIBAND WifiBand = iota
    FREQUENCY24GHZ_WIFIBAND
    FREQUENCY50GHZ_WIFIBAND
    FREQUENCY60GHZ_WIFIBAND
    UNKNOWNFUTUREVALUE_WIFIBAND
)

func (i WifiBand) String() string {
    return []string{"UNKNOWN", "FREQUENCY24GHZ", "FREQUENCY50GHZ", "FREQUENCY60GHZ", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWifiBand(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_WIFIBAND, nil
        case "FREQUENCY24GHZ":
            return FREQUENCY24GHZ_WIFIBAND, nil
        case "FREQUENCY50GHZ":
            return FREQUENCY50GHZ_WIFIBAND, nil
        case "FREQUENCY60GHZ":
            return FREQUENCY60GHZ_WIFIBAND, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WIFIBAND, nil
    }
    return 0, errors.New("Unknown WifiBand value: " + v)
}
func SerializeWifiBand(values []WifiBand) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
