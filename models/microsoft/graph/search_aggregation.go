package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// searchAggregation 
type SearchAggregation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Defines the actual buckets of the computed aggregation.
    buckets []SearchBucket;
    // Defines on which field the aggregation was computed on.
    field *string;
}
// NewSearchAggregation instantiates a new searchAggregation and sets the default values.
func NewSearchAggregation()(*SearchAggregation) {
    m := &SearchAggregation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchAggregation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBuckets gets the buckets property value. Defines the actual buckets of the computed aggregation.
func (m *SearchAggregation) GetBuckets()([]SearchBucket) {
    if m == nil {
        return nil
    } else {
        return m.buckets
    }
}
// GetField gets the field property value. Defines on which field the aggregation was computed on.
func (m *SearchAggregation) GetField()(*string) {
    if m == nil {
        return nil
    } else {
        return m.field
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchAggregation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["buckets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchBucket() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SearchBucket, len(val))
            for i, v := range val {
                res[i] = *(v.(*SearchBucket))
            }
            m.SetBuckets(res)
        }
        return nil
    }
    res["field"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetField(val)
        }
        return nil
    }
    return res
}
func (m *SearchAggregation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchAggregation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBuckets()))
        for i, v := range m.GetBuckets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("buckets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("field", m.GetField())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchAggregation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBuckets sets the buckets property value. Defines the actual buckets of the computed aggregation.
func (m *SearchAggregation) SetBuckets(value []SearchBucket)() {
    m.buckets = value
}
// SetField sets the field property value. Defines on which field the aggregation was computed on.
func (m *SearchAggregation) SetField(value *string)() {
    m.field = value
}
