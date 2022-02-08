package graph
import (
    "strings"
    "errors"
)
// 
type MobileAppPublishingState int

const (
    NOTPUBLISHED_MOBILEAPPPUBLISHINGSTATE MobileAppPublishingState = iota
    PROCESSING_MOBILEAPPPUBLISHINGSTATE
    PUBLISHED_MOBILEAPPPUBLISHINGSTATE
)

func (i MobileAppPublishingState) String() string {
    return []string{"NOTPUBLISHED", "PROCESSING", "PUBLISHED"}[i]
}
func ParseMobileAppPublishingState(v string) (interface{}, error) {
    result := NOTPUBLISHED_MOBILEAPPPUBLISHINGSTATE
    switch strings.ToUpper(v) {
        case "NOTPUBLISHED":
            result = NOTPUBLISHED_MOBILEAPPPUBLISHINGSTATE
        case "PROCESSING":
            result = PROCESSING_MOBILEAPPPUBLISHINGSTATE
        case "PUBLISHED":
            result = PUBLISHED_MOBILEAPPPUBLISHINGSTATE
        default:
            return 0, errors.New("Unknown MobileAppPublishingState value: " + v)
    }
    return &result, nil
}
func SerializeMobileAppPublishingState(values []MobileAppPublishingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
