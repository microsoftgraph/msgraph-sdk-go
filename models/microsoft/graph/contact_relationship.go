package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "PARENT":
            return PARENT_CONTACTRELATIONSHIP, nil
        case "RELATIVE":
            return RELATIVE_CONTACTRELATIONSHIP, nil
        case "AIDE":
            return AIDE_CONTACTRELATIONSHIP, nil
        case "DOCTOR":
            return DOCTOR_CONTACTRELATIONSHIP, nil
        case "GUARDIAN":
            return GUARDIAN_CONTACTRELATIONSHIP, nil
        case "CHILD":
            return CHILD_CONTACTRELATIONSHIP, nil
        case "OTHER":
            return OTHER_CONTACTRELATIONSHIP, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONTACTRELATIONSHIP, nil
    }
    return 0, errors.New("Unknown ContactRelationship value: " + v)
}
func SerializeContactRelationship(values []ContactRelationship) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
