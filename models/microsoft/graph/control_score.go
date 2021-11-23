package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// controlScore 
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
// NewControlScore instantiates a new controlScore and sets the default values.
func NewControlScore()(*ControlScore) {
    m := &ControlScore{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ControlScore) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetControlCategory gets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
func (m *ControlScore) GetControlCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlCategory
    }
}
// GetControlName gets the controlName property value. Control unique name.
func (m *ControlScore) GetControlName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlName
    }
}
// GetDescription gets the description property value. Description of the control.
func (m *ControlScore) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetScore gets the score property value. Tenant achieved score for the control (it varies day by day depending on tenant operations on the control).
func (m *ControlScore) GetScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.score
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ControlScore) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["controlCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlCategory(val)
        }
        return nil
    }
    res["controlName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlName(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["score"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    return res
}
func (m *ControlScore) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ControlScore) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetControlCategory sets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
func (m *ControlScore) SetControlCategory(value *string)() {
    m.controlCategory = value
}
// SetControlName sets the controlName property value. Control unique name.
func (m *ControlScore) SetControlName(value *string)() {
    m.controlName = value
}
// SetDescription sets the description property value. Description of the control.
func (m *ControlScore) SetDescription(value *string)() {
    m.description = value
}
// SetScore sets the score property value. Tenant achieved score for the control (it varies day by day depending on tenant operations on the control).
func (m *ControlScore) SetScore(value *float64)() {
    m.score = value
}
