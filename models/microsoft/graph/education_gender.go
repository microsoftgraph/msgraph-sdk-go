package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "FEMALE":
            return FEMALE_EDUCATIONGENDER, nil
        case "MALE":
            return MALE_EDUCATIONGENDER, nil
        case "OTHER":
            return OTHER_EDUCATIONGENDER, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EDUCATIONGENDER, nil
    }
    return 0, errors.New("Unknown EducationGender value: " + v)
}
func SerializeEducationGender(values []EducationGender) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
