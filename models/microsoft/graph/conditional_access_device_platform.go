package graph
import (
    "strings"
    "errors"
)
type ConditionalAccessDevicePlatform int

const (
    ANDROID_CONDITIONALACCESSDEVICEPLATFORM ConditionalAccessDevicePlatform = iota
    IOS_CONDITIONALACCESSDEVICEPLATFORM
    WINDOWS_CONDITIONALACCESSDEVICEPLATFORM
    WINDOWSPHONE_CONDITIONALACCESSDEVICEPLATFORM
    MACOS_CONDITIONALACCESSDEVICEPLATFORM
    ALL_CONDITIONALACCESSDEVICEPLATFORM
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSDEVICEPLATFORM
)

func (i ConditionalAccessDevicePlatform) String() string {
    return []string{"ANDROID", "IOS", "WINDOWS", "WINDOWSPHONE", "MACOS", "ALL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConditionalAccessDevicePlatform(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ANDROID":
            return ANDROID_CONDITIONALACCESSDEVICEPLATFORM, nil
        case "IOS":
            return IOS_CONDITIONALACCESSDEVICEPLATFORM, nil
        case "WINDOWS":
            return WINDOWS_CONDITIONALACCESSDEVICEPLATFORM, nil
        case "WINDOWSPHONE":
            return WINDOWSPHONE_CONDITIONALACCESSDEVICEPLATFORM, nil
        case "MACOS":
            return MACOS_CONDITIONALACCESSDEVICEPLATFORM, nil
        case "ALL":
            return ALL_CONDITIONALACCESSDEVICEPLATFORM, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONDITIONALACCESSDEVICEPLATFORM, nil
    }
    return 0, errors.New("Unknown ConditionalAccessDevicePlatform value: " + v)
}
func SerializeConditionalAccessDevicePlatform(values []ConditionalAccessDevicePlatform) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
