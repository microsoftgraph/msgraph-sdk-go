package models
import (
    "errors"
    "strings"
)
// Type of start menu app list visibility.
type WindowsStartMenuAppListVisibilityType int

const (
    // User defined. Default value.
    USERDEFINED_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE WindowsStartMenuAppListVisibilityType = iota
    // Collapse the app list on the start menu.
    COLLAPSE_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
    // Removes the app list entirely from the start menu.
    REMOVE_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
    // Disables the corresponding toggle (Collapse or Remove) in the Settings app.
    DISABLESETTINGSAPP_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
)

func (i WindowsStartMenuAppListVisibilityType) String() string {
    var values []string
    for p := WindowsStartMenuAppListVisibilityType(1); p <= DISABLESETTINGSAPP_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"userDefined", "collapse", "remove", "disableSettingsApp"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWindowsStartMenuAppListVisibilityType(v string) (any, error) {
    var result WindowsStartMenuAppListVisibilityType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "userDefined":
                result |= USERDEFINED_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
            case "collapse":
                result |= COLLAPSE_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
            case "remove":
                result |= REMOVE_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
            case "disableSettingsApp":
                result |= DISABLESETTINGSAPP_WINDOWSSTARTMENUAPPLISTVISIBILITYTYPE
            default:
                return 0, errors.New("Unknown WindowsStartMenuAppListVisibilityType value: " + v)
        }
    }
    return &result, nil
}
func SerializeWindowsStartMenuAppListVisibilityType(values []WindowsStartMenuAppListVisibilityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsStartMenuAppListVisibilityType) isMultiValue() bool {
    return true
}
