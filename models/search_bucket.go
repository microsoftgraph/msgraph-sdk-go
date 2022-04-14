package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchBucket 
type SearchBucket struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A token containing the encoded filter to aggregate search matches by the specific key value. To use the filter, pass the token as part of the aggregationFilter property in a searchRequest object, in the format '{field}:/'{aggregationFilterToken}/''. See an example.
    aggregationFilterToken *string
    // The approximate number of search matches that share the same value specified in the key property. Note that this number is not the exact number of matches.
    count *int32
    // The discrete value of the field that an aggregation was computed on.
    key *string
}
// NewSearchBucket instantiates a new searchBucket and sets the default values.
func NewSearchBucket()(*SearchBucket) {
    m := &SearchBucket{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchBucketFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchBucketFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchBucket(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchBucket) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAggregationFilterToken gets the aggregationFilterToken property value. A token containing the encoded filter to aggregate search matches by the specific key value. To use the filter, pass the token as part of the aggregationFilter property in a searchRequest object, in the format '{field}:/'{aggregationFilterToken}/''. See an example.
func (m *SearchBucket) GetAggregationFilterToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.aggregationFilterToken
    }
}
// GetCount gets the count property value. The approximate number of search matches that share the same value specified in the key property. Note that this number is not the exact number of matches.
func (m *SearchBucket) GetCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.count
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchBucket) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aggregationFilterToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAggregationFilterToken(val)
        }
        return nil
    }
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    return res
}
// GetKey gets the key property value. The discrete value of the field that an aggregation was computed on.
func (m *SearchBucket) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// Serialize serializes information the current object
func (m *SearchBucket) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchBucket) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAggregationFilterToken sets the aggregationFilterToken property value. A token containing the encoded filter to aggregate search matches by the specific key value. To use the filter, pass the token as part of the aggregationFilter property in a searchRequest object, in the format '{field}:/'{aggregationFilterToken}/''. See an example.
func (m *SearchBucket) SetAggregationFilterToken(value *string)() {
    if m != nil {
        m.aggregationFilterToken = value
    }
}
// SetCount sets the count property value. The approximate number of search matches that share the same value specified in the key property. Note that this number is not the exact number of matches.
func (m *SearchBucket) SetCount(value *int32)() {
    if m != nil {
        m.count = value
    }
}
// SetKey sets the key property value. The discrete value of the field that an aggregation was computed on.
func (m *SearchBucket) SetKey(value *string)() {
    if m != nil {
        m.key = value
    }
}
