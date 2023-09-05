package models
import (
    "errors"
    "strings"
)
// 
type TemplateScenarios int

const (
    NEW_TEMPLATESCENARIOS TemplateScenarios = iota
    SECUREFOUNDATION_TEMPLATESCENARIOS
    ZEROTRUST_TEMPLATESCENARIOS
    REMOTEWORK_TEMPLATESCENARIOS
    PROTECTADMINS_TEMPLATESCENARIOS
    EMERGINGTHREATS_TEMPLATESCENARIOS
    UNKNOWNFUTUREVALUE_TEMPLATESCENARIOS
)

func (i TemplateScenarios) String() string {
    var values []string
    for p := TemplateScenarios(1); p <= UNKNOWNFUTUREVALUE_TEMPLATESCENARIOS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"new", "secureFoundation", "zeroTrust", "remoteWork", "protectAdmins", "emergingThreats", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseTemplateScenarios(v string) (any, error) {
    var result TemplateScenarios
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "new":
                result |= NEW_TEMPLATESCENARIOS
            case "secureFoundation":
                result |= SECUREFOUNDATION_TEMPLATESCENARIOS
            case "zeroTrust":
                result |= ZEROTRUST_TEMPLATESCENARIOS
            case "remoteWork":
                result |= REMOTEWORK_TEMPLATESCENARIOS
            case "protectAdmins":
                result |= PROTECTADMINS_TEMPLATESCENARIOS
            case "emergingThreats":
                result |= EMERGINGTHREATS_TEMPLATESCENARIOS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_TEMPLATESCENARIOS
            default:
                return 0, errors.New("Unknown TemplateScenarios value: " + v)
        }
    }
    return &result, nil
}
func SerializeTemplateScenarios(values []TemplateScenarios) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TemplateScenarios) isMultiValue() bool {
    return true
}
