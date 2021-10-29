package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SearchBucket struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A token containing the encoded filter to aggregate search matches by the specific key value. To use the filter, pass the token as part of the aggregationFilter property in a searchRequest object, in the format '{field}:/'{aggregationFilterToken}/''. See an example.
    aggregationFilterToken *string;
    // The approximate number of search matches that share the same value specified in the key property. Note that this number is not the exact number of matches.
    count *int32;
    // The discrete value of the field that an aggregation was computed on.
    key *string;
}
// Instantiates a new searchBucket and sets the default values.
func NewSearchBucket()(*SearchBucket) {
    m := &SearchBucket{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchBucket) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the aggregationFilterToken property value. A token containing the encoded filter to aggregate search matches by the specific key value. To use the filter, pass the token as part of the aggregationFilter property in a searchRequest object, in the format '{field}:/'{aggregationFilterToken}/''. See an example.
func (m *SearchBucket) GetAggregationFilterToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aggregationFilterToken
    }
}
// Gets the count property value. The approximate number of search matches that share the same value specified in the key property. Note that this number is not the exact number of matches.
func (m *SearchBucket) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
// Gets the key property value. The discrete value of the field that an aggregation was computed on.
func (m *SearchBucket) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// The deserialization information for the current model
func (m *SearchBucket) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["aggregationFilterToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAggregationFilterToken(val)
        return nil
    }
    res["count"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCount(val)
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    return res
}
func (m *SearchBucket) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SearchBucket) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("aggregationFilterToken", m.GetAggregationFilterToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key", m.GetKey())
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
func (m *SearchBucket) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the aggregationFilterToken property value. A token containing the encoded filter to aggregate search matches by the specific key value. To use the filter, pass the token as part of the aggregationFilter property in a searchRequest object, in the format '{field}:/'{aggregationFilterToken}/''. See an example.
// Parameters:
//  - value : Value to set for the aggregationFilterToken property.
func (m *SearchBucket) SetAggregationFilterToken(value *string)() {
    m.aggregationFilterToken = value
}
// Sets the count property value. The approximate number of search matches that share the same value specified in the key property. Note that this number is not the exact number of matches.
// Parameters:
//  - value : Value to set for the count property.
func (m *SearchBucket) SetCount(value *int32)() {
    m.count = value
}
// Sets the key property value. The discrete value of the field that an aggregation was computed on.
// Parameters:
//  - value : Value to set for the key property.
func (m *SearchBucket) SetKey(value *string)() {
    m.key = value
}
