package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type ManagedAppFlaggedReason int

const (
    NONE_MANAGEDAPPFLAGGEDREASON ManagedAppFlaggedReason = iota
    ROOTEDDEVICE_MANAGEDAPPFLAGGEDREASON
)

func (i ManagedAppFlaggedReason) String() string {
    return []string{"NONE", "ROOTEDDEVICE"}[i]
}
func ParseManagedAppFlaggedReason(v string) (interface{}, error) {
    result := NONE_MANAGEDAPPFLAGGEDREASON
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_MANAGEDAPPFLAGGEDREASON
        case "ROOTEDDEVICE":
            result = ROOTEDDEVICE_MANAGEDAPPFLAGGEDREASON
        default:
            return 0, errors.New("Unknown ManagedAppFlaggedReason value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppFlaggedReason(values []ManagedAppFlaggedReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
