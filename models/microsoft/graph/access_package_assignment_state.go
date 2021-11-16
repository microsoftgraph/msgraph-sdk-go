package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "DELIVERING":
            return DELIVERING_ACCESSPACKAGEASSIGNMENTSTATE, nil
        case "PARTIALLYDELIVERED":
            return PARTIALLYDELIVERED_ACCESSPACKAGEASSIGNMENTSTATE, nil
        case "DELIVERED":
            return DELIVERED_ACCESSPACKAGEASSIGNMENTSTATE, nil
        case "EXPIRED":
            return EXPIRED_ACCESSPACKAGEASSIGNMENTSTATE, nil
        case "DELIVERYFAILED":
            return DELIVERYFAILED_ACCESSPACKAGEASSIGNMENTSTATE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ACCESSPACKAGEASSIGNMENTSTATE, nil
    }
    return 0, errors.New("Unknown AccessPackageAssignmentState value: " + v)
}
func SerializeAccessPackageAssignmentState(values []AccessPackageAssignmentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
