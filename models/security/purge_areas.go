package security
import (
    "errors"
    "strings"
)
// 
type PurgeAreas int

const (
    MAILBOXES_PURGEAREAS PurgeAreas = iota
    TEAMSMESSAGES_PURGEAREAS
    UNKNOWNFUTUREVALUE_PURGEAREAS
)

func (i PurgeAreas) String() string {
    var values []string
    for p := PurgeAreas(1); p <= UNKNOWNFUTUREVALUE_PURGEAREAS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"mailboxes", "teamsMessages", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParsePurgeAreas(v string) (any, error) {
    var result PurgeAreas
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "mailboxes":
                result |= MAILBOXES_PURGEAREAS
            case "teamsMessages":
                result |= TEAMSMESSAGES_PURGEAREAS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_PURGEAREAS
            default:
                return 0, errors.New("Unknown PurgeAreas value: " + v)
        }
    }
    return &result, nil
}
func SerializePurgeAreas(values []PurgeAreas) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PurgeAreas) isMultiValue() bool {
    return true
}
