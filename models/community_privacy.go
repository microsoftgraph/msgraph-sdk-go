package models
type CommunityPrivacy int

const (
    PUBLIC_COMMUNITYPRIVACY CommunityPrivacy = iota
    PRIVATE_COMMUNITYPRIVACY
    UNKNOWNFUTUREVALUE_COMMUNITYPRIVACY
)

func (i CommunityPrivacy) String() string {
    return []string{"public", "private", "unknownFutureValue"}[i]
}
func ParseCommunityPrivacy(v string) (any, error) {
    result := PUBLIC_COMMUNITYPRIVACY
    switch v {
        case "public":
            result = PUBLIC_COMMUNITYPRIVACY
        case "private":
            result = PRIVATE_COMMUNITYPRIVACY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_COMMUNITYPRIVACY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCommunityPrivacy(values []CommunityPrivacy) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CommunityPrivacy) isMultiValue() bool {
    return false
}
