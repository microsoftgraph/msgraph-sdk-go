package graph
import (
    "strings"
    "errors"
)
// 
type SettingSourceType int

const (
    DEVICECONFIGURATION_SETTINGSOURCETYPE SettingSourceType = iota
    DEVICEINTENT_SETTINGSOURCETYPE
)

func (i SettingSourceType) String() string {
    return []string{"DEVICECONFIGURATION", "DEVICEINTENT"}[i]
}
func ParseSettingSourceType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DEVICECONFIGURATION":
            return DEVICECONFIGURATION_SETTINGSOURCETYPE, nil
        case "DEVICEINTENT":
            return DEVICEINTENT_SETTINGSOURCETYPE, nil
    }
    return 0, errors.New("Unknown SettingSourceType value: " + v)
}
func SerializeSettingSourceType(values []SettingSourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
