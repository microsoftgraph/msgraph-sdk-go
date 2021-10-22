package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "MAIL":
            return MAIL_THREATASSESSMENTCONTENTTYPE, nil
        case "URL":
            return URL_THREATASSESSMENTCONTENTTYPE, nil
        case "FILE":
            return FILE_THREATASSESSMENTCONTENTTYPE, nil
    }
    return 0, errors.New("Unknown ThreatAssessmentContentType value: " + v)
}
func SerializeThreatAssessmentContentType(values []ThreatAssessmentContentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
