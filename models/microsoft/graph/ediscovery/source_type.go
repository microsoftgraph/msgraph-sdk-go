package ediscovery
import (
    "strings"
    "errors"
)
// 
type SourceType int

const (
    MAILBOX_SOURCETYPE SourceType = iota
    SITE_SOURCETYPE
)

func (i SourceType) String() string {
    return []string{"MAILBOX", "SITE"}[i]
}
func ParseSourceType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "MAILBOX":
            return MAILBOX_SOURCETYPE, nil
        case "SITE":
            return SITE_SOURCETYPE, nil
    }
    return 0, errors.New("Unknown SourceType value: " + v)
}
func SerializeSourceType(values []SourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
