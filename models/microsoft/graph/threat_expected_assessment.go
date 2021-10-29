package graph
import (
    "strings"
    "errors"
)
// 
type ThreatExpectedAssessment int

const (
    BLOCK_THREATEXPECTEDASSESSMENT ThreatExpectedAssessment = iota
    UNBLOCK_THREATEXPECTEDASSESSMENT
)

func (i ThreatExpectedAssessment) String() string {
    return []string{"BLOCK", "UNBLOCK"}[i]
}
func ParseThreatExpectedAssessment(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "BLOCK":
            return BLOCK_THREATEXPECTEDASSESSMENT, nil
        case "UNBLOCK":
            return UNBLOCK_THREATEXPECTEDASSESSMENT, nil
    }
    return 0, errors.New("Unknown ThreatExpectedAssessment value: " + v)
}
func SerializeThreatExpectedAssessment(values []ThreatExpectedAssessment) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
