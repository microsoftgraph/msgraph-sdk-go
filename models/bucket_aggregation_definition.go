package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BucketAggregationDefinition 
type BucketAggregationDefinition struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // True to specify the sort order as descending. The default is false, with the sort order as ascending. Optional.
    isDescending *bool
    // The minimum number of items that should be present in the aggregation to be returned in a bucket. Optional.
    minimumCount *int32
    // A filter to define a matching criteria. The key should start with the specified prefix to be returned in the response. Optional.
    prefixFilter *string
    // Specifies the manual ranges to compute the aggregations. This is only valid for non-string refiners of date or numeric type. Optional.
    ranges []BucketAggregationRangeable
    // The sortBy property
    sortBy *BucketAggregationSortProperty
}
// NewBucketAggregationDefinition instantiates a new bucketAggregationDefinition and sets the default values.
func NewBucketAggregationDefinition()(*BucketAggregationDefinition) {
    m := &BucketAggregationDefinition{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBucketAggregationDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBucketAggregationDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBucketAggregationDefinition(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BucketAggregationDefinition) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BucketAggregationDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isDescending"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDescending(val)
        }
        return nil
    }
    res["minimumCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumCount(val)
        }
        return nil
    }
    res["prefixFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrefixFilter(val)
        }
        return nil
    }
    res["ranges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBucketAggregationRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BucketAggregationRangeable, len(val))
            for i, v := range val {
                res[i] = v.(BucketAggregationRangeable)
            }
            m.SetRanges(res)
        }
        return nil
    }
    res["sortBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBucketAggregationSortProperty)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSortBy(val.(*BucketAggregationSortProperty))
        }
        return nil
    }
    return res
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
func (m *BucketAggregationDefinition) GetRanges()([]BucketAggregationRangeable) {
    if m == nil {
        return nil
    } else {
        return m.ranges
    }
}
// GetSortBy gets the sortBy property value. The sortBy property
func (m *BucketAggregationDefinition) GetSortBy()(*BucketAggregationSortProperty) {
    if m == nil {
        return nil
    } else {
        return m.sortBy
    }
}
// Serialize serializes information the current object
func (m *BucketAggregationDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRanges()))
        for i, v := range m.GetRanges() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("ranges", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSortBy() != nil {
        cast := (*m.GetSortBy()).String()
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
func (m *BucketAggregationDefinition) SetRanges(value []BucketAggregationRangeable)() {
    if m != nil {
        m.ranges = value
    }
}
// SetSortBy sets the sortBy property value. The sortBy property
func (m *BucketAggregationDefinition) SetSortBy(value *BucketAggregationSortProperty)() {
    if m != nil {
        m.sortBy = value
    }
}
