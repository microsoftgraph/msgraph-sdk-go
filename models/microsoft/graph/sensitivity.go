package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NORMAL":
            return NORMAL_SENSITIVITY, nil
        case "PERSONAL":
            return PERSONAL_SENSITIVITY, nil
        case "PRIVATE":
            return PRIVATE_SENSITIVITY, nil
        case "CONFIDENTIAL":
            return CONFIDENTIAL_SENSITIVITY, nil
    }
    return 0, errors.New("Unknown Sensitivity value: " + v)
}
func SerializeSensitivity(values []Sensitivity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
