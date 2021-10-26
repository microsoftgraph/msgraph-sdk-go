package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "CHECKPOLICY":
            return CHECKPOLICY_THREATASSESSMENTRESULTTYPE, nil
        case "RESCAN":
            return RESCAN_THREATASSESSMENTRESULTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_THREATASSESSMENTRESULTTYPE, nil
    }
    return 0, errors.New("Unknown ThreatAssessmentResultType value: " + v)
}
func SerializeThreatAssessmentResultType(values []ThreatAssessmentResultType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
