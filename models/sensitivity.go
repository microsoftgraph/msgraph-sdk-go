package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type Sensitivity int

const (
    NORMAL_SENSITIVITY Sensitivity = iota
    PERSONAL_SENSITIVITY
    PRIVATE_SENSITIVITY
    CONFIDENTIAL_SENSITIVITY
)

func (i Sensitivity) String() string {
    return []string{"NORMAL", "PERSONAL", "PRIVATE", "CONFIDENTIAL"}[i]
}
func ParseSensitivity(v string) (interface{}, error) {
    result := NORMAL_SENSITIVITY
    switch strings.ToUpper(v) {
        case "NORMAL":
            result = NORMAL_SENSITIVITY
        case "PERSONAL":
            result = PERSONAL_SENSITIVITY
        case "PRIVATE":
            result = PRIVATE_SENSITIVITY
        case "CONFIDENTIAL":
            result = CONFIDENTIAL_SENSITIVITY
        default:
            return 0, errors.New("Unknown Sensitivity value: " + v)
    }
    return &result, nil
}
func SerializeSensitivity(values []Sensitivity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
