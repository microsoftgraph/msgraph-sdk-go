package graph
import (
    "strings"
    "errors"
)
type ThreatAssessmentStatus int

const (
    PENDING_THREATASSESSMENTSTATUS ThreatAssessmentStatus = iota
    COMPLETED_THREATASSESSMENTSTATUS
)

func (i ThreatAssessmentStatus) String() string {
    return []string{"PENDING", "COMPLETED"}[i]
}
func ParseThreatAssessmentStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "PENDING":
            return PENDING_THREATASSESSMENTSTATUS, nil
        case "COMPLETED":
            return COMPLETED_THREATASSESSMENTSTATUS, nil
    }
    return 0, errors.New("Unknown ThreatAssessmentStatus value: " + v)
}
func SerializeThreatAssessmentStatus(values []ThreatAssessmentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
