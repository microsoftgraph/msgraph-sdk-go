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
    switch strings.ToUpper(v) {
        case "NOTAPPLICABLE":
            return NOTAPPLICABLE_INSTALLSTATE, nil
        case "INSTALLED":
            return INSTALLED_INSTALLSTATE, nil
        case "FAILED":
            return FAILED_INSTALLSTATE, nil
        case "NOTINSTALLED":
            return NOTINSTALLED_INSTALLSTATE, nil
        case "UNINSTALLFAILED":
            return UNINSTALLFAILED_INSTALLSTATE, nil
        case "UNKNOWN":
            return UNKNOWN_INSTALLSTATE, nil
    }
    return 0, errors.New("Unknown InstallState value: " + v)
}
func SerializeInstallState(values []InstallState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
