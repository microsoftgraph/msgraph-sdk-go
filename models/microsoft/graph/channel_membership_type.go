package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "STANDARD":
            return STANDARD_CHANNELMEMBERSHIPTYPE, nil
        case "PRIVATE":
            return PRIVATE_CHANNELMEMBERSHIPTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CHANNELMEMBERSHIPTYPE, nil
    }
    return 0, errors.New("Unknown ChannelMembershipType value: " + v)
}
func SerializeChannelMembershipType(values []ChannelMembershipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
