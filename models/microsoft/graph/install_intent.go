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
    switch strings.ToUpper(v) {
        case "AVAILABLE":
            return AVAILABLE_INSTALLINTENT, nil
        case "REQUIRED":
            return REQUIRED_INSTALLINTENT, nil
        case "UNINSTALL":
            return UNINSTALL_INSTALLINTENT, nil
        case "AVAILABLEWITHOUTENROLLMENT":
            return AVAILABLEWITHOUTENROLLMENT_INSTALLINTENT, nil
    }
    return 0, errors.New("Unknown InstallIntent value: " + v)
}
func SerializeInstallIntent(values []InstallIntent) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
