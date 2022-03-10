package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityProtectionRoot singleton.
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
    result := NOTDEFINED_RISKDETECTIONTIMINGTYPE
    switch strings.ToUpper(v) {
        case "NOTDEFINED":
            result = NOTDEFINED_RISKDETECTIONTIMINGTYPE
        case "REALTIME":
            result = REALTIME_RISKDETECTIONTIMINGTYPE
        case "NEARREALTIME":
            result = NEARREALTIME_RISKDETECTIONTIMINGTYPE
        case "OFFLINE":
            result = OFFLINE_RISKDETECTIONTIMINGTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RISKDETECTIONTIMINGTYPE
        default:
            return 0, errors.New("Unknown RiskDetectionTimingType value: " + v)
    }
    return &result, nil
}
func SerializeRiskDetectionTimingType(values []RiskDetectionTimingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
