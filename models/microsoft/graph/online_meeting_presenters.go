package graph
import (
    "strings"
    "errors"
)
// 
type OnlineMeetingPresenters int

const (
    EVERYONE_ONLINEMEETINGPRESENTERS OnlineMeetingPresenters = iota
    ORGANIZATION_ONLINEMEETINGPRESENTERS
    ROLEISPRESENTER_ONLINEMEETINGPRESENTERS
    ORGANIZER_ONLINEMEETINGPRESENTERS
    UNKNOWNFUTUREVALUE_ONLINEMEETINGPRESENTERS
)

func (i OnlineMeetingPresenters) String() string {
    return []string{"EVERYONE", "ORGANIZATION", "ROLEISPRESENTER", "ORGANIZER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseOnlineMeetingPresenters(v string) (interface{}, error) {
    result := EVERYONE_ONLINEMEETINGPRESENTERS
    switch strings.ToUpper(v) {
        case "EVERYONE":
            result = EVERYONE_ONLINEMEETINGPRESENTERS
        case "ORGANIZATION":
            result = ORGANIZATION_ONLINEMEETINGPRESENTERS
        case "ROLEISPRESENTER":
            result = ROLEISPRESENTER_ONLINEMEETINGPRESENTERS
        case "ORGANIZER":
            result = ORGANIZER_ONLINEMEETINGPRESENTERS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ONLINEMEETINGPRESENTERS
        default:
            return 0, errors.New("Unknown OnlineMeetingPresenters value: " + v)
    }
    return &result, nil
}
func SerializeOnlineMeetingPresenters(values []OnlineMeetingPresenters) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
