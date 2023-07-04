package models
import (
    "errors"
)
// 
type OutlierContainerType int

const (
    GROUP_OUTLIERCONTAINERTYPE OutlierContainerType = iota
    UNKNOWNFUTUREVALUE_OUTLIERCONTAINERTYPE
)

func (i OutlierContainerType) String() string {
    return []string{"group", "unknownFutureValue"}[i]
}
func ParseOutlierContainerType(v string) (any, error) {
    result := GROUP_OUTLIERCONTAINERTYPE
    switch v {
        case "group":
            result = GROUP_OUTLIERCONTAINERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OUTLIERCONTAINERTYPE
        default:
            return 0, errors.New("Unknown OutlierContainerType value: " + v)
    }
    return &result, nil
}
func SerializeOutlierContainerType(values []OutlierContainerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
