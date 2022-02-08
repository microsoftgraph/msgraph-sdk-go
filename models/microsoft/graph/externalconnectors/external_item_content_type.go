package externalconnectors
import (
    "strings"
    "errors"
)
// 
type ExternalItemContentType int

const (
    TEXT_EXTERNALITEMCONTENTTYPE ExternalItemContentType = iota
    HTML_EXTERNALITEMCONTENTTYPE
    UNKNOWNFUTUREVALUE_EXTERNALITEMCONTENTTYPE
)

func (i ExternalItemContentType) String() string {
    return []string{"TEXT", "HTML", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseExternalItemContentType(v string) (interface{}, error) {
    result := TEXT_EXTERNALITEMCONTENTTYPE
    switch strings.ToUpper(v) {
        case "TEXT":
            result = TEXT_EXTERNALITEMCONTENTTYPE
        case "HTML":
            result = HTML_EXTERNALITEMCONTENTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EXTERNALITEMCONTENTTYPE
        default:
            return 0, errors.New("Unknown ExternalItemContentType value: " + v)
    }
    return &result, nil
}
func SerializeExternalItemContentType(values []ExternalItemContentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
