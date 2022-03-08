package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the informationProtection singleton.
type ThreatAssessmentResultType int

const (
    CHECKPOLICY_THREATASSESSMENTRESULTTYPE ThreatAssessmentResultType = iota
    RESCAN_THREATASSESSMENTRESULTTYPE
    UNKNOWNFUTUREVALUE_THREATASSESSMENTRESULTTYPE
)

func (i ThreatAssessmentResultType) String() string {
    return []string{"CHECKPOLICY", "RESCAN", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseThreatAssessmentResultType(v string) (interface{}, error) {
    result := CHECKPOLICY_THREATASSESSMENTRESULTTYPE
    switch strings.ToUpper(v) {
        case "CHECKPOLICY":
            result = CHECKPOLICY_THREATASSESSMENTRESULTTYPE
        case "RESCAN":
            result = RESCAN_THREATASSESSMENTRESULTTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_THREATASSESSMENTRESULTTYPE
        default:
            return 0, errors.New("Unknown ThreatAssessmentResultType value: " + v)
    }
    return &result, nil
}
func SerializeThreatAssessmentResultType(values []ThreatAssessmentResultType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
