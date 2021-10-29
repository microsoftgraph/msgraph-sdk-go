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
    switch strings.ToUpper(v) {
        case "EVERYONE":
            return EVERYONE_ONLINEMEETINGPRESENTERS, nil
        case "ORGANIZATION":
            return ORGANIZATION_ONLINEMEETINGPRESENTERS, nil
        case "ROLEISPRESENTER":
            return ROLEISPRESENTER_ONLINEMEETINGPRESENTERS, nil
        case "ORGANIZER":
            return ORGANIZER_ONLINEMEETINGPRESENTERS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ONLINEMEETINGPRESENTERS, nil
    }
    return 0, errors.New("Unknown OnlineMeetingPresenters value: " + v)
}
func SerializeOnlineMeetingPresenters(values []OnlineMeetingPresenters) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
