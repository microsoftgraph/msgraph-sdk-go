package models
import (
    "errors"
    "strings"
)
// 
type OnlineMeetingVideoDisabledReason int

const (
    WATERMARKPROTECTION_ONLINEMEETINGVIDEODISABLEDREASON OnlineMeetingVideoDisabledReason = iota
    UNKNOWNFUTUREVALUE_ONLINEMEETINGVIDEODISABLEDREASON
)

func (i OnlineMeetingVideoDisabledReason) String() string {
    var values []string
    for p := OnlineMeetingVideoDisabledReason(1); p <= UNKNOWNFUTUREVALUE_ONLINEMEETINGVIDEODISABLEDREASON; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"watermarkProtection", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseOnlineMeetingVideoDisabledReason(v string) (any, error) {
    var result OnlineMeetingVideoDisabledReason
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "watermarkProtection":
                result |= WATERMARKPROTECTION_ONLINEMEETINGVIDEODISABLEDREASON
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ONLINEMEETINGVIDEODISABLEDREASON
            default:
                return 0, errors.New("Unknown OnlineMeetingVideoDisabledReason value: " + v)
        }
    }
    return &result, nil
}
func SerializeOnlineMeetingVideoDisabledReason(values []OnlineMeetingVideoDisabledReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OnlineMeetingVideoDisabledReason) isMultiValue() bool {
    return true
}
