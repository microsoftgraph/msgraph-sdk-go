package graph
import (
    "strings"
    "errors"
)
// 
type ServiceHealthOrigin int

const (
    MICROSOFT_SERVICEHEALTHORIGIN ServiceHealthOrigin = iota
    THIRDPARTY_SERVICEHEALTHORIGIN
    CUSTOMER_SERVICEHEALTHORIGIN
    UNKNOWNFUTUREVALUE_SERVICEHEALTHORIGIN
)

func (i ServiceHealthOrigin) String() string {
    return []string{"MICROSOFT", "THIRDPARTY", "CUSTOMER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseServiceHealthOrigin(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "MICROSOFT":
            return MICROSOFT_SERVICEHEALTHORIGIN, nil
        case "THIRDPARTY":
            return THIRDPARTY_SERVICEHEALTHORIGIN, nil
        case "CUSTOMER":
            return CUSTOMER_SERVICEHEALTHORIGIN, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SERVICEHEALTHORIGIN, nil
    }
    return 0, errors.New("Unknown ServiceHealthOrigin value: " + v)
}
func SerializeServiceHealthOrigin(values []ServiceHealthOrigin) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
