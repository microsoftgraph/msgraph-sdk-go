package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AggregationOption struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    bucketDefinition *BucketAggregationDefinition;
    // Computes aggregation on the field while the field exists in current entity type. Required.
    field *string;
    // The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional.
    size *int32;
}
// Instantiates a new aggregationOption and sets the default values.
func NewAggregationOption()(*AggregationOption) {
    m := &AggregationOption{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AggregationOption) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the bucketDefinition property value. 
func (m *AggregationOption) GetBucketDefinition()(*BucketAggregationDefinition) {
    if m == nil {
        return nil
    } else {
        return m.bucketDefinition
    }
}
// Gets the field property value. Computes aggregation on the field while the field exists in current entity type. Required.
func (m *AggregationOption) GetField()(*string) {
    if m == nil {
        return nil
    } else {
        return m.field
    }
}
// Gets the size property value. The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional.
func (m *AggregationOption) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// The deserialization information for the current model
func (m *AggregationOption) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bucketDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBucketAggregationDefinition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBucketDefinition(val.(*BucketAggregationDefinition))
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
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
func (m *AggregationOption) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AggregationOption) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bucketDefinition", m.GetBucketDefinition())
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
        err := writer.WriteInt32Value("size", m.GetSize())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AggregationOption) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the bucketDefinition property value. 
// Parameters:
//  - value : Value to set for the bucketDefinition property.
func (m *AggregationOption) SetBucketDefinition(value *BucketAggregationDefinition)() {
    m.bucketDefinition = value
}
// Sets the field property value. Computes aggregation on the field while the field exists in current entity type. Required.
// Parameters:
//  - value : Value to set for the field property.
func (m *AggregationOption) SetField(value *string)() {
    m.field = value
}
// Sets the size property value. The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional.
// Parameters:
//  - value : Value to set for the size property.
func (m *AggregationOption) SetSize(value *int32)() {
    m.size = value
}
