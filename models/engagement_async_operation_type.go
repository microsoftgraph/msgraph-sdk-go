package models
type EngagementAsyncOperationType int

const (
    CREATECOMMUNITY_ENGAGEMENTASYNCOPERATIONTYPE EngagementAsyncOperationType = iota
    UNKNOWNFUTUREVALUE_ENGAGEMENTASYNCOPERATIONTYPE
)

func (i EngagementAsyncOperationType) String() string {
    return []string{"createCommunity", "unknownFutureValue"}[i]
}
func ParseEngagementAsyncOperationType(v string) (any, error) {
    result := CREATECOMMUNITY_ENGAGEMENTASYNCOPERATIONTYPE
    switch v {
        case "createCommunity":
            result = CREATECOMMUNITY_ENGAGEMENTASYNCOPERATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENGAGEMENTASYNCOPERATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEngagementAsyncOperationType(values []EngagementAsyncOperationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EngagementAsyncOperationType) isMultiValue() bool {
    return false
}
