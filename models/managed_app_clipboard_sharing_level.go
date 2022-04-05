package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type ManagedAppClipboardSharingLevel int

const (
    ALLAPPS_MANAGEDAPPCLIPBOARDSHARINGLEVEL ManagedAppClipboardSharingLevel = iota
    MANAGEDAPPSWITHPASTEIN_MANAGEDAPPCLIPBOARDSHARINGLEVEL
    MANAGEDAPPS_MANAGEDAPPCLIPBOARDSHARINGLEVEL
    BLOCKED_MANAGEDAPPCLIPBOARDSHARINGLEVEL
)

func (i ManagedAppClipboardSharingLevel) String() string {
    return []string{"ALLAPPS", "MANAGEDAPPSWITHPASTEIN", "MANAGEDAPPS", "BLOCKED"}[i]
}
func ParseManagedAppClipboardSharingLevel(v string) (interface{}, error) {
    result := ALLAPPS_MANAGEDAPPCLIPBOARDSHARINGLEVEL
    switch strings.ToUpper(v) {
        case "ALLAPPS":
            result = ALLAPPS_MANAGEDAPPCLIPBOARDSHARINGLEVEL
        case "MANAGEDAPPSWITHPASTEIN":
            result = MANAGEDAPPSWITHPASTEIN_MANAGEDAPPCLIPBOARDSHARINGLEVEL
        case "MANAGEDAPPS":
            result = MANAGEDAPPS_MANAGEDAPPCLIPBOARDSHARINGLEVEL
        case "BLOCKED":
            result = BLOCKED_MANAGEDAPPCLIPBOARDSHARINGLEVEL
        default:
            return 0, errors.New("Unknown ManagedAppClipboardSharingLevel value: " + v)
    }
    return &result, nil
}
func SerializeManagedAppClipboardSharingLevel(values []ManagedAppClipboardSharingLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
