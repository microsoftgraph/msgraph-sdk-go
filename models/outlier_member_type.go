package models
import (
    "errors"
)
// 
type OutlierMemberType int

const (
    USER_OUTLIERMEMBERTYPE OutlierMemberType = iota
    UNKNOWNFUTUREVALUE_OUTLIERMEMBERTYPE
)

func (i OutlierMemberType) String() string {
    return []string{"user", "unknownFutureValue"}[i]
}
func ParseOutlierMemberType(v string) (any, error) {
    result := USER_OUTLIERMEMBERTYPE
    switch v {
        case "user":
            result = USER_OUTLIERMEMBERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OUTLIERMEMBERTYPE
        default:
            return 0, errors.New("Unknown OutlierMemberType value: " + v)
    }
    return &result, nil
}
func SerializeOutlierMemberType(values []OutlierMemberType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
