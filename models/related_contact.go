package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RelatedContact 
type RelatedContact struct {
    // Indicates whether the user has been consented to access student data.
    accessConsent *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the contact. Required.
    displayName *string
    // Email address of the contact.
    emailAddress *string
    // Mobile phone number of the contact.
    mobilePhone *string
    // Relationship to the user. Possible values are: parent, relative, aide, doctor, guardian, child, other, unknownFutureValue.
    relationship *ContactRelationship
}
// NewRelatedContact instantiates a new relatedContact and sets the default values.
func NewRelatedContact()(*RelatedContact) {
    m := &RelatedContact{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRelatedContactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRelatedContactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelatedContact(), nil
}
// GetAccessConsent gets the accessConsent property value. Indicates whether the user has been consented to access student data.
func (m *RelatedContact) GetAccessConsent()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accessConsent
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RelatedContact) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Name of the contact. Required.
func (m *RelatedContact) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEmailAddress gets the emailAddress property value. Email address of the contact.
func (m *RelatedContact) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RelatedContact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessConsent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessConsent(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["mobilePhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["relationship"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseContactRelationship)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelationship(val.(*ContactRelationship))
        }
        return nil
    }
    return res
}
// GetMobilePhone gets the mobilePhone property value. Mobile phone number of the contact.
func (m *RelatedContact) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// GetRelationship gets the relationship property value. Relationship to the user. Possible values are: parent, relative, aide, doctor, guardian, child, other, unknownFutureValue.
func (m *RelatedContact) GetRelationship()(*ContactRelationship) {
    if m == nil {
        return nil
    } else {
        return m.relationship
    }
}
// Serialize serializes information the current object
func (m *RelatedContact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetRelationship()).String()
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
// SetAccessConsent sets the accessConsent property value. Indicates whether the user has been consented to access student data.
func (m *RelatedContact) SetAccessConsent(value *bool)() {
    if m != nil {
        m.accessConsent = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RelatedContact) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Name of the contact. Required.
func (m *RelatedContact) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEmailAddress sets the emailAddress property value. Email address of the contact.
func (m *RelatedContact) SetEmailAddress(value *string)() {
    if m != nil {
        m.emailAddress = value
    }
}
// SetMobilePhone sets the mobilePhone property value. Mobile phone number of the contact.
func (m *RelatedContact) SetMobilePhone(value *string)() {
    if m != nil {
        m.mobilePhone = value
    }
}
// SetRelationship sets the relationship property value. Relationship to the user. Possible values are: parent, relative, aide, doctor, guardian, child, other, unknownFutureValue.
func (m *RelatedContact) SetRelationship(value *ContactRelationship)() {
    if m != nil {
        m.relationship = value
    }
}
