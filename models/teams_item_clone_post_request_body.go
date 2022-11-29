package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsItemClonePostRequestBody provides operations to call the clone method.
type TeamsItemClonePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The classification property
    classification *string
    // The description property
    description *string
    // The displayName property
    displayName *string
    // The mailNickname property
    mailNickname *string
    // The partsToClone property
    partsToClone *ClonableTeamParts
    // The visibility property
    visibility *TeamVisibilityType
}
// NewTeamsItemClonePostRequestBody instantiates a new TeamsItemClonePostRequestBody and sets the default values.
func NewTeamsItemClonePostRequestBody()(*TeamsItemClonePostRequestBody) {
    m := &TeamsItemClonePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamsItemClonePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsItemClonePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsItemClonePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamsItemClonePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClassification gets the classification property value. The classification property
func (m *TeamsItemClonePostRequestBody) GetClassification()(*string) {
    return m.classification
}
// GetDescription gets the description property value. The description property
func (m *TeamsItemClonePostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *TeamsItemClonePostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsItemClonePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetClassification)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["mailNickname"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMailNickname)
    res["partsToClone"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseClonableTeamParts , m.SetPartsToClone)
    res["visibility"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseTeamVisibilityType , m.SetVisibility)
    return res
}
// GetMailNickname gets the mailNickname property value. The mailNickname property
func (m *TeamsItemClonePostRequestBody) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetPartsToClone gets the partsToClone property value. The partsToClone property
func (m *TeamsItemClonePostRequestBody) GetPartsToClone()(*ClonableTeamParts) {
    return m.partsToClone
}
// GetVisibility gets the visibility property value. The visibility property
func (m *TeamsItemClonePostRequestBody) GetVisibility()(*TeamVisibilityType) {
    return m.visibility
}
// Serialize serializes information the current object
func (m *TeamsItemClonePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("classification", m.GetClassification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
        err := writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    if m.GetPartsToClone() != nil {
        cast := (*m.GetPartsToClone()).String()
        err := writer.WriteStringValue("partsToClone", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVisibility() != nil {
        cast := (*m.GetVisibility()).String()
        err := writer.WriteStringValue("visibility", &cast)
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
func (m *TeamsItemClonePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClassification sets the classification property value. The classification property
func (m *TeamsItemClonePostRequestBody) SetClassification(value *string)() {
    m.classification = value
}
// SetDescription sets the description property value. The description property
func (m *TeamsItemClonePostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *TeamsItemClonePostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetMailNickname sets the mailNickname property value. The mailNickname property
func (m *TeamsItemClonePostRequestBody) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetPartsToClone sets the partsToClone property value. The partsToClone property
func (m *TeamsItemClonePostRequestBody) SetPartsToClone(value *ClonableTeamParts)() {
    m.partsToClone = value
}
// SetVisibility sets the visibility property value. The visibility property
func (m *TeamsItemClonePostRequestBody) SetVisibility(value *TeamVisibilityType)() {
    m.visibility = value
}
