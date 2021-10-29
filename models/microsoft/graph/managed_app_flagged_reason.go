package graph
import (
    "strings"
    "errors"
)
// 
type ManagedAppFlaggedReason int

const (
    NONE_MANAGEDAPPFLAGGEDREASON ManagedAppFlaggedReason = iota
    ROOTEDDEVICE_MANAGEDAPPFLAGGEDREASON
)

func (i ManagedAppFlaggedReason) String() string {
    return []string{"NONE", "ROOTEDDEVICE"}[i]
}
func ParseManagedAppFlaggedReason(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_MANAGEDAPPFLAGGEDREASON, nil
        case "ROOTEDDEVICE":
            return ROOTEDDEVICE_MANAGEDAPPFLAGGEDREASON, nil
    }
    return 0, errors.New("Unknown ManagedAppFlaggedReason value: " + v)
}
func SerializeManagedAppFlaggedReason(values []ManagedAppFlaggedReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
