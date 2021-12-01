package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DriveRecipient 
type DriveRecipient struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The alias of the domain object, for cases where an email address is unavailable (e.g. security groups).
    alias *string;
    // The email address for the recipient, if the recipient has an associated email address.
    email *string;
    // The unique identifier for the recipient in the directory.
    objectId *string;
}
// NewDriveRecipient instantiates a new driveRecipient and sets the default values.
func NewDriveRecipient()(*DriveRecipient) {
    m := &DriveRecipient{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveRecipient) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAlias gets the alias property value. The alias of the domain object, for cases where an email address is unavailable (e.g. security groups).
func (m *DriveRecipient) GetAlias()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alias
    }
}
// GetEmail gets the email property value. The email address for the recipient, if the recipient has an associated email address.
func (m *DriveRecipient) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// GetObjectId gets the objectId property value. The unique identifier for the recipient in the directory.
func (m *DriveRecipient) GetObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveRecipient) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["alias"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlias(val)
        }
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["objectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectId(val)
        }
        return nil
    }
    return res
}
func (m *DriveRecipient) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveRecipient) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAlias sets the alias property value. The alias of the domain object, for cases where an email address is unavailable (e.g. security groups).
func (m *DriveRecipient) SetAlias(value *string)() {
    if m != nil {
        m.alias = value
    }
}
// SetEmail sets the email property value. The email address for the recipient, if the recipient has an associated email address.
func (m *DriveRecipient) SetEmail(value *string)() {
    if m != nil {
        m.email = value
    }
}
// SetObjectId sets the objectId property value. The unique identifier for the recipient in the directory.
func (m *DriveRecipient) SetObjectId(value *string)() {
    if m != nil {
        m.objectId = value
    }
}
