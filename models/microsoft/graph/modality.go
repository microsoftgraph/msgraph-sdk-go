package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the cloudCommunications singleton.
type Modality int

const (
    AUDIO_MODALITY Modality = iota
    VIDEO_MODALITY
    VIDEOBASEDSCREENSHARING_MODALITY
    DATA_MODALITY
    UNKNOWNFUTUREVALUE_MODALITY
)

func (i Modality) String() string {
    return []string{"AUDIO", "VIDEO", "VIDEOBASEDSCREENSHARING", "DATA", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseModality(v string) (interface{}, error) {
    result := AUDIO_MODALITY
    switch strings.ToUpper(v) {
        case "AUDIO":
            result = AUDIO_MODALITY
        case "VIDEO":
            result = VIDEO_MODALITY
        case "VIDEOBASEDSCREENSHARING":
            result = VIDEOBASEDSCREENSHARING_MODALITY
        case "DATA":
            result = DATA_MODALITY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MODALITY
        default:
            return 0, errors.New("Unknown Modality value: " + v)
    }
    return &result, nil
}
func SerializeModality(values []Modality) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
