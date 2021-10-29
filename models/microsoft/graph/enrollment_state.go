package graph
import (
    "strings"
    "errors"
)
// 
type EnrollmentState int

const (
    UNKNOWN_ENROLLMENTSTATE EnrollmentState = iota
    ENROLLED_ENROLLMENTSTATE
    PENDINGRESET_ENROLLMENTSTATE
    FAILED_ENROLLMENTSTATE
    NOTCONTACTED_ENROLLMENTSTATE
)

func (i EnrollmentState) String() string {
    return []string{"UNKNOWN", "ENROLLED", "PENDINGRESET", "FAILED", "NOTCONTACTED"}[i]
}
func ParseEnrollmentState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_ENROLLMENTSTATE, nil
        case "ENROLLED":
            return ENROLLED_ENROLLMENTSTATE, nil
        case "PENDINGRESET":
            return PENDINGRESET_ENROLLMENTSTATE, nil
        case "FAILED":
            return FAILED_ENROLLMENTSTATE, nil
        case "NOTCONTACTED":
            return NOTCONTACTED_ENROLLMENTSTATE, nil
    }
    return 0, errors.New("Unknown EnrollmentState value: " + v)
}
func SerializeEnrollmentState(values []EnrollmentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
