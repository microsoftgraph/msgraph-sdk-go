package graph
import (
    "strings"
    "errors"
)
// 
type BucketAggregationSortProperty int

const (
    COUNT_BUCKETAGGREGATIONSORTPROPERTY BucketAggregationSortProperty = iota
    KEYASSTRING_BUCKETAGGREGATIONSORTPROPERTY
    KEYASNUMBER_BUCKETAGGREGATIONSORTPROPERTY
    UNKNOWNFUTUREVALUE_BUCKETAGGREGATIONSORTPROPERTY
)

func (i BucketAggregationSortProperty) String() string {
    return []string{"COUNT", "KEYASSTRING", "KEYASNUMBER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseBucketAggregationSortProperty(v string) (interface{}, error) {
    result := COUNT_BUCKETAGGREGATIONSORTPROPERTY
    switch strings.ToUpper(v) {
        case "COUNT":
            result = COUNT_BUCKETAGGREGATIONSORTPROPERTY
        case "KEYASSTRING":
            result = KEYASSTRING_BUCKETAGGREGATIONSORTPROPERTY
        case "KEYASNUMBER":
            result = KEYASNUMBER_BUCKETAGGREGATIONSORTPROPERTY
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_BUCKETAGGREGATIONSORTPROPERTY
        default:
            return 0, errors.New("Unknown BucketAggregationSortProperty value: " + v)
    }
    return &result, nil
}
func SerializeBucketAggregationSortProperty(values []BucketAggregationSortProperty) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
