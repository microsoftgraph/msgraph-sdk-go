package externalconnectors
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "TEXT":
            return TEXT_EXTERNALITEMCONTENTTYPE, nil
        case "HTML":
            return HTML_EXTERNALITEMCONTENTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EXTERNALITEMCONTENTTYPE, nil
    }
    return 0, errors.New("Unknown ExternalItemContentType value: " + v)
}
func SerializeExternalItemContentType(values []ExternalItemContentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
