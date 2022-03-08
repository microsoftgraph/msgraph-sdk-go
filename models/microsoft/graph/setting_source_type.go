package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type SettingSourceType int

const (
    DEVICECONFIGURATION_SETTINGSOURCETYPE SettingSourceType = iota
    DEVICEINTENT_SETTINGSOURCETYPE
)

func (i SettingSourceType) String() string {
    return []string{"DEVICECONFIGURATION", "DEVICEINTENT"}[i]
}
func ParseSettingSourceType(v string) (interface{}, error) {
    result := DEVICECONFIGURATION_SETTINGSOURCETYPE
    switch strings.ToUpper(v) {
        case "DEVICECONFIGURATION":
            result = DEVICECONFIGURATION_SETTINGSOURCETYPE
        case "DEVICEINTENT":
            result = DEVICEINTENT_SETTINGSOURCETYPE
        default:
            return 0, errors.New("Unknown SettingSourceType value: " + v)
    }
    return &result, nil
}
func SerializeSettingSourceType(values []SettingSourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
