package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the informationProtection singleton.
type ThreatAssessmentContentType int

const (
    MAIL_THREATASSESSMENTCONTENTTYPE ThreatAssessmentContentType = iota
    URL_THREATASSESSMENTCONTENTTYPE
    FILE_THREATASSESSMENTCONTENTTYPE
)

func (i ThreatAssessmentContentType) String() string {
    return []string{"MAIL", "URL", "FILE"}[i]
}
func ParseThreatAssessmentContentType(v string) (interface{}, error) {
    result := MAIL_THREATASSESSMENTCONTENTTYPE
    switch strings.ToUpper(v) {
        case "MAIL":
            result = MAIL_THREATASSESSMENTCONTENTTYPE
        case "URL":
            result = URL_THREATASSESSMENTCONTENTTYPE
        case "FILE":
            result = FILE_THREATASSESSMENTCONTENTTYPE
        default:
            return 0, errors.New("Unknown ThreatAssessmentContentType value: " + v)
    }
    return &result, nil
}
func SerializeThreatAssessmentContentType(values []ThreatAssessmentContentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
