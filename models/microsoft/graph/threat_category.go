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
    switch strings.ToUpper(v) {
        case "UNDEFINED":
            return UNDEFINED_THREATCATEGORY, nil
        case "SPAM":
            return SPAM_THREATCATEGORY, nil
        case "PHISHING":
            return PHISHING_THREATCATEGORY, nil
        case "MALWARE":
            return MALWARE_THREATCATEGORY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_THREATCATEGORY, nil
    }
    return 0, errors.New("Unknown ThreatCategory value: " + v)
}
func SerializeThreatCategory(values []ThreatCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
