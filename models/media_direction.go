package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type MediaDirection int

const (
    INACTIVE_MEDIADIRECTION MediaDirection = iota
    SENDONLY_MEDIADIRECTION
    RECEIVEONLY_MEDIADIRECTION
    SENDRECEIVE_MEDIADIRECTION
)

func (i MediaDirection) String() string {
    return []string{"INACTIVE", "SENDONLY", "RECEIVEONLY", "SENDRECEIVE"}[i]
}
func ParseMediaDirection(v string) (interface{}, error) {
    result := INACTIVE_MEDIADIRECTION
    switch strings.ToUpper(v) {
        case "INACTIVE":
            result = INACTIVE_MEDIADIRECTION
        case "SENDONLY":
            result = SENDONLY_MEDIADIRECTION
        case "RECEIVEONLY":
            result = RECEIVEONLY_MEDIADIRECTION
        case "SENDRECEIVE":
            result = SENDRECEIVE_MEDIADIRECTION
        default:
            return 0, errors.New("Unknown MediaDirection value: " + v)
    }
    return &result, nil
}
func SerializeMediaDirection(values []MediaDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
