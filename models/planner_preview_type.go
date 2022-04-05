package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type PlannerPreviewType int

const (
    AUTOMATIC_PLANNERPREVIEWTYPE PlannerPreviewType = iota
    NOPREVIEW_PLANNERPREVIEWTYPE
    CHECKLIST_PLANNERPREVIEWTYPE
    DESCRIPTION_PLANNERPREVIEWTYPE
    REFERENCE_PLANNERPREVIEWTYPE
)

func (i PlannerPreviewType) String() string {
    return []string{"AUTOMATIC", "NOPREVIEW", "CHECKLIST", "DESCRIPTION", "REFERENCE"}[i]
}
func ParsePlannerPreviewType(v string) (interface{}, error) {
    result := AUTOMATIC_PLANNERPREVIEWTYPE
    switch strings.ToUpper(v) {
        case "AUTOMATIC":
            result = AUTOMATIC_PLANNERPREVIEWTYPE
        case "NOPREVIEW":
            result = NOPREVIEW_PLANNERPREVIEWTYPE
        case "CHECKLIST":
            result = CHECKLIST_PLANNERPREVIEWTYPE
        case "DESCRIPTION":
            result = DESCRIPTION_PLANNERPREVIEWTYPE
        case "REFERENCE":
            result = REFERENCE_PLANNERPREVIEWTYPE
        default:
            return 0, errors.New("Unknown PlannerPreviewType value: " + v)
    }
    return &result, nil
}
func SerializePlannerPreviewType(values []PlannerPreviewType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
