package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceDecisionItemResource 
type AccessReviewInstanceDecisionItemResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Display name of the resource
    displayName *string
    // Identifier of the resource
    id *string
    // Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
    type_escaped *string
}
// NewAccessReviewInstanceDecisionItemResource instantiates a new accessReviewInstanceDecisionItemResource and sets the default values.
func NewAccessReviewInstanceDecisionItemResource()(*AccessReviewInstanceDecisionItemResource) {
    m := &AccessReviewInstanceDecisionItemResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessReviewInstanceDecisionItemResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewInstanceDecisionItemResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewInstanceDecisionItemResource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewInstanceDecisionItemResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Display name of the resource
func (m *AccessReviewInstanceDecisionItemResource) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewInstanceDecisionItemResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Identifier of the resource
func (m *AccessReviewInstanceDecisionItemResource) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetType gets the type property value. Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
func (m *AccessReviewInstanceDecisionItemResource) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Serialize serializes information the current object
func (m *AccessReviewInstanceDecisionItemResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *AccessReviewInstanceDecisionItemResource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the resource
func (m *AccessReviewInstanceDecisionItemResource) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetId sets the id property value. Identifier of the resource
func (m *AccessReviewInstanceDecisionItemResource) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetType sets the type property value. Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
func (m *AccessReviewInstanceDecisionItemResource) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
