package models
import (
    "errors"
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
    return []string{"None", "Add", "Update", "Delete"}[i]
}
func ParseObjectFlowTypes(v string) (any, error) {
    result := NONE_OBJECTFLOWTYPES
    switch v {
        case "None":
            result = NONE_OBJECTFLOWTYPES
        case "Add":
            result = ADD_OBJECTFLOWTYPES
        case "Update":
            result = UPDATE_OBJECTFLOWTYPES
        case "Delete":
            result = DELETE_OBJECTFLOWTYPES
        default:
            return 0, errors.New("Unknown ObjectFlowTypes value: " + v)
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
