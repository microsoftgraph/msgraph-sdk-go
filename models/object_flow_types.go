package models
import (
    "errors"
    "strings"
)
// 
type ObjectFlowTypes int

const (
    NONE_OBJECTFLOWTYPES ObjectFlowTypes = iota
    ADD_OBJECTFLOWTYPES
    UPDATE_OBJECTFLOWTYPES
    DELETE_OBJECTFLOWTYPES
)

func (i ObjectFlowTypes) String() string {
    var values []string
    for p := ObjectFlowTypes(1); p <= DELETE_OBJECTFLOWTYPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"None", "Add", "Update", "Delete"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseObjectFlowTypes(v string) (any, error) {
    var result ObjectFlowTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "None":
                result |= NONE_OBJECTFLOWTYPES
            case "Add":
                result |= ADD_OBJECTFLOWTYPES
            case "Update":
                result |= UPDATE_OBJECTFLOWTYPES
            case "Delete":
                result |= DELETE_OBJECTFLOWTYPES
            default:
                return 0, errors.New("Unknown ObjectFlowTypes value: " + v)
        }
    }
    return &result, nil
}
func SerializeObjectFlowTypes(values []ObjectFlowTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ObjectFlowTypes) isMultiValue() bool {
    return true
}
