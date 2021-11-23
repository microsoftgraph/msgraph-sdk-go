package percentrank_exc

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// PercentRank_ExcRequestBody 
type PercentRank_ExcRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    array *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json;
    // 
    significance *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json;
    // 
    x *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json;
}
// NewPercentRank_ExcRequestBody instantiates a new percentRank_ExcRequestBody and sets the default values.
func NewPercentRank_ExcRequestBody()(*PercentRank_ExcRequestBody) {
    m := &PercentRank_ExcRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PercentRank_ExcRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetArray gets the array property value. 
func (m *PercentRank_ExcRequestBody) GetArray()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json) {
    if m == nil {
        return nil
    } else {
        return m.array
    }
}
// GetSignificance gets the significance property value. 
func (m *PercentRank_ExcRequestBody) GetSignificance()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json) {
    if m == nil {
        return nil
    } else {
        return m.significance
    }
}
// GetX gets the x property value. 
func (m *PercentRank_ExcRequestBody) GetX()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json) {
    if m == nil {
        return nil
    } else {
        return m.x
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PercentRank_ExcRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["array"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArray(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json))
        }
        return nil
    }
    res["significance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignificance(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json))
        }
        return nil
    }
    res["x"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetX(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json))
        }
        return nil
    }
    return res
}
func (m *PercentRank_ExcRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PercentRank_ExcRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("array", m.GetArray())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("significance", m.GetSignificance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("x", m.GetX())
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
func (m *PercentRank_ExcRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetArray sets the array property value. 
func (m *PercentRank_ExcRequestBody) SetArray(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json)() {
    m.array = value
}
// SetSignificance sets the significance property value. 
func (m *PercentRank_ExcRequestBody) SetSignificance(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json)() {
    m.significance = value
}
// SetX sets the x property value. 
func (m *PercentRank_ExcRequestBody) SetX(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Json)() {
    m.x = value
}
