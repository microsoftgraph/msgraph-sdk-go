package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ControlScore struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Control action category (Identity, Data, Device, Apps, Infrastructure).
    controlCategory *string;
    // Control unique name.
    controlName *string;
    // Description of the control.
    description *string;
    // Tenant achieved score for the control (it varies day by day depending on tenant operations on the control).
    score *float64;
}
// Instantiates a new controlScore and sets the default values.
func NewControlScore()(*ControlScore) {
    m := &ControlScore{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ControlScore) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
func (m *ControlScore) GetControlCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlCategory
    }
}
// Gets the controlName property value. Control unique name.
func (m *ControlScore) GetControlName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlName
    }
}
// Gets the description property value. Description of the control.
func (m *ControlScore) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the score property value. Tenant achieved score for the control (it varies day by day depending on tenant operations on the control).
func (m *ControlScore) GetScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.score
    }
}
// The deserialization information for the current model
func (m *ControlScore) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["controlCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetControlCategory(val)
        return nil
    }
    res["controlName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetControlName(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["score"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetScore(val)
        return nil
    }
    return res
}
func (m *ControlScore) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ControlScore) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("controlCategory", m.GetControlCategory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("controlName", m.GetControlName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("score", m.GetScore())
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
func (m *ControlScore) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
// Parameters:
//  - value : Value to set for the controlCategory property.
func (m *ControlScore) SetControlCategory(value *string)() {
    m.controlCategory = value
}
// Sets the controlName property value. Control unique name.
// Parameters:
//  - value : Value to set for the controlName property.
func (m *ControlScore) SetControlName(value *string)() {
    m.controlName = value
}
// Sets the description property value. Description of the control.
// Parameters:
//  - value : Value to set for the description property.
func (m *ControlScore) SetDescription(value *string)() {
    m.description = value
}
// Sets the score property value. Tenant achieved score for the control (it varies day by day depending on tenant operations on the control).
// Parameters:
//  - value : Value to set for the score property.
func (m *ControlScore) SetScore(value *float64)() {
    m.score = value
}
