package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "AUTOMATIC":
            return AUTOMATIC_PLANNERPREVIEWTYPE, nil
        case "NOPREVIEW":
            return NOPREVIEW_PLANNERPREVIEWTYPE, nil
        case "CHECKLIST":
            return CHECKLIST_PLANNERPREVIEWTYPE, nil
        case "DESCRIPTION":
            return DESCRIPTION_PLANNERPREVIEWTYPE, nil
        case "REFERENCE":
            return REFERENCE_PLANNERPREVIEWTYPE, nil
    }
    return 0, errors.New("Unknown PlannerPreviewType value: " + v)
}
func SerializePlannerPreviewType(values []PlannerPreviewType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
