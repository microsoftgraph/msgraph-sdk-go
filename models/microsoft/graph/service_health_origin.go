package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the admin singleton.
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
    result := MICROSOFT_SERVICEHEALTHORIGIN
    switch strings.ToUpper(v) {
        case "MICROSOFT":
            result = MICROSOFT_SERVICEHEALTHORIGIN
        case "THIRDPARTY":
            result = THIRDPARTY_SERVICEHEALTHORIGIN
        case "CUSTOMER":
            result = CUSTOMER_SERVICEHEALTHORIGIN
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SERVICEHEALTHORIGIN
        default:
            return 0, errors.New("Unknown ServiceHealthOrigin value: " + v)
    }
    return &result, nil
}
func SerializeServiceHealthOrigin(values []ServiceHealthOrigin) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
