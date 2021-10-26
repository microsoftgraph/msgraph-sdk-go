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
    switch strings.ToUpper(v) {
        case "NOTPUBLISHED":
            return NOTPUBLISHED_MOBILEAPPPUBLISHINGSTATE, nil
        case "PROCESSING":
            return PROCESSING_MOBILEAPPPUBLISHINGSTATE, nil
        case "PUBLISHED":
            return PUBLISHED_MOBILEAPPPUBLISHINGSTATE, nil
    }
    return 0, errors.New("Unknown MobileAppPublishingState value: " + v)
}
func SerializeMobileAppPublishingState(values []MobileAppPublishingState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
