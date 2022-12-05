package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemDrivesItemRootCopyPostRequestBody provides operations to call the copy method.
type GroupsItemDrivesItemRootCopyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name property
    name *string
    // The parentReference property
    parentReference ItemReferenceable
}
// NewGroupsItemDrivesItemRootCopyPostRequestBody instantiates a new GroupsItemDrivesItemRootCopyPostRequestBody and sets the default values.
func NewGroupsItemDrivesItemRootCopyPostRequestBody()(*GroupsItemDrivesItemRootCopyPostRequestBody) {
    m := &GroupsItemDrivesItemRootCopyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemDrivesItemRootCopyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemDrivesItemRootCopyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemDrivesItemRootCopyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemDrivesItemRootCopyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemDrivesItemRootCopyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parentReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentReference(val.(ItemReferenceable))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
func (m *GroupsItemDrivesItemRootCopyPostRequestBody) GetName()(*string) {
    return m.name
}
// GetParentReference gets the parentReference property value. The parentReference property
func (m *GroupsItemDrivesItemRootCopyPostRequestBody) GetParentReference()(ItemReferenceable) {
    return m.parentReference
}
// Serialize serializes information the current object
func (m *GroupsItemDrivesItemRootCopyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parentReference", m.GetParentReference())
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
func (m *GroupsItemDrivesItemRootCopyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *GroupsItemDrivesItemRootCopyPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetParentReference sets the parentReference property value. The parentReference property
func (m *GroupsItemDrivesItemRootCopyPostRequestBody) SetParentReference(value ItemReferenceable)() {
    m.parentReference = value
}
