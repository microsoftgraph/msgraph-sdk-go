package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the policyRoot singleton.
type ConditionalAccessDevicePlatform int

const (
    ANDROID_CONDITIONALACCESSDEVICEPLATFORM ConditionalAccessDevicePlatform = iota
    IOS_CONDITIONALACCESSDEVICEPLATFORM
    WINDOWS_CONDITIONALACCESSDEVICEPLATFORM
    WINDOWSPHONE_CONDITIONALACCESSDEVICEPLATFORM
    MACOS_CONDITIONALACCESSDEVICEPLATFORM
    ALL_CONDITIONALACCESSDEVICEPLATFORM
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSDEVICEPLATFORM
    LINUX_CONDITIONALACCESSDEVICEPLATFORM
)

func (i ConditionalAccessDevicePlatform) String() string {
    return []string{"ANDROID", "IOS", "WINDOWS", "WINDOWSPHONE", "MACOS", "ALL", "UNKNOWNFUTUREVALUE", "LINUX"}[i]
}
func ParseConditionalAccessDevicePlatform(v string) (interface{}, error) {
    result := ANDROID_CONDITIONALACCESSDEVICEPLATFORM
    switch strings.ToUpper(v) {
        case "ANDROID":
            result = ANDROID_CONDITIONALACCESSDEVICEPLATFORM
        case "IOS":
            result = IOS_CONDITIONALACCESSDEVICEPLATFORM
        case "WINDOWS":
            result = WINDOWS_CONDITIONALACCESSDEVICEPLATFORM
        case "WINDOWSPHONE":
            result = WINDOWSPHONE_CONDITIONALACCESSDEVICEPLATFORM
        case "MACOS":
            result = MACOS_CONDITIONALACCESSDEVICEPLATFORM
        case "ALL":
            result = ALL_CONDITIONALACCESSDEVICEPLATFORM
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONDITIONALACCESSDEVICEPLATFORM
        case "LINUX":
            result = LINUX_CONDITIONALACCESSDEVICEPLATFORM
        default:
            return 0, errors.New("Unknown ConditionalAccessDevicePlatform value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessDevicePlatform(values []ConditionalAccessDevicePlatform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
