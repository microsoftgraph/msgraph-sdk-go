package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AlertTrigger 
type AlertTrigger struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the property serving as a detection trigger.
    name *string;
    // Type of the property in the key:value pair for interpretation. For example, String, Boolean etc.
    type *string;
    // Value of the property serving as a detection trigger.
    value *string;
}
// NewAlertTrigger instantiates a new alertTrigger and sets the default values.
func NewAlertTrigger()(*AlertTrigger) {
    m := &AlertTrigger{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertTrigger) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetName gets the name property value. Name of the property serving as a detection trigger.
func (m *AlertTrigger) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetType gets the type property value. Type of the property in the key:value pair for interpretation. For example, String, Boolean etc.
func (m *AlertTrigger) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type
    }
}
// GetValue gets the value property value. Value of the property serving as a detection trigger.
func (m *AlertTrigger) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertTrigger) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
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
func (m *AlertTrigger) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AlertTrigger) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *AlertTrigger) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. Name of the property serving as a detection trigger.
func (m *AlertTrigger) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetType sets the type property value. Type of the property in the key:value pair for interpretation. For example, String, Boolean etc.
func (m *AlertTrigger) SetType(value *string)() {
    if m != nil {
        m.type = value
    }
}
// SetValue sets the value property value. Value of the property serving as a detection trigger.
func (m *AlertTrigger) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}
