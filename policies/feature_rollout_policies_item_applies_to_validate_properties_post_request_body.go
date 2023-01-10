package policies

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody 
type FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The displayName property
    displayName *string
    // The entityType property
    entityType *string
    // The mailNickname property
    mailNickname *string
    // The onBehalfOfUserId property
    onBehalfOfUserId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewFeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody instantiates a new FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody()(*FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) {
    m := &FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetEntityType gets the entityType property value. The entityType property
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) GetEntityType()(*string) {
    return m.entityType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["entityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityType(val)
        }
        return nil
    }
    res["mailNickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["onBehalfOfUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnBehalfOfUserId(val)
        }
        return nil
    }
    return res
}
// GetMailNickname gets the mailNickname property value. The mailNickname property
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetOnBehalfOfUserId gets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) GetOnBehalfOfUserId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.onBehalfOfUserId
}
// Serialize serializes information the current object
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("entityType", m.GetEntityType())
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
    {
        err := writer.WriteUUIDValue("onBehalfOfUserId", m.GetOnBehalfOfUserId())
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
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEntityType sets the entityType property value. The entityType property
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) SetEntityType(value *string)() {
    m.entityType = value
}
// SetMailNickname sets the mailNickname property value. The mailNickname property
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetOnBehalfOfUserId sets the onBehalfOfUserId property value. The onBehalfOfUserId property
func (m *FeatureRolloutPoliciesItemAppliesToValidatePropertiesPostRequestBody) SetOnBehalfOfUserId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.onBehalfOfUserId = value
}
