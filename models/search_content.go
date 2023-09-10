package models
import (
    "errors"
    "strings"
)
// 
type SearchContent int

const (
    SHAREDCONTENT_SEARCHCONTENT SearchContent = iota
    PRIVATECONTENT_SEARCHCONTENT
    UNKNOWNFUTUREVALUE_SEARCHCONTENT
)

func (i SearchContent) String() string {
    var values []string
    for p := SearchContent(1); p <= UNKNOWNFUTUREVALUE_SEARCHCONTENT; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"sharedContent", "privateContent", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseSearchContent(v string) (any, error) {
    var result SearchContent
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "sharedContent":
                result |= SHAREDCONTENT_SEARCHCONTENT
            case "privateContent":
                result |= PRIVATECONTENT_SEARCHCONTENT
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_SEARCHCONTENT
            default:
                return 0, errors.New("Unknown SearchContent value: " + v)
        }
    }
    return &result, nil
}
func SerializeSearchContent(values []SearchContent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SearchContent) isMultiValue() bool {
    return true
}
