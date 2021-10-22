package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "ALLAPPS":
            return ALLAPPS_MANAGEDAPPCLIPBOARDSHARINGLEVEL, nil
        case "MANAGEDAPPSWITHPASTEIN":
            return MANAGEDAPPSWITHPASTEIN_MANAGEDAPPCLIPBOARDSHARINGLEVEL, nil
        case "MANAGEDAPPS":
            return MANAGEDAPPS_MANAGEDAPPCLIPBOARDSHARINGLEVEL, nil
        case "BLOCKED":
            return BLOCKED_MANAGEDAPPCLIPBOARDSHARINGLEVEL, nil
    }
    return 0, errors.New("Unknown ManagedAppClipboardSharingLevel value: " + v)
}
func SerializeManagedAppClipboardSharingLevel(values []ManagedAppClipboardSharingLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
