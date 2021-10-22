package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DriveRecipient struct {
    additionalData map[string]interface{};
    alias *string;
    email *string;
    objectId *string;
}
func NewDriveRecipient()(*DriveRecipient) {
    m := &DriveRecipient{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DriveRecipient) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DriveRecipient) GetAlias()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alias
    }
}
func (m *DriveRecipient) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
func (m *DriveRecipient) GetObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectId
    }
}
func (m *DriveRecipient) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["alias"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAlias(val)
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["objectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetObjectId(val)
        return nil
    }
    return res
}
func (m *DriveRecipient) IsNil()(bool) {
    return m == nil
}
func (m *DriveRecipient) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alias", m.GetAlias())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("objectId", m.GetObjectId())
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
func (m *DriveRecipient) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DriveRecipient) SetAlias(value *string)() {
    m.alias = value
}
func (m *DriveRecipient) SetEmail(value *string)() {
    m.email = value
}
func (m *DriveRecipient) SetObjectId(value *string)() {
    m.objectId = value
}
