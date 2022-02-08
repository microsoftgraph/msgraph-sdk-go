package graph
import (
    "strings"
    "errors"
)
// 
type InstallIntent int

const (
    AVAILABLE_INSTALLINTENT InstallIntent = iota
    REQUIRED_INSTALLINTENT
    UNINSTALL_INSTALLINTENT
    AVAILABLEWITHOUTENROLLMENT_INSTALLINTENT
)

func (i InstallIntent) String() string {
    return []string{"AVAILABLE", "REQUIRED", "UNINSTALL", "AVAILABLEWITHOUTENROLLMENT"}[i]
}
func ParseInstallIntent(v string) (interface{}, error) {
    result := AVAILABLE_INSTALLINTENT
    switch strings.ToUpper(v) {
        case "AVAILABLE":
            result = AVAILABLE_INSTALLINTENT
        case "REQUIRED":
            result = REQUIRED_INSTALLINTENT
        case "UNINSTALL":
            result = UNINSTALL_INSTALLINTENT
        case "AVAILABLEWITHOUTENROLLMENT":
            result = AVAILABLEWITHOUTENROLLMENT_INSTALLINTENT
        default:
            return 0, errors.New("Unknown InstallIntent value: " + v)
    }
    return &result, nil
}
func SerializeInstallIntent(values []InstallIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
