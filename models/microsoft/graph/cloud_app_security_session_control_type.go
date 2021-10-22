package graph
import (
    "strings"
    "errors"
)
type CloudAppSecuritySessionControlType int

const (
    MCASCONFIGURED_CLOUDAPPSECURITYSESSIONCONTROLTYPE CloudAppSecuritySessionControlType = iota
    MONITORONLY_CLOUDAPPSECURITYSESSIONCONTROLTYPE
    BLOCKDOWNLOADS_CLOUDAPPSECURITYSESSIONCONTROLTYPE
    UNKNOWNFUTUREVALUE_CLOUDAPPSECURITYSESSIONCONTROLTYPE
)

func (i CloudAppSecuritySessionControlType) String() string {
    return []string{"MCASCONFIGURED", "MONITORONLY", "BLOCKDOWNLOADS", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudAppSecuritySessionControlType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "MCASCONFIGURED":
            return MCASCONFIGURED_CLOUDAPPSECURITYSESSIONCONTROLTYPE, nil
        case "MONITORONLY":
            return MONITORONLY_CLOUDAPPSECURITYSESSIONCONTROLTYPE, nil
        case "BLOCKDOWNLOADS":
            return BLOCKDOWNLOADS_CLOUDAPPSECURITYSESSIONCONTROLTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDAPPSECURITYSESSIONCONTROLTYPE, nil
    }
    return 0, errors.New("Unknown CloudAppSecuritySessionControlType value: " + v)
}
func SerializeCloudAppSecuritySessionControlType(values []CloudAppSecuritySessionControlType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
