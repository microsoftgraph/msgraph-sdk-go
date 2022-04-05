package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of drive entities.
type ChannelMembershipType int

const (
    STANDARD_CHANNELMEMBERSHIPTYPE ChannelMembershipType = iota
    PRIVATE_CHANNELMEMBERSHIPTYPE
    UNKNOWNFUTUREVALUE_CHANNELMEMBERSHIPTYPE
)

func (i ChannelMembershipType) String() string {
    return []string{"STANDARD", "PRIVATE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseChannelMembershipType(v string) (interface{}, error) {
    result := STANDARD_CHANNELMEMBERSHIPTYPE
    switch strings.ToUpper(v) {
        case "STANDARD":
            result = STANDARD_CHANNELMEMBERSHIPTYPE
        case "PRIVATE":
            result = PRIVATE_CHANNELMEMBERSHIPTYPE
        case "UNKNOWNFUTUREVALUE":
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
