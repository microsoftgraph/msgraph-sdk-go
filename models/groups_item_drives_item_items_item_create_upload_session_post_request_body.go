package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody provides operations to call the createUploadSession method.
type GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The item property
    item DriveItemUploadablePropertiesable
}
// NewGroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody instantiates a new GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody and sets the default values.
func NewGroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody()(*GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody) {
    m := &GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["item"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateDriveItemUploadablePropertiesFromDiscriminatorValue , m.SetItem)
    return res
}
// GetItem gets the item property value. The item property
func (m *GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody) GetItem()(DriveItemUploadablePropertiesable) {
    return m.item
}
// Serialize serializes information the current object
func (m *GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("item", m.GetItem())
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
func (m *GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetItem sets the item property value. The item property
func (m *GroupsItemDrivesItemItemsItemCreateUploadSessionPostRequestBody) SetItem(value DriveItemUploadablePropertiesable)() {
    m.item = value
}
