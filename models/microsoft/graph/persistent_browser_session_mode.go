package graph
import (
    "strings"
    "errors"
)
type PersistentBrowserSessionMode int

const (
    ALWAYS_PERSISTENTBROWSERSESSIONMODE PersistentBrowserSessionMode = iota
    NEVER_PERSISTENTBROWSERSESSIONMODE
)

func (i PersistentBrowserSessionMode) String() string {
    return []string{"ALWAYS", "NEVER"}[i]
}
func ParsePersistentBrowserSessionMode(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ALWAYS":
            return ALWAYS_PERSISTENTBROWSERSESSIONMODE, nil
        case "NEVER":
            return NEVER_PERSISTENTBROWSERSESSIONMODE, nil
    }
    return 0, errors.New("Unknown PersistentBrowserSessionMode value: " + v)
}
func SerializePersistentBrowserSessionMode(values []PersistentBrowserSessionMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
