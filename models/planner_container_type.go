package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type PlannerContainerType int

const (
    GROUP_PLANNERCONTAINERTYPE PlannerContainerType = iota
    UNKNOWNFUTUREVALUE_PLANNERCONTAINERTYPE
    ROSTER_PLANNERCONTAINERTYPE
)

func (i PlannerContainerType) String() string {
    return []string{"GROUP", "UNKNOWNFUTUREVALUE", "ROSTER"}[i]
}
func ParsePlannerContainerType(v string) (interface{}, error) {
    result := GROUP_PLANNERCONTAINERTYPE
    switch strings.ToUpper(v) {
        case "GROUP":
            result = GROUP_PLANNERCONTAINERTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_PLANNERCONTAINERTYPE
        case "ROSTER":
            result = ROSTER_PLANNERCONTAINERTYPE
        default:
            return 0, errors.New("Unknown PlannerContainerType value: " + v)
    }
    return &result, nil
}
func SerializePlannerContainerType(values []PlannerContainerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
