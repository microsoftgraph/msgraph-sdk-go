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
    switch strings.ToUpper(v) {
        case "COUNT":
            return COUNT_BUCKETAGGREGATIONSORTPROPERTY, nil
        case "KEYASSTRING":
            return KEYASSTRING_BUCKETAGGREGATIONSORTPROPERTY, nil
        case "KEYASNUMBER":
            return KEYASNUMBER_BUCKETAGGREGATIONSORTPROPERTY, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_BUCKETAGGREGATIONSORTPROPERTY, nil
    }
    return 0, errors.New("Unknown BucketAggregationSortProperty value: " + v)
}
func SerializeBucketAggregationSortProperty(values []BucketAggregationSortProperty) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
