package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody provides operations to call the copyToNotebook method.
type UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The groupId property
    groupId *string
    // The id property
    id *string
    // The renameAs property
    renameAs *string
    // The siteCollectionId property
    siteCollectionId *string
    // The siteId property
    siteId *string
}
// NewUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody instantiates a new UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody and sets the default values.
func NewUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody()(*UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) {
    m := &UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGroupId)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["renameAs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRenameAs)
    res["siteCollectionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSiteCollectionId)
    res["siteId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSiteId)
    return res
}
// GetGroupId gets the groupId property value. The groupId property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetGroupId()(*string) {
    return m.groupId
}
// GetId gets the id property value. The id property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetId()(*string) {
    return m.id
}
// GetRenameAs gets the renameAs property value. The renameAs property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetRenameAs()(*string) {
    return m.renameAs
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetSiteCollectionId()(*string) {
    return m.siteCollectionId
}
// GetSiteId gets the siteId property value. The siteId property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetSiteId()(*string) {
    return m.siteId
}
// Serialize serializes information the current object
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
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
        err := writer.WriteStringValue("renameAs", m.GetRenameAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteCollectionId", m.GetSiteCollectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteId", m.GetSiteId())
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
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetGroupId(value *string)() {
    m.groupId = value
}
// SetId sets the id property value. The id property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetId(value *string)() {
    m.id = value
}
// SetRenameAs sets the renameAs property value. The renameAs property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetRenameAs(value *string)() {
    m.renameAs = value
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetSiteCollectionId(value *string)() {
    m.siteCollectionId = value
}
// SetSiteId sets the siteId property value. The siteId property
func (m *UsersItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetSiteId(value *string)() {
    m.siteId = value
}
