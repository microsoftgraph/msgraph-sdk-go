package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the policyRoot singleton.
type StagedFeatureName int

const (
    PASSTHROUGHAUTHENTICATION_STAGEDFEATURENAME StagedFeatureName = iota
    SEAMLESSSSO_STAGEDFEATURENAME
    PASSWORDHASHSYNC_STAGEDFEATURENAME
    EMAILASALTERNATEID_STAGEDFEATURENAME
    UNKNOWNFUTUREVALUE_STAGEDFEATURENAME
    CERTIFICATEBASEDAUTHENTICATION_STAGEDFEATURENAME
    MULTIFACTORAUTHENTICATION_STAGEDFEATURENAME
)

func (i StagedFeatureName) String() string {
    return []string{"PASSTHROUGHAUTHENTICATION", "SEAMLESSSSO", "PASSWORDHASHSYNC", "EMAILASALTERNATEID", "UNKNOWNFUTUREVALUE", "CERTIFICATEBASEDAUTHENTICATION", "MULTIFACTORAUTHENTICATION"}[i]
}
func ParseStagedFeatureName(v string) (interface{}, error) {
    result := PASSTHROUGHAUTHENTICATION_STAGEDFEATURENAME
    switch strings.ToUpper(v) {
        case "PASSTHROUGHAUTHENTICATION":
            result = PASSTHROUGHAUTHENTICATION_STAGEDFEATURENAME
        case "SEAMLESSSSO":
            result = SEAMLESSSSO_STAGEDFEATURENAME
        case "PASSWORDHASHSYNC":
            result = PASSWORDHASHSYNC_STAGEDFEATURENAME
        case "EMAILASALTERNATEID":
            result = EMAILASALTERNATEID_STAGEDFEATURENAME
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_STAGEDFEATURENAME
        case "CERTIFICATEBASEDAUTHENTICATION":
            result = CERTIFICATEBASEDAUTHENTICATION_STAGEDFEATURENAME
        case "MULTIFACTORAUTHENTICATION":
            result = MULTIFACTORAUTHENTICATION_STAGEDFEATURENAME
        default:
            return 0, errors.New("Unknown StagedFeatureName value: " + v)
    }
    return &result, nil
}
func SerializeStagedFeatureName(values []StagedFeatureName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
