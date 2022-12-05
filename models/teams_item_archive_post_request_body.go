package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsItemArchivePostRequestBody provides operations to call the archive method.
type TeamsItemArchivePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The shouldSetSpoSiteReadOnlyForMembers property
    shouldSetSpoSiteReadOnlyForMembers *bool
}
// NewTeamsItemArchivePostRequestBody instantiates a new TeamsItemArchivePostRequestBody and sets the default values.
func NewTeamsItemArchivePostRequestBody()(*TeamsItemArchivePostRequestBody) {
    m := &TeamsItemArchivePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamsItemArchivePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsItemArchivePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsItemArchivePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamsItemArchivePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsItemArchivePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["shouldSetSpoSiteReadOnlyForMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldSetSpoSiteReadOnlyForMembers(val)
        }
        return nil
    }
    return res
}
// GetShouldSetSpoSiteReadOnlyForMembers gets the shouldSetSpoSiteReadOnlyForMembers property value. The shouldSetSpoSiteReadOnlyForMembers property
func (m *TeamsItemArchivePostRequestBody) GetShouldSetSpoSiteReadOnlyForMembers()(*bool) {
    return m.shouldSetSpoSiteReadOnlyForMembers
}
// Serialize serializes information the current object
func (m *TeamsItemArchivePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("shouldSetSpoSiteReadOnlyForMembers", m.GetShouldSetSpoSiteReadOnlyForMembers())
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
func (m *TeamsItemArchivePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetShouldSetSpoSiteReadOnlyForMembers sets the shouldSetSpoSiteReadOnlyForMembers property value. The shouldSetSpoSiteReadOnlyForMembers property
func (m *TeamsItemArchivePostRequestBody) SetShouldSetSpoSiteReadOnlyForMembers(value *bool)() {
    m.shouldSetSpoSiteReadOnlyForMembers = value
}
