package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RelatedContact struct {
    // Indicates whether the user has been consented to access student data.
    accessConsent *bool;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Name of the contact. Required.
    displayName *string;
    // Primary email address of the contact.
    emailAddress *string;
    // Mobile phone number of the contact.
    mobilePhone *string;
    // Relationship to the user. Possible values are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue.
    relationship *ContactRelationship;
}
// Instantiates a new relatedContact and sets the default values.
func NewRelatedContact()(*RelatedContact) {
    m := &RelatedContact{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the accessConsent property value. Indicates whether the user has been consented to access student data.
func (m *RelatedContact) GetAccessConsent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accessConsent
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RelatedContact) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. Name of the contact. Required.
func (m *RelatedContact) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the emailAddress property value. Primary email address of the contact.
func (m *RelatedContact) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// Gets the mobilePhone property value. Mobile phone number of the contact.
func (m *RelatedContact) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// Gets the relationship property value. Relationship to the user. Possible values are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue.
func (m *RelatedContact) GetRelationship()(*ContactRelationship) {
    if m == nil {
        return nil
    } else {
        return m.relationship
    }
}
// The deserialization information for the current model
func (m *RelatedContact) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessConsent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessConsent(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["mobilePhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["relationship"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseContactRelationship)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ContactRelationship)
            m.SetRelationship(&cast)
        }
        return nil
    }
    return res
}
func (m *RelatedContact) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RelatedContact) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("accessConsent", m.GetAccessConsent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
        if err != nil {
            return err
        }
    }
    if m.GetRelationship() != nil {
        cast := m.GetRelationship().String()
        err := writer.WriteStringValue("relationship", &cast)
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
// Sets the accessConsent property value. Indicates whether the user has been consented to access student data.
// Parameters:
//  - value : Value to set for the accessConsent property.
func (m *RelatedContact) SetAccessConsent(value *bool)() {
    m.accessConsent = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RelatedContact) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. Name of the contact. Required.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *RelatedContact) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the emailAddress property value. Primary email address of the contact.
// Parameters:
//  - value : Value to set for the emailAddress property.
func (m *RelatedContact) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// Sets the mobilePhone property value. Mobile phone number of the contact.
// Parameters:
//  - value : Value to set for the mobilePhone property.
func (m *RelatedContact) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// Sets the relationship property value. Relationship to the user. Possible values are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue.
// Parameters:
//  - value : Value to set for the relationship property.
func (m *RelatedContact) SetRelationship(value *ContactRelationship)() {
    m.relationship = value
}
