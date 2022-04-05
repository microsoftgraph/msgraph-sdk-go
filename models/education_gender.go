package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
type EducationGender int

const (
    FEMALE_EDUCATIONGENDER EducationGender = iota
    MALE_EDUCATIONGENDER
    OTHER_EDUCATIONGENDER
    UNKNOWNFUTUREVALUE_EDUCATIONGENDER
)

func (i EducationGender) String() string {
    return []string{"FEMALE", "MALE", "OTHER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEducationGender(v string) (interface{}, error) {
    result := FEMALE_EDUCATIONGENDER
    switch strings.ToUpper(v) {
        case "FEMALE":
            result = FEMALE_EDUCATIONGENDER
        case "MALE":
            result = MALE_EDUCATIONGENDER
        case "OTHER":
            result = OTHER_EDUCATIONGENDER
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONGENDER
        default:
            return 0, errors.New("Unknown EducationGender value: " + v)
    }
    return &result, nil
}
func SerializeEducationGender(values []EducationGender) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
