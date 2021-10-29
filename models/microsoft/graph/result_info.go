package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ResultInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The result code.
    code *int32;
    // The message.
    message *string;
    // The result sub-code.
    subcode *int32;
}
// Instantiates a new resultInfo and sets the default values.
func NewResultInfo()(*ResultInfo) {
    m := &ResultInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResultInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the code property value. The result code.
func (m *ResultInfo) GetCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// Gets the message property value. The message.
func (m *ResultInfo) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// Gets the subcode property value. The result sub-code.
func (m *ResultInfo) GetSubcode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.subcode
    }
}
// The deserialization information for the current model
func (m *ResultInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetCode(val)
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["subcode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSubcode(val)
        return nil
    }
    return res
}
func (m *ResultInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ResultInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("subcode", m.GetSubcode())
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
func (m *ResultInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the code property value. The result code.
// Parameters:
//  - value : Value to set for the code property.
func (m *ResultInfo) SetCode(value *int32)() {
    m.code = value
}
// Sets the message property value. The message.
// Parameters:
//  - value : Value to set for the message property.
func (m *ResultInfo) SetMessage(value *string)() {
    m.message = value
}
// Sets the subcode property value. The result sub-code.
// Parameters:
//  - value : Value to set for the subcode property.
func (m *ResultInfo) SetSubcode(value *int32)() {
    m.subcode = value
}
