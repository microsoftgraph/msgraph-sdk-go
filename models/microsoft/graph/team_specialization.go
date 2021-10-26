package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_TEAMSPECIALIZATION, nil
        case "EDUCATIONSTANDARD":
            return EDUCATIONSTANDARD_TEAMSPECIALIZATION, nil
        case "EDUCATIONCLASS":
            return EDUCATIONCLASS_TEAMSPECIALIZATION, nil
        case "EDUCATIONPROFESSIONALLEARNINGCOMMUNITY":
            return EDUCATIONPROFESSIONALLEARNINGCOMMUNITY_TEAMSPECIALIZATION, nil
        case "EDUCATIONSTAFF":
            return EDUCATIONSTAFF_TEAMSPECIALIZATION, nil
        case "HEALTHCARESTANDARD":
            return HEALTHCARESTANDARD_TEAMSPECIALIZATION, nil
        case "HEALTHCARECARECOORDINATION":
            return HEALTHCARECARECOORDINATION_TEAMSPECIALIZATION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMSPECIALIZATION, nil
    }
    return 0, errors.New("Unknown TeamSpecialization value: " + v)
}
func SerializeTeamSpecialization(values []TeamSpecialization) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
