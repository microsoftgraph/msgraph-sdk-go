package graph
import (
    "strings"
    "errors"
)
// 
type ThreatAssessmentRequestSource int

const (
    UNDEFINED_THREATASSESSMENTREQUESTSOURCE ThreatAssessmentRequestSource = iota
    USER_THREATASSESSMENTREQUESTSOURCE
    ADMINISTRATOR_THREATASSESSMENTREQUESTSOURCE
)

func (i ThreatAssessmentRequestSource) String() string {
    return []string{"UNDEFINED", "USER", "ADMINISTRATOR"}[i]
}
func ParseThreatAssessmentRequestSource(v string) (interface{}, error) {
    result := UNDEFINED_THREATASSESSMENTREQUESTSOURCE
    switch strings.ToUpper(v) {
        case "UNDEFINED":
            result = UNDEFINED_THREATASSESSMENTREQUESTSOURCE
        case "USER":
            result = USER_THREATASSESSMENTREQUESTSOURCE
        case "ADMINISTRATOR":
            result = ADMINISTRATOR_THREATASSESSMENTREQUESTSOURCE
        default:
            return 0, errors.New("Unknown ThreatAssessmentRequestSource value: " + v)
    }
    return &result, nil
}
func SerializeThreatAssessmentRequestSource(values []ThreatAssessmentRequestSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
