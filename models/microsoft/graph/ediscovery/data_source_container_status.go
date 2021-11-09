package ediscovery
import (
    "strings"
    "errors"
)
// 
type DataSourceContainerStatus int

const (
    ACTIVE_DATASOURCECONTAINERSTATUS DataSourceContainerStatus = iota
    RELEASED_DATASOURCECONTAINERSTATUS
    UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS
)

func (i DataSourceContainerStatus) String() string {
    return []string{"ACTIVE", "RELEASED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDataSourceContainerStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ACTIVE":
            return ACTIVE_DATASOURCECONTAINERSTATUS, nil
        case "RELEASED":
            return RELEASED_DATASOURCECONTAINERSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DATASOURCECONTAINERSTATUS, nil
    }
    return 0, errors.New("Unknown DataSourceContainerStatus value: " + v)
}
func SerializeDataSourceContainerStatus(values []DataSourceContainerStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
