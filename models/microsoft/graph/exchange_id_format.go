package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "ENTRYID":
            return ENTRYID_EXCHANGEIDFORMAT, nil
        case "EWSID":
            return EWSID_EXCHANGEIDFORMAT, nil
        case "IMMUTABLEENTRYID":
            return IMMUTABLEENTRYID_EXCHANGEIDFORMAT, nil
        case "RESTID":
            return RESTID_EXCHANGEIDFORMAT, nil
        case "RESTIMMUTABLEENTRYID":
            return RESTIMMUTABLEENTRYID_EXCHANGEIDFORMAT, nil
    }
    return 0, errors.New("Unknown ExchangeIdFormat value: " + v)
}
func SerializeExchangeIdFormat(values []ExchangeIdFormat) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
