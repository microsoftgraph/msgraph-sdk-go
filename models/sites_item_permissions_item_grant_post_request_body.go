package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SitesItemPermissionsItemGrantPostRequestBody provides operations to call the grant method.
type SitesItemPermissionsItemGrantPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The recipients property
    recipients []DriveRecipientable
    // The roles property
    roles []string
}
// NewSitesItemPermissionsItemGrantPostRequestBody instantiates a new SitesItemPermissionsItemGrantPostRequestBody and sets the default values.
func NewSitesItemPermissionsItemGrantPostRequestBody()(*SitesItemPermissionsItemGrantPostRequestBody) {
    m := &SitesItemPermissionsItemGrantPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSitesItemPermissionsItemGrantPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSitesItemPermissionsItemGrantPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSitesItemPermissionsItemGrantPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SitesItemPermissionsItemGrantPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SitesItemPermissionsItemGrantPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["recipients"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDriveRecipientFromDiscriminatorValue , m.SetRecipients)
    res["roles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoles)
    return res
}
// GetRecipients gets the recipients property value. The recipients property
func (m *SitesItemPermissionsItemGrantPostRequestBody) GetRecipients()([]DriveRecipientable) {
    return m.recipients
}
// GetRoles gets the roles property value. The roles property
func (m *SitesItemPermissionsItemGrantPostRequestBody) GetRoles()([]string) {
    return m.roles
}
// Serialize serializes information the current object
func (m *SitesItemPermissionsItemGrantPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetRecipients() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRecipients())
        err := writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        err := writer.WriteCollectionOfStringValues("roles", m.GetRoles())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SitesItemPermissionsItemGrantPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *SitesItemPermissionsItemGrantPostRequestBody) SetRecipients(value []DriveRecipientable)() {
    m.recipients = value
}
// SetRoles sets the roles property value. The roles property
func (m *SitesItemPermissionsItemGrantPostRequestBody) SetRoles(value []string)() {
    m.roles = value
}
