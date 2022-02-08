package graph
import (
    "strings"
    "errors"
)
// 
type ThreatAssessmentStatus int

const (
    PENDING_THREATASSESSMENTSTATUS ThreatAssessmentStatus = iota
    COMPLETED_THREATASSESSMENTSTATUS
)

func (i ThreatAssessmentStatus) String() string {
    return []string{"PENDING", "COMPLETED"}[i]
}
func ParseThreatAssessmentStatus(v string) (interface{}, error) {
    result := PENDING_THREATASSESSMENTSTATUS
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_THREATASSESSMENTSTATUS
        case "COMPLETED":
            result = COMPLETED_THREATASSESSMENTSTATUS
        default:
            return 0, errors.New("Unknown ThreatAssessmentStatus value: " + v)
    }
    return &result, nil
}
func SerializeThreatAssessmentStatus(values []ThreatAssessmentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
