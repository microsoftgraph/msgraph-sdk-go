package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentitySet 
type IdentitySet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Optional. The application associated with this action.
    application *Identity;
    // Optional. The device associated with this action.
    device *Identity;
    // Optional. The user associated with this action.
    user *Identity;
}
// NewIdentitySet instantiates a new identitySet and sets the default values.
func NewIdentitySet()(*IdentitySet) {
    m := &IdentitySet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentitySet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplication gets the application property value. Optional. The application associated with this action.
func (m *IdentitySet) GetApplication()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// GetDevice gets the device property value. Optional. The device associated with this action.
func (m *IdentitySet) GetDevice()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.device
    }
}
// GetUser gets the user property value. Optional. The user associated with this action.
func (m *IdentitySet) GetUser()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentitySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(*Identity))
        }
        return nil
    }
    res["device"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(*Identity))
        }
        return nil
    }
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(*Identity))
        }
        return nil
    }
    return res
}
func (m *IdentitySet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IdentitySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *IdentitySet) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplication sets the application property value. Optional. The application associated with this action.
func (m *IdentitySet) SetApplication(value *Identity)() {
    if m != nil {
        m.application = value
    }
}
// SetDevice sets the device property value. Optional. The device associated with this action.
func (m *IdentitySet) SetDevice(value *Identity)() {
    if m != nil {
        m.device = value
    }
}
// SetUser sets the user property value. Optional. The user associated with this action.
func (m *IdentitySet) SetUser(value *Identity)() {
    if m != nil {
        m.user = value
    }
}
