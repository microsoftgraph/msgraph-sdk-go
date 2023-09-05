package models
import (
    "errors"
    "strings"
)
// 
type OnlineMeetingContentSharingDisabledReason int

const (
    WATERMARKPROTECTION_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON OnlineMeetingContentSharingDisabledReason = iota
    UNKNOWNFUTUREVALUE_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON
)

func (i OnlineMeetingContentSharingDisabledReason) String() string {
    var values []string
    for p := OnlineMeetingContentSharingDisabledReason(1); p <= UNKNOWNFUTUREVALUE_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"watermarkProtection", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseOnlineMeetingContentSharingDisabledReason(v string) (any, error) {
    var result OnlineMeetingContentSharingDisabledReason
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "watermarkProtection":
                result |= WATERMARKPROTECTION_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON
            default:
                return 0, errors.New("Unknown OnlineMeetingContentSharingDisabledReason value: " + v)
        }
    }
    return &result, nil
}
func SerializeOnlineMeetingContentSharingDisabledReason(values []OnlineMeetingContentSharingDisabledReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OnlineMeetingContentSharingDisabledReason) isMultiValue() bool {
    return true
}
