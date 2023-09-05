package security
import (
    "errors"
    "strings"
)
// 
type SourceType int

const (
    MAILBOX_SOURCETYPE SourceType = iota
    SITE_SOURCETYPE
    UNKNOWNFUTUREVALUE_SOURCETYPE
)

func (i SourceType) String() string {
    var values []string
    for p := SourceType(1); p <= UNKNOWNFUTUREVALUE_SOURCETYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"mailbox", "site", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseSourceType(v string) (any, error) {
    var result SourceType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "mailbox":
                result |= MAILBOX_SOURCETYPE
            case "site":
                result |= SITE_SOURCETYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_SOURCETYPE
            default:
                return 0, errors.New("Unknown SourceType value: " + v)
        }
    }
    return &result, nil
}
func SerializeSourceType(values []SourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SourceType) isMultiValue() bool {
    return true
}
