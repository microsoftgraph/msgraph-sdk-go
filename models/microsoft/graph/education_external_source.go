package graph
import (
    "strings"
    "errors"
)
// 
type EducationExternalSource int

const (
    SIS_EDUCATIONEXTERNALSOURCE EducationExternalSource = iota
    MANUAL_EDUCATIONEXTERNALSOURCE
    UNKNOWNFUTUREVALUE_EDUCATIONEXTERNALSOURCE
)

func (i EducationExternalSource) String() string {
    return []string{"SIS", "MANUAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseEducationExternalSource(v string) (interface{}, error) {
    result := SIS_EDUCATIONEXTERNALSOURCE
    switch strings.ToUpper(v) {
        case "SIS":
            result = SIS_EDUCATIONEXTERNALSOURCE
        case "MANUAL":
            result = MANUAL_EDUCATIONEXTERNALSOURCE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONEXTERNALSOURCE
        default:
            return 0, errors.New("Unknown EducationExternalSource value: " + v)
    }
    return &result, nil
}
func SerializeEducationExternalSource(values []EducationExternalSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
