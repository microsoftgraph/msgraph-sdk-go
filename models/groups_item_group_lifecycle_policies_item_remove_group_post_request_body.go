package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody provides operations to call the removeGroup method.
type GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The groupId property
    groupId *string
}
// NewGroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody instantiates a new GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody and sets the default values.
func NewGroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody()(*GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody) {
    m := &GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGroupId)
    return res
}
// GetGroupId gets the groupId property value. The groupId property
func (m *GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody) GetGroupId()(*string) {
    return m.groupId
}
// Serialize serializes information the current object
func (m *GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
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
func (m *GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *GroupsItemGroupLifecyclePoliciesItemRemoveGroupPostRequestBody) SetGroupId(value *string)() {
    m.groupId = value
}
