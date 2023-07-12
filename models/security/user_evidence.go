package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserEvidence 
type UserEvidence struct {
    AlertEvidence
}
// NewUserEvidence instantiates a new userEvidence and sets the default values.
func NewUserEvidence()(*UserEvidence) {
    m := &UserEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateUserEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserEvidence) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserAccount gets the userAccount property value. The user account details.
func (m *UserEvidence) GetUserAccount()(UserAccountable) {
    val, err := m.GetBackingStore().Get("userAccount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserAccountable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserEvidence) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAccount sets the userAccount property value. The user account details.
func (m *UserEvidence) SetUserAccount(value UserAccountable)() {
    err := m.GetBackingStore().Set("userAccount", value)
    if err != nil {
        panic(err)
    }
}
// UserEvidenceable 
type UserEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetUserAccount()(UserAccountable)
    SetOdataType(value *string)()
    SetUserAccount(value UserAccountable)()
}
