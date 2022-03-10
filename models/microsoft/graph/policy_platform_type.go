package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type PolicyPlatformType int

const (
    ANDROID_POLICYPLATFORMTYPE PolicyPlatformType = iota
    ANDROIDFORWORK_POLICYPLATFORMTYPE
    IOS_POLICYPLATFORMTYPE
    MACOS_POLICYPLATFORMTYPE
    WINDOWSPHONE81_POLICYPLATFORMTYPE
    WINDOWS81ANDLATER_POLICYPLATFORMTYPE
    WINDOWS10ANDLATER_POLICYPLATFORMTYPE
    ALL_POLICYPLATFORMTYPE
)

func (i PolicyPlatformType) String() string {
    return []string{"ANDROID", "ANDROIDFORWORK", "IOS", "MACOS", "WINDOWSPHONE81", "WINDOWS81ANDLATER", "WINDOWS10ANDLATER", "ALL"}[i]
}
func ParsePolicyPlatformType(v string) (interface{}, error) {
    result := ANDROID_POLICYPLATFORMTYPE
    switch strings.ToUpper(v) {
        case "ANDROID":
            result = ANDROID_POLICYPLATFORMTYPE
        case "ANDROIDFORWORK":
            result = ANDROIDFORWORK_POLICYPLATFORMTYPE
        case "IOS":
            result = IOS_POLICYPLATFORMTYPE
        case "MACOS":
            result = MACOS_POLICYPLATFORMTYPE
        case "WINDOWSPHONE81":
            result = WINDOWSPHONE81_POLICYPLATFORMTYPE
        case "WINDOWS81ANDLATER":
            result = WINDOWS81ANDLATER_POLICYPLATFORMTYPE
        case "WINDOWS10ANDLATER":
            result = WINDOWS10ANDLATER_POLICYPLATFORMTYPE
        case "ALL":
            result = ALL_POLICYPLATFORMTYPE
        default:
            return 0, errors.New("Unknown PolicyPlatformType value: " + v)
    }
    return &result, nil
}
func SerializePolicyPlatformType(values []PolicyPlatformType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
