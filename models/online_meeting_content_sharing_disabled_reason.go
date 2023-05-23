package models
import (
    "errors"
)
// 
type OnlineMeetingContentSharingDisabledReason int

const (
    WATERMARKPROTECTION_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON OnlineMeetingContentSharingDisabledReason = iota
    UNKNOWNFUTUREVALUE_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON
)

func (i OnlineMeetingContentSharingDisabledReason) String() string {
    return []string{"watermarkProtection", "unknownFutureValue"}[i]
}
func ParseOnlineMeetingContentSharingDisabledReason(v string) (any, error) {
    result := WATERMARKPROTECTION_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON
    switch v {
        case "watermarkProtection":
            result = WATERMARKPROTECTION_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ONLINEMEETINGCONTENTSHARINGDISABLEDREASON
        default:
            return 0, errors.New("Unknown OnlineMeetingContentSharingDisabledReason value: " + v)
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
