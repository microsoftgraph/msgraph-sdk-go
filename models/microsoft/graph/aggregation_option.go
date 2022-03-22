package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AggregationOption 
type AggregationOption struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    bucketDefinition BucketAggregationDefinitionable;
    // Computes aggregation on the field while the field exists in current entity type. Required.
    field *string;
    // The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional.
    size *int32;
}
// NewAggregationOption instantiates a new aggregationOption and sets the default values.
func NewAggregationOption()(*AggregationOption) {
    m := &AggregationOption{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAggregationOptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAggregationOptionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAggregationOption(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AggregationOption) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBucketDefinition gets the bucketDefinition property value. 
func (m *AggregationOption) GetBucketDefinition()(BucketAggregationDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.bucketDefinition
    }
}
// GetField gets the field property value. Computes aggregation on the field while the field exists in current entity type. Required.
func (m *AggregationOption) GetField()(*string) {
    if m == nil {
        return nil
    } else {
        return m.field
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AggregationOption) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bucketDefinition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateBucketAggregationDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBucketDefinition(val.(BucketAggregationDefinitionable))
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
// GetSize gets the size property value. The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional.
func (m *AggregationOption) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AggregationOption) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBucketDefinition sets the bucketDefinition property value. 
func (m *AggregationOption) SetBucketDefinition(value BucketAggregationDefinitionable)() {
    if m != nil {
        m.bucketDefinition = value
    }
}
// SetField sets the field property value. Computes aggregation on the field while the field exists in current entity type. Required.
func (m *AggregationOption) SetField(value *string)() {
    if m != nil {
        m.field = value
    }
}
// SetSize sets the size property value. The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional.
func (m *AggregationOption) SetSize(value *int32)() {
    if m != nil {
        m.size = value
    }
}
