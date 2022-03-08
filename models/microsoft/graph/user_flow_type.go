package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityContainer singleton.
type UserFlowType int

const (
    SIGNUP_USERFLOWTYPE UserFlowType = iota
    SIGNIN_USERFLOWTYPE
    SIGNUPORSIGNIN_USERFLOWTYPE
    PASSWORDRESET_USERFLOWTYPE
    PROFILEUPDATE_USERFLOWTYPE
    RESOURCEOWNER_USERFLOWTYPE
    UNKNOWNFUTUREVALUE_USERFLOWTYPE
)

func (i UserFlowType) String() string {
    return []string{"SIGNUP", "SIGNIN", "SIGNUPORSIGNIN", "PASSWORDRESET", "PROFILEUPDATE", "RESOURCEOWNER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUserFlowType(v string) (interface{}, error) {
    result := SIGNUP_USERFLOWTYPE
    switch strings.ToUpper(v) {
        case "SIGNUP":
            result = SIGNUP_USERFLOWTYPE
        case "SIGNIN":
            result = SIGNIN_USERFLOWTYPE
        case "SIGNUPORSIGNIN":
            result = SIGNUPORSIGNIN_USERFLOWTYPE
        case "PASSWORDRESET":
            result = PASSWORDRESET_USERFLOWTYPE
        case "PROFILEUPDATE":
            result = PROFILEUPDATE_USERFLOWTYPE
        case "RESOURCEOWNER":
            result = RESOURCEOWNER_USERFLOWTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_USERFLOWTYPE
        default:
            return 0, errors.New("Unknown UserFlowType value: " + v)
    }
    return &result, nil
}
func SerializeUserFlowType(values []UserFlowType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
