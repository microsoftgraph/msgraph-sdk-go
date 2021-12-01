package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InternetMessageHeader 
type InternetMessageHeader struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents the key in a key-value pair.
    name *string;
    // The value in a key-value pair.
    value *string;
}
// NewInternetMessageHeader instantiates a new internetMessageHeader and sets the default values.
func NewInternetMessageHeader()(*InternetMessageHeader) {
    m := &InternetMessageHeader{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InternetMessageHeader) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetName gets the name property value. Represents the key in a key-value pair.
func (m *InternetMessageHeader) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetValue gets the value property value. The value in a key-value pair.
func (m *InternetMessageHeader) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InternetMessageHeader) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
func (m *InternetMessageHeader) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InternetMessageHeader) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *InternetMessageHeader) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. Represents the key in a key-value pair.
func (m *InternetMessageHeader) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetValue sets the value property value. The value in a key-value pair.
func (m *InternetMessageHeader) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
