package graph
import (
    "strings"
    "errors"
)
// 
type StagedFeatureName int

const (
    PASSTHROUGHAUTHENTICATION_STAGEDFEATURENAME StagedFeatureName = iota
    SEAMLESSSSO_STAGEDFEATURENAME
    PASSWORDHASHSYNC_STAGEDFEATURENAME
    EMAILASALTERNATEID_STAGEDFEATURENAME
    UNKNOWNFUTUREVALUE_STAGEDFEATURENAME
)

func (i StagedFeatureName) String() string {
    return []string{"PASSTHROUGHAUTHENTICATION", "SEAMLESSSSO", "PASSWORDHASHSYNC", "EMAILASALTERNATEID", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseStagedFeatureName(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "PASSTHROUGHAUTHENTICATION":
            return PASSTHROUGHAUTHENTICATION_STAGEDFEATURENAME, nil
        case "SEAMLESSSSO":
            return SEAMLESSSSO_STAGEDFEATURENAME, nil
        case "PASSWORDHASHSYNC":
            return PASSWORDHASHSYNC_STAGEDFEATURENAME, nil
        case "EMAILASALTERNATEID":
            return EMAILASALTERNATEID_STAGEDFEATURENAME, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_STAGEDFEATURENAME, nil
    }
    return 0, errors.New("Unknown StagedFeatureName value: " + v)
}
func SerializeStagedFeatureName(values []StagedFeatureName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
