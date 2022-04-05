package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
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
    result := UNKNOWN_ENROLLMENTSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_ENROLLMENTSTATE
        case "ENROLLED":
            result = ENROLLED_ENROLLMENTSTATE
        case "PENDINGRESET":
            result = PENDINGRESET_ENROLLMENTSTATE
        case "FAILED":
            result = FAILED_ENROLLMENTSTATE
        case "NOTCONTACTED":
            result = NOTCONTACTED_ENROLLMENTSTATE
        default:
            return 0, errors.New("Unknown EnrollmentState value: " + v)
    }
    return &result, nil
}
func SerializeEnrollmentState(values []EnrollmentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
