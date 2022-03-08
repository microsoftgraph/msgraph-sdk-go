package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of group entities.
type TeamSpecialization int

const (
    NONE_TEAMSPECIALIZATION TeamSpecialization = iota
    EDUCATIONSTANDARD_TEAMSPECIALIZATION
    EDUCATIONCLASS_TEAMSPECIALIZATION
    EDUCATIONPROFESSIONALLEARNINGCOMMUNITY_TEAMSPECIALIZATION
    EDUCATIONSTAFF_TEAMSPECIALIZATION
    HEALTHCARESTANDARD_TEAMSPECIALIZATION
    HEALTHCARECARECOORDINATION_TEAMSPECIALIZATION
    UNKNOWNFUTUREVALUE_TEAMSPECIALIZATION
)

func (i TeamSpecialization) String() string {
    return []string{"NONE", "EDUCATIONSTANDARD", "EDUCATIONCLASS", "EDUCATIONPROFESSIONALLEARNINGCOMMUNITY", "EDUCATIONSTAFF", "HEALTHCARESTANDARD", "HEALTHCARECARECOORDINATION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamSpecialization(v string) (interface{}, error) {
    result := NONE_TEAMSPECIALIZATION
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_TEAMSPECIALIZATION
        case "EDUCATIONSTANDARD":
            result = EDUCATIONSTANDARD_TEAMSPECIALIZATION
        case "EDUCATIONCLASS":
            result = EDUCATIONCLASS_TEAMSPECIALIZATION
        case "EDUCATIONPROFESSIONALLEARNINGCOMMUNITY":
            result = EDUCATIONPROFESSIONALLEARNINGCOMMUNITY_TEAMSPECIALIZATION
        case "EDUCATIONSTAFF":
            result = EDUCATIONSTAFF_TEAMSPECIALIZATION
        case "HEALTHCARESTANDARD":
            result = HEALTHCARESTANDARD_TEAMSPECIALIZATION
        case "HEALTHCARECARECOORDINATION":
            result = HEALTHCARECARECOORDINATION_TEAMSPECIALIZATION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMSPECIALIZATION
        default:
            return 0, errors.New("Unknown TeamSpecialization value: " + v)
    }
    return &result, nil
}
func SerializeTeamSpecialization(values []TeamSpecialization) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
