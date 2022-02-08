package graph
import (
    "strings"
    "errors"
)
// 
type ServiceUpdateCategory int

const (
    PREVENTORFIXISSUE_SERVICEUPDATECATEGORY ServiceUpdateCategory = iota
    PLANFORCHANGE_SERVICEUPDATECATEGORY
    STAYINFORMED_SERVICEUPDATECATEGORY
    UNKNOWNFUTUREVALUE_SERVICEUPDATECATEGORY
)

func (i ServiceUpdateCategory) String() string {
    return []string{"PREVENTORFIXISSUE", "PLANFORCHANGE", "STAYINFORMED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseServiceUpdateCategory(v string) (interface{}, error) {
    result := PREVENTORFIXISSUE_SERVICEUPDATECATEGORY
    switch strings.ToUpper(v) {
        case "PREVENTORFIXISSUE":
            result = PREVENTORFIXISSUE_SERVICEUPDATECATEGORY
        case "PLANFORCHANGE":
            result = PLANFORCHANGE_SERVICEUPDATECATEGORY
        case "STAYINFORMED":
            result = STAYINFORMED_SERVICEUPDATECATEGORY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SERVICEUPDATECATEGORY
        default:
            return 0, errors.New("Unknown ServiceUpdateCategory value: " + v)
    }
    return &result, nil
}
func SerializeServiceUpdateCategory(values []ServiceUpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
