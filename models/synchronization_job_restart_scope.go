package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := SynchronizationJobRestartScope(1); p <= FORCEDELETES_SYNCHRONIZATIONJOBRESTARTSCOPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"None", "ConnectorDataStore", "Escrows", "Watermark", "QuarantineState", "Full", "ForceDeletes"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseSynchronizationJobRestartScope(v string) (any, error) {
    var result SynchronizationJobRestartScope
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "None":
                result |= NONE_SYNCHRONIZATIONJOBRESTARTSCOPE
            case "ConnectorDataStore":
                result |= CONNECTORDATASTORE_SYNCHRONIZATIONJOBRESTARTSCOPE
            case "Escrows":
                result |= ESCROWS_SYNCHRONIZATIONJOBRESTARTSCOPE
            case "Watermark":
                result |= WATERMARK_SYNCHRONIZATIONJOBRESTARTSCOPE
            case "QuarantineState":
                result |= QUARANTINESTATE_SYNCHRONIZATIONJOBRESTARTSCOPE
            case "Full":
                result |= FULL_SYNCHRONIZATIONJOBRESTARTSCOPE
            case "ForceDeletes":
                result |= FORCEDELETES_SYNCHRONIZATIONJOBRESTARTSCOPE
            default:
                return 0, errors.New("Unknown SynchronizationJobRestartScope value: " + v)
        }
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
func (i SynchronizationJobRestartScope) isMultiValue() bool {
    return true
}
