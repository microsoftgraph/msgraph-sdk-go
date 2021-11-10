package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new driveRecipient and sets the default values.
func NewDriveRecipient()(*DriveRecipient) {
    m := &DriveRecipient{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveRecipient) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the alias property value. The alias of the domain object, for cases where an email address is unavailable (e.g. security groups).
func (m *DriveRecipient) GetAlias()(*string) {
    if m == nil {
        return nil
    } else {
        return m.alias
    }
}
// Gets the email property value. The email address for the recipient, if the recipient has an associated email address.
func (m *DriveRecipient) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the objectId property value. The unique identifier for the recipient in the directory.
func (m *DriveRecipient) GetObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectId
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DriveRecipient) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the alias property value. The alias of the domain object, for cases where an email address is unavailable (e.g. security groups).
// Parameters:
//  - value : Value to set for the alias property.
func (m *DriveRecipient) SetAlias(value *string)() {
    m.alias = value
}
// Sets the email property value. The email address for the recipient, if the recipient has an associated email address.
// Parameters:
//  - value : Value to set for the email property.
func (m *DriveRecipient) SetEmail(value *string)() {
    m.email = value
}
// Sets the objectId property value. The unique identifier for the recipient in the directory.
// Parameters:
//  - value : Value to set for the objectId property.
func (m *DriveRecipient) SetObjectId(value *string)() {
    m.objectId = value
}
