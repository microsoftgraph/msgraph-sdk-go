package graph
import (
    "strings"
    "errors"
)
// 
type ThreatCategory int

const (
    UNDEFINED_THREATCATEGORY ThreatCategory = iota
    SPAM_THREATCATEGORY
    PHISHING_THREATCATEGORY
    MALWARE_THREATCATEGORY
    UNKNOWNFUTUREVALUE_THREATCATEGORY
)

func (i ThreatCategory) String() string {
    return []string{"UNDEFINED", "SPAM", "PHISHING", "MALWARE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseThreatCategory(v string) (interface{}, error) {
    result := UNDEFINED_THREATCATEGORY
    switch strings.ToUpper(v) {
        case "UNDEFINED":
            result = UNDEFINED_THREATCATEGORY
        case "SPAM":
            result = SPAM_THREATCATEGORY
        case "PHISHING":
            result = PHISHING_THREATCATEGORY
        case "MALWARE":
            result = MALWARE_THREATCATEGORY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_THREATCATEGORY
        default:
            return 0, errors.New("Unknown ThreatCategory value: " + v)
    }
    return &result, nil
}
func SerializeThreatCategory(values []ThreatCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
