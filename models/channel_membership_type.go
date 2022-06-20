package models
import (
    "errors"
)
// Provides operations to manage the collection of application entities.
type ChannelMembershipType int

const (
    STANDARD_CHANNELMEMBERSHIPTYPE ChannelMembershipType = iota
    PRIVATE_CHANNELMEMBERSHIPTYPE
    UNKNOWNFUTUREVALUE_CHANNELMEMBERSHIPTYPE
)

func (i ChannelMembershipType) String() string {
    return []string{"standard", "private", "unknownFutureValue"}[i]
}
func ParseChannelMembershipType(v string) (interface{}, error) {
    result := STANDARD_CHANNELMEMBERSHIPTYPE
    switch v {
        case "standard":
            result = STANDARD_CHANNELMEMBERSHIPTYPE
        case "private":
            result = PRIVATE_CHANNELMEMBERSHIPTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CHANNELMEMBERSHIPTYPE
        default:
            return 0, errors.New("Unknown ChannelMembershipType value: " + v)
    }
    return &result, nil
}
func SerializeChannelMembershipType(values []ChannelMembershipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
