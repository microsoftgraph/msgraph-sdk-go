package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_RECORDINGSTATUS, nil
        case "NOTRECORDING":
            return NOTRECORDING_RECORDINGSTATUS, nil
        case "RECORDING":
            return RECORDING_RECORDINGSTATUS, nil
        case "FAILED":
            return FAILED_RECORDINGSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_RECORDINGSTATUS, nil
    }
    return 0, errors.New("Unknown RecordingStatus value: " + v)
}
func SerializeRecordingStatus(values []RecordingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
