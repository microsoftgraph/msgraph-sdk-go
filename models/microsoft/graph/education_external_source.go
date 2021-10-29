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
    switch strings.ToUpper(v) {
        case "SIS":
            return SIS_EDUCATIONEXTERNALSOURCE, nil
        case "MANUAL":
            return MANUAL_EDUCATIONEXTERNALSOURCE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EDUCATIONEXTERNALSOURCE, nil
    }
    return 0, errors.New("Unknown EducationExternalSource value: " + v)
}
func SerializeEducationExternalSource(values []EducationExternalSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
