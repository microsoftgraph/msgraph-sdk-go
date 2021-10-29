package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "ANDROID":
            return ANDROID_POLICYPLATFORMTYPE, nil
        case "ANDROIDFORWORK":
            return ANDROIDFORWORK_POLICYPLATFORMTYPE, nil
        case "IOS":
            return IOS_POLICYPLATFORMTYPE, nil
        case "MACOS":
            return MACOS_POLICYPLATFORMTYPE, nil
        case "WINDOWSPHONE81":
            return WINDOWSPHONE81_POLICYPLATFORMTYPE, nil
        case "WINDOWS81ANDLATER":
            return WINDOWS81ANDLATER_POLICYPLATFORMTYPE, nil
        case "WINDOWS10ANDLATER":
            return WINDOWS10ANDLATER_POLICYPLATFORMTYPE, nil
        case "ALL":
            return ALL_POLICYPLATFORMTYPE, nil
    }
    return 0, errors.New("Unknown PolicyPlatformType value: " + v)
}
func SerializePolicyPlatformType(values []PolicyPlatformType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
