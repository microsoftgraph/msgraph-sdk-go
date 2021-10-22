package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNDEFINED":
            return UNDEFINED_THREATASSESSMENTREQUESTSOURCE, nil
        case "USER":
            return USER_THREATASSESSMENTREQUESTSOURCE, nil
        case "ADMINISTRATOR":
            return ADMINISTRATOR_THREATASSESSMENTREQUESTSOURCE, nil
    }
    return 0, errors.New("Unknown ThreatAssessmentRequestSource value: " + v)
}
func SerializeThreatAssessmentRequestSource(values []ThreatAssessmentRequestSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
