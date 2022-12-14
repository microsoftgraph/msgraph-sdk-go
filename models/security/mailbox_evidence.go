package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailboxEvidence 
type MailboxEvidence struct {
    AlertEvidence
    // The displayName property
    displayName *string
    // The primaryAddress property
    primaryAddress *string
    // The userAccount property
    userAccount UserAccountable
}
// NewMailboxEvidence instantiates a new MailboxEvidence and sets the default values.
func NewMailboxEvidence()(*MailboxEvidence) {
    m := &MailboxEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateMailboxEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailboxEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxEvidence(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *MailboxEvidence) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailboxEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
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
    res["primaryAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryAddress(val)
        }
        return nil
    }
    res["userAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccount(val.(UserAccountable))
        }
        return nil
    }
    return res
}
// GetPrimaryAddress gets the primaryAddress property value. The primaryAddress property
func (m *MailboxEvidence) GetPrimaryAddress()(*string) {
    return m.primaryAddress
}
// GetUserAccount gets the userAccount property value. The userAccount property
func (m *MailboxEvidence) GetUserAccount()(UserAccountable) {
    return m.userAccount
}
// Serialize serializes information the current object
func (m *MailboxEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryAddress", m.GetPrimaryAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userAccount", m.GetUserAccount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *MailboxEvidence) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPrimaryAddress sets the primaryAddress property value. The primaryAddress property
func (m *MailboxEvidence) SetPrimaryAddress(value *string)() {
    m.primaryAddress = value
}
// SetUserAccount sets the userAccount property value. The userAccount property
func (m *MailboxEvidence) SetUserAccount(value UserAccountable)() {
    m.userAccount = value
}
