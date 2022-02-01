package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BucketAggregationDefinition 
type BucketAggregationDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // True to specify the sort order as descending. The default is false, with the sort order as ascending. Optional.
    isDescending *bool;
    // The minimum number of items that should be present in the aggregation to be returned in a bucket. Optional.
    minimumCount *int32;
    // A filter to define a matching criteria. The key should start with the specified prefix to be returned in the response. Optional.
    prefixFilter *string;
    // Specifies the manual ranges to compute the aggregations. This is only valid for non-string refiners of date or numeric type. Optional.
    ranges []BucketAggregationRange;
    // The possible values are count to sort by the number of matches in the aggregation, keyAsStringto sort alphabeticaly based on the key in the aggregation, keyAsNumber for numerical sorting based on the key in the aggregation. Required.
    sortBy *BucketAggregationSortProperty;
}
// NewBucketAggregationDefinition instantiates a new bucketAggregationDefinition and sets the default values.
func NewBucketAggregationDefinition()(*BucketAggregationDefinition) {
    m := &BucketAggregationDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BucketAggregationDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsDescending gets the isDescending property value. True to specify the sort order as descending. The default is false, with the sort order as ascending. Optional.
func (m *BucketAggregationDefinition) GetIsDescending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDescending
    }
}
// GetMinimumCount gets the minimumCount property value. The minimum number of items that should be present in the aggregation to be returned in a bucket. Optional.
func (m *BucketAggregationDefinition) GetMinimumCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumCount
    }
}
// GetPrefixFilter gets the prefixFilter property value. A filter to define a matching criteria. The key should start with the specified prefix to be returned in the response. Optional.
func (m *BucketAggregationDefinition) GetPrefixFilter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.prefixFilter
    }
}
// GetRanges gets the ranges property value. Specifies the manual ranges to compute the aggregations. This is only valid for non-string refiners of date or numeric type. Optional.
func (m *BucketAggregationDefinition) GetRanges()([]BucketAggregationRange) {
    if m == nil {
        return nil
    } else {
        return m.ranges
    }
}
// GetSortBy gets the sortBy property value. The possible values are count to sort by the number of matches in the aggregation, keyAsStringto sort alphabeticaly based on the key in the aggregation, keyAsNumber for numerical sorting based on the key in the aggregation. Required.
func (m *BucketAggregationDefinition) GetSortBy()(*BucketAggregationSortProperty) {
    if m == nil {
        return nil
    } else {
        return m.sortBy
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BucketAggregationDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDescending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDescending(val)
        }
        return nil
    }
    res["minimumCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumCount(val)
        }
        return nil
    }
    res["prefixFilter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrefixFilter(val)
        }
        return nil
    }
    res["ranges"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBucketAggregationRange() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BucketAggregationRange, len(val))
            for i, v := range val {
                res[i] = *(v.(*BucketAggregationRange))
            }
            m.SetRanges(res)
        }
        return nil
    }
    res["sortBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBucketAggregationSortProperty)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(BucketAggregationSortProperty)
            m.SetSortBy(&cast)
        }
        return nil
    }
    return res
}
func (m *BucketAggregationDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BucketAggregationDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDescending", m.GetIsDescending())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minimumCount", m.GetMinimumCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("prefixFilter", m.GetPrefixFilter())
        if err != nil {
            return err
        }
    }
    if m.GetRanges() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRanges()))
        for i, v := range m.GetRanges() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("ranges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSortBy() != nil {
        cast := m.GetSortBy().String()
        err := writer.WriteStringValue("sortBy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BucketAggregationDefinition) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsDescending sets the isDescending property value. True to specify the sort order as descending. The default is false, with the sort order as ascending. Optional.
func (m *BucketAggregationDefinition) SetIsDescending(value *bool)() {
    if m != nil {
        m.isDescending = value
    }
}
// SetMinimumCount sets the minimumCount property value. The minimum number of items that should be present in the aggregation to be returned in a bucket. Optional.
func (m *BucketAggregationDefinition) SetMinimumCount(value *int32)() {
    if m != nil {
        m.minimumCount = value
    }
}
// SetPrefixFilter sets the prefixFilter property value. A filter to define a matching criteria. The key should start with the specified prefix to be returned in the response. Optional.
func (m *BucketAggregationDefinition) SetPrefixFilter(value *string)() {
    if m != nil {
        m.prefixFilter = value
    }
}
// SetRanges sets the ranges property value. Specifies the manual ranges to compute the aggregations. This is only valid for non-string refiners of date or numeric type. Optional.
func (m *BucketAggregationDefinition) SetRanges(value []BucketAggregationRange)() {
    if m != nil {
        m.ranges = value
    }
}
// SetSortBy sets the sortBy property value. The possible values are count to sort by the number of matches in the aggregation, keyAsStringto sort alphabeticaly based on the key in the aggregation, keyAsNumber for numerical sorting based on the key in the aggregation. Required.
func (m *BucketAggregationDefinition) SetSortBy(value *BucketAggregationSortProperty)() {
    if m != nil {
        m.sortBy = value
    }
}
