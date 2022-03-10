package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the translateExchangeIds method.
type ExchangeIdFormat int

const (
    ENTRYID_EXCHANGEIDFORMAT ExchangeIdFormat = iota
    EWSID_EXCHANGEIDFORMAT
    IMMUTABLEENTRYID_EXCHANGEIDFORMAT
    RESTID_EXCHANGEIDFORMAT
    RESTIMMUTABLEENTRYID_EXCHANGEIDFORMAT
)

func (i ExchangeIdFormat) String() string {
    return []string{"ENTRYID", "EWSID", "IMMUTABLEENTRYID", "RESTID", "RESTIMMUTABLEENTRYID"}[i]
}
func ParseExchangeIdFormat(v string) (interface{}, error) {
    result := ENTRYID_EXCHANGEIDFORMAT
    switch strings.ToUpper(v) {
        case "ENTRYID":
            result = ENTRYID_EXCHANGEIDFORMAT
        case "EWSID":
            result = EWSID_EXCHANGEIDFORMAT
        case "IMMUTABLEENTRYID":
            result = IMMUTABLEENTRYID_EXCHANGEIDFORMAT
        case "RESTID":
            result = RESTID_EXCHANGEIDFORMAT
        case "RESTIMMUTABLEENTRYID":
            result = RESTIMMUTABLEENTRYID_EXCHANGEIDFORMAT
        default:
            return 0, errors.New("Unknown ExchangeIdFormat value: " + v)
    }
    return &result, nil
}
func SerializeExchangeIdFormat(values []ExchangeIdFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
