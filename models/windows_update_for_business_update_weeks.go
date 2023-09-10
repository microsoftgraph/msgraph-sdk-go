package models
import (
    "errors"
    "strings"
)
// Scheduled the update installation on the weeks of the month
type WindowsUpdateForBusinessUpdateWeeks int

const (
    // Allow the user to set.
    USERDEFINED_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS WindowsUpdateForBusinessUpdateWeeks = iota
    // Scheduled the update installation on the first week of the month
    FIRSTWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
    // Scheduled the update installation on the second week of the month
    SECONDWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
    // Scheduled the update installation on the third week of the month
    THIRDWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
    // Scheduled the update installation on the fourth week of the month
    FOURTHWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
    // Scheduled the update installation on every week of the month
    EVERYWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
    // Evolvable enum member
    UNKNOWNFUTUREVALUE_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
)

func (i WindowsUpdateForBusinessUpdateWeeks) String() string {
    var values []string
    for p := WindowsUpdateForBusinessUpdateWeeks(1); p <= UNKNOWNFUTUREVALUE_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"userDefined", "firstWeek", "secondWeek", "thirdWeek", "fourthWeek", "everyWeek", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseWindowsUpdateForBusinessUpdateWeeks(v string) (any, error) {
    var result WindowsUpdateForBusinessUpdateWeeks
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "userDefined":
                result |= USERDEFINED_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
            case "firstWeek":
                result |= FIRSTWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
            case "secondWeek":
                result |= SECONDWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
            case "thirdWeek":
                result |= THIRDWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
            case "fourthWeek":
                result |= FOURTHWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
            case "everyWeek":
                result |= EVERYWEEK_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_WINDOWSUPDATEFORBUSINESSUPDATEWEEKS
            default:
                return 0, errors.New("Unknown WindowsUpdateForBusinessUpdateWeeks value: " + v)
        }
    }
    return &result, nil
}
func SerializeWindowsUpdateForBusinessUpdateWeeks(values []WindowsUpdateForBusinessUpdateWeeks) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsUpdateForBusinessUpdateWeeks) isMultiValue() bool {
    return true
}
