package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserAttributeValuesItem 
type UserAttributeValuesItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Determines whether the value is set as the default.
    isDefault *bool;
    // The display name of the property displayed to the user in the user flow.
    name *string;
    // The value that is set when this item is selected.
    value *string;
}
// NewUserAttributeValuesItem instantiates a new userAttributeValuesItem and sets the default values.
func NewUserAttributeValuesItem()(*UserAttributeValuesItem) {
    m := &UserAttributeValuesItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAttributeValuesItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsDefault gets the isDefault property value. Determines whether the value is set as the default.
func (m *UserAttributeValuesItem) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetName gets the name property value. The display name of the property displayed to the user in the user flow.
func (m *UserAttributeValuesItem) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetValue gets the value property value. The value that is set when this item is selected.
func (m *UserAttributeValuesItem) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAttributeValuesItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
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
func (m *UserAttributeValuesItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserAttributeValuesItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
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
func (m *UserAttributeValuesItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsDefault sets the isDefault property value. Determines whether the value is set as the default.
func (m *UserAttributeValuesItem) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetName sets the name property value. The display name of the property displayed to the user in the user flow.
func (m *UserAttributeValuesItem) SetName(value *string)() {
    m.name = value
}
// SetValue sets the value property value. The value that is set when this item is selected.
func (m *UserAttributeValuesItem) SetValue(value *string)() {
    m.value = value
}
