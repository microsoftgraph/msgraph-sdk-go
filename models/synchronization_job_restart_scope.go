package models
import (
    "errors"
)
// 
type SynchronizationJobRestartScope int

const (
    NONE_SYNCHRONIZATIONJOBRESTARTSCOPE SynchronizationJobRestartScope = iota
    CONNECTORDATASTORE_SYNCHRONIZATIONJOBRESTARTSCOPE
    ESCROWS_SYNCHRONIZATIONJOBRESTARTSCOPE
    WATERMARK_SYNCHRONIZATIONJOBRESTARTSCOPE
    QUARANTINESTATE_SYNCHRONIZATIONJOBRESTARTSCOPE
    FULL_SYNCHRONIZATIONJOBRESTARTSCOPE
    FORCEDELETES_SYNCHRONIZATIONJOBRESTARTSCOPE
)

func (i SynchronizationJobRestartScope) String() string {
    return []string{"None", "ConnectorDataStore", "Escrows", "Watermark", "QuarantineState", "Full", "ForceDeletes"}[i]
}
func ParseSynchronizationJobRestartScope(v string) (any, error) {
    result := NONE_SYNCHRONIZATIONJOBRESTARTSCOPE
    switch v {
        case "None":
            result = NONE_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "ConnectorDataStore":
            result = CONNECTORDATASTORE_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "Escrows":
            result = ESCROWS_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "Watermark":
            result = WATERMARK_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "QuarantineState":
            result = QUARANTINESTATE_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "Full":
            result = FULL_SYNCHRONIZATIONJOBRESTARTSCOPE
        case "ForceDeletes":
            result = FORCEDELETES_SYNCHRONIZATIONJOBRESTARTSCOPE
        default:
            return 0, errors.New("Unknown SynchronizationJobRestartScope value: " + v)
    }
    return &result, nil
}
func SerializeSynchronizationJobRestartScope(values []SynchronizationJobRestartScope) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
