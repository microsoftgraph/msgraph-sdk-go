package graph
import (
    "strings"
    "errors"
)
// 
type RiskDetectionTimingType int

const (
    NOTDEFINED_RISKDETECTIONTIMINGTYPE RiskDetectionTimingType = iota
    REALTIME_RISKDETECTIONTIMINGTYPE
    NEARREALTIME_RISKDETECTIONTIMINGTYPE
    OFFLINE_RISKDETECTIONTIMINGTYPE
    UNKNOWNFUTUREVALUE_RISKDETECTIONTIMINGTYPE
)

func (i RiskDetectionTimingType) String() string {
    return []string{"NOTDEFINED", "REALTIME", "NEARREALTIME", "OFFLINE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRiskDetectionTimingType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTDEFINED":
            return NOTDEFINED_RISKDETECTIONTIMINGTYPE, nil
        case "REALTIME":
            return REALTIME_RISKDETECTIONTIMINGTYPE, nil
        case "NEARREALTIME":
            return NEARREALTIME_RISKDETECTIONTIMINGTYPE, nil
        case "OFFLINE":
            return OFFLINE_RISKDETECTIONTIMINGTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_RISKDETECTIONTIMINGTYPE, nil
    }
    return 0, errors.New("Unknown RiskDetectionTimingType value: " + v)
}
func SerializeRiskDetectionTimingType(values []RiskDetectionTimingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
