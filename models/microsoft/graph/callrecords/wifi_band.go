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
    result := UNKNOWN_WIFIBAND
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_WIFIBAND
        case "FREQUENCY24GHZ":
            result = FREQUENCY24GHZ_WIFIBAND
        case "FREQUENCY50GHZ":
            result = FREQUENCY50GHZ_WIFIBAND
        case "FREQUENCY60GHZ":
            result = FREQUENCY60GHZ_WIFIBAND
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WIFIBAND
        default:
            return 0, errors.New("Unknown WifiBand value: " + v)
    }
    return &result, nil
}
func SerializeWifiBand(values []WifiBand) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
