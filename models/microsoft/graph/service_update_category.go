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
    switch strings.ToUpper(v) {
        case "PREVENTORFIXISSUE":
            return PREVENTORFIXISSUE_SERVICEUPDATECATEGORY, nil
        case "PLANFORCHANGE":
            return PLANFORCHANGE_SERVICEUPDATECATEGORY, nil
        case "STAYINFORMED":
            return STAYINFORMED_SERVICEUPDATECATEGORY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SERVICEUPDATECATEGORY, nil
    }
    return 0, errors.New("Unknown ServiceUpdateCategory value: " + v)
}
func SerializeServiceUpdateCategory(values []ServiceUpdateCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
