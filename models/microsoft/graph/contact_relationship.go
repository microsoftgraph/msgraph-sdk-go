package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
type ContactRelationship int

const (
    PARENT_CONTACTRELATIONSHIP ContactRelationship = iota
    RELATIVE_CONTACTRELATIONSHIP
    AIDE_CONTACTRELATIONSHIP
    DOCTOR_CONTACTRELATIONSHIP
    GUARDIAN_CONTACTRELATIONSHIP
    CHILD_CONTACTRELATIONSHIP
    OTHER_CONTACTRELATIONSHIP
    UNKNOWNFUTUREVALUE_CONTACTRELATIONSHIP
)

func (i ContactRelationship) String() string {
    return []string{"PARENT", "RELATIVE", "AIDE", "DOCTOR", "GUARDIAN", "CHILD", "OTHER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseContactRelationship(v string) (interface{}, error) {
    result := PARENT_CONTACTRELATIONSHIP
    switch strings.ToUpper(v) {
        case "PARENT":
            result = PARENT_CONTACTRELATIONSHIP
        case "RELATIVE":
            result = RELATIVE_CONTACTRELATIONSHIP
        case "AIDE":
            result = AIDE_CONTACTRELATIONSHIP
        case "DOCTOR":
            result = DOCTOR_CONTACTRELATIONSHIP
        case "GUARDIAN":
            result = GUARDIAN_CONTACTRELATIONSHIP
        case "CHILD":
            result = CHILD_CONTACTRELATIONSHIP
        case "OTHER":
            result = OTHER_CONTACTRELATIONSHIP
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONTACTRELATIONSHIP
        default:
            return 0, errors.New("Unknown ContactRelationship value: " + v)
    }
    return &result, nil
}
func SerializeContactRelationship(values []ContactRelationship) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
