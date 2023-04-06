package security
import (
    "errors"
)
// 
type QueryType int

const (
    FILES_QUERYTYPE QueryType = iota
    MESSAGES_QUERYTYPE
    UNKNOWNFUTUREVALUE_QUERYTYPE
)

func (i QueryType) String() string {
    return []string{"files", "messages", "unknownFutureValue"}[i]
}
func ParseQueryType(v string) (any, error) {
    result := FILES_QUERYTYPE
    switch v {
        case "files":
            result = FILES_QUERYTYPE
        case "messages":
            result = MESSAGES_QUERYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_QUERYTYPE
        default:
            return 0, errors.New("Unknown QueryType value: " + v)
    }
    return &result, nil
}
func SerializeQueryType(values []QueryType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
