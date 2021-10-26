package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "ALL":
            return ALL_CONDITIONALACCESSCLIENTAPP, nil
        case "BROWSER":
            return BROWSER_CONDITIONALACCESSCLIENTAPP, nil
        case "MOBILEAPPSANDDESKTOPCLIENTS":
            return MOBILEAPPSANDDESKTOPCLIENTS_CONDITIONALACCESSCLIENTAPP, nil
        case "EXCHANGEACTIVESYNC":
            return EXCHANGEACTIVESYNC_CONDITIONALACCESSCLIENTAPP, nil
        case "EASSUPPORTED":
            return EASSUPPORTED_CONDITIONALACCESSCLIENTAPP, nil
        case "OTHER":
            return OTHER_CONDITIONALACCESSCLIENTAPP, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONDITIONALACCESSCLIENTAPP, nil
    }
    return 0, errors.New("Unknown ConditionalAccessClientApp value: " + v)
}
func SerializeConditionalAccessClientApp(values []ConditionalAccessClientApp) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
