package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityGovernance singleton.
type AccessPackageAssignmentState int

const (
    DELIVERING_ACCESSPACKAGEASSIGNMENTSTATE AccessPackageAssignmentState = iota
    PARTIALLYDELIVERED_ACCESSPACKAGEASSIGNMENTSTATE
    DELIVERED_ACCESSPACKAGEASSIGNMENTSTATE
    EXPIRED_ACCESSPACKAGEASSIGNMENTSTATE
    DELIVERYFAILED_ACCESSPACKAGEASSIGNMENTSTATE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGEASSIGNMENTSTATE
)

func (i AccessPackageAssignmentState) String() string {
    return []string{"DELIVERING", "PARTIALLYDELIVERED", "DELIVERED", "EXPIRED", "DELIVERYFAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAccessPackageAssignmentState(v string) (interface{}, error) {
    result := DELIVERING_ACCESSPACKAGEASSIGNMENTSTATE
    switch strings.ToUpper(v) {
        case "DELIVERING":
            result = DELIVERING_ACCESSPACKAGEASSIGNMENTSTATE
        case "PARTIALLYDELIVERED":
            result = PARTIALLYDELIVERED_ACCESSPACKAGEASSIGNMENTSTATE
        case "DELIVERED":
            result = DELIVERED_ACCESSPACKAGEASSIGNMENTSTATE
        case "EXPIRED":
            result = EXPIRED_ACCESSPACKAGEASSIGNMENTSTATE
        case "DELIVERYFAILED":
            result = DELIVERYFAILED_ACCESSPACKAGEASSIGNMENTSTATE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGEASSIGNMENTSTATE
        default:
            return 0, errors.New("Unknown AccessPackageAssignmentState value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageAssignmentState(values []AccessPackageAssignmentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
