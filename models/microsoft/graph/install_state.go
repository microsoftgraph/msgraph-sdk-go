package graph
import (
    "strings"
    "errors"
)
// 
type InstallState int

const (
    NOTAPPLICABLE_INSTALLSTATE InstallState = iota
    INSTALLED_INSTALLSTATE
    FAILED_INSTALLSTATE
    NOTINSTALLED_INSTALLSTATE
    UNINSTALLFAILED_INSTALLSTATE
    UNKNOWN_INSTALLSTATE
)

func (i InstallState) String() string {
    return []string{"NOTAPPLICABLE", "INSTALLED", "FAILED", "NOTINSTALLED", "UNINSTALLFAILED", "UNKNOWN"}[i]
}
func ParseInstallState(v string) (interface{}, error) {
    result := NOTAPPLICABLE_INSTALLSTATE
    switch strings.ToUpper(v) {
        case "NOTAPPLICABLE":
            result = NOTAPPLICABLE_INSTALLSTATE
        case "INSTALLED":
            result = INSTALLED_INSTALLSTATE
        case "FAILED":
            result = FAILED_INSTALLSTATE
        case "NOTINSTALLED":
            result = NOTINSTALLED_INSTALLSTATE
        case "UNINSTALLFAILED":
            result = UNINSTALLFAILED_INSTALLSTATE
        case "UNKNOWN":
            result = UNKNOWN_INSTALLSTATE
        default:
            return 0, errors.New("Unknown InstallState value: " + v)
    }
    return &result, nil
}
func SerializeInstallState(values []InstallState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
