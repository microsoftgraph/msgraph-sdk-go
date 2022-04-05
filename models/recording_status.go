package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type RecordingStatus int

const (
    UNKNOWN_RECORDINGSTATUS RecordingStatus = iota
    NOTRECORDING_RECORDINGSTATUS
    RECORDING_RECORDINGSTATUS
    FAILED_RECORDINGSTATUS
    UNKNOWNFUTUREVALUE_RECORDINGSTATUS
)

func (i RecordingStatus) String() string {
    return []string{"UNKNOWN", "NOTRECORDING", "RECORDING", "FAILED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRecordingStatus(v string) (interface{}, error) {
    result := UNKNOWN_RECORDINGSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_RECORDINGSTATUS
        case "NOTRECORDING":
            result = NOTRECORDING_RECORDINGSTATUS
        case "RECORDING":
            result = RECORDING_RECORDINGSTATUS
        case "FAILED":
            result = FAILED_RECORDINGSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RECORDINGSTATUS
        default:
            return 0, errors.New("Unknown RecordingStatus value: " + v)
    }
    return &result, nil
}
func SerializeRecordingStatus(values []RecordingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
