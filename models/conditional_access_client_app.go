package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityContainer singleton.
type ConditionalAccessClientApp int

const (
    ALL_CONDITIONALACCESSCLIENTAPP ConditionalAccessClientApp = iota
    BROWSER_CONDITIONALACCESSCLIENTAPP
    MOBILEAPPSANDDESKTOPCLIENTS_CONDITIONALACCESSCLIENTAPP
    EXCHANGEACTIVESYNC_CONDITIONALACCESSCLIENTAPP
    EASSUPPORTED_CONDITIONALACCESSCLIENTAPP
    OTHER_CONDITIONALACCESSCLIENTAPP
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSCLIENTAPP
)

func (i ConditionalAccessClientApp) String() string {
    return []string{"ALL", "BROWSER", "MOBILEAPPSANDDESKTOPCLIENTS", "EXCHANGEACTIVESYNC", "EASSUPPORTED", "OTHER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConditionalAccessClientApp(v string) (interface{}, error) {
    result := ALL_CONDITIONALACCESSCLIENTAPP
    switch strings.ToUpper(v) {
        case "ALL":
            result = ALL_CONDITIONALACCESSCLIENTAPP
        case "BROWSER":
            result = BROWSER_CONDITIONALACCESSCLIENTAPP
        case "MOBILEAPPSANDDESKTOPCLIENTS":
            result = MOBILEAPPSANDDESKTOPCLIENTS_CONDITIONALACCESSCLIENTAPP
        case "EXCHANGEACTIVESYNC":
            result = EXCHANGEACTIVESYNC_CONDITIONALACCESSCLIENTAPP
        case "EASSUPPORTED":
            result = EASSUPPORTED_CONDITIONALACCESSCLIENTAPP
        case "OTHER":
            result = OTHER_CONDITIONALACCESSCLIENTAPP
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONDITIONALACCESSCLIENTAPP
        default:
            return 0, errors.New("Unknown ConditionalAccessClientApp value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessClientApp(values []ConditionalAccessClientApp) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
