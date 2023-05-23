package models
import (
    "errors"
)
// 
type SynchronizationStatusCode int

const (
    NOTCONFIGURED_SYNCHRONIZATIONSTATUSCODE SynchronizationStatusCode = iota
    NOTRUN_SYNCHRONIZATIONSTATUSCODE
    ACTIVE_SYNCHRONIZATIONSTATUSCODE
    PAUSED_SYNCHRONIZATIONSTATUSCODE
    QUARANTINE_SYNCHRONIZATIONSTATUSCODE
)

func (i SynchronizationStatusCode) String() string {
    return []string{"NotConfigured", "NotRun", "Active", "Paused", "Quarantine"}[i]
}
func ParseSynchronizationStatusCode(v string) (any, error) {
    result := NOTCONFIGURED_SYNCHRONIZATIONSTATUSCODE
    switch v {
        case "NotConfigured":
            result = NOTCONFIGURED_SYNCHRONIZATIONSTATUSCODE
        case "NotRun":
            result = NOTRUN_SYNCHRONIZATIONSTATUSCODE
        case "Active":
            result = ACTIVE_SYNCHRONIZATIONSTATUSCODE
        case "Paused":
            result = PAUSED_SYNCHRONIZATIONSTATUSCODE
        case "Quarantine":
            result = QUARANTINE_SYNCHRONIZATIONSTATUSCODE
        default:
            return 0, errors.New("Unknown SynchronizationStatusCode value: " + v)
    }
    return &result, nil
}
func SerializeSynchronizationStatusCode(values []SynchronizationStatusCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
