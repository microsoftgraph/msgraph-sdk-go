package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody provides operations to call the copyToNotebook method.
type GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody struct {
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
// NewGroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody instantiates a new GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody and sets the default values.
func NewGroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody()(*GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) {
    m := &GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGroupId)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["renameAs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRenameAs)
    res["siteCollectionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSiteCollectionId)
    res["siteId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSiteId)
    return res
}
// GetGroupId gets the groupId property value. The groupId property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetGroupId()(*string) {
    return m.groupId
}
// GetId gets the id property value. The id property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetId()(*string) {
    return m.id
}
// GetRenameAs gets the renameAs property value. The renameAs property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetRenameAs()(*string) {
    return m.renameAs
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetSiteCollectionId()(*string) {
    return m.siteCollectionId
}
// GetSiteId gets the siteId property value. The siteId property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) GetSiteId()(*string) {
    return m.siteId
}
// Serialize serializes information the current object
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGroupId sets the groupId property value. The groupId property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetGroupId(value *string)() {
    m.groupId = value
}
// SetId sets the id property value. The id property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetId(value *string)() {
    m.id = value
}
// SetRenameAs sets the renameAs property value. The renameAs property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetRenameAs(value *string)() {
    m.renameAs = value
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetSiteCollectionId(value *string)() {
    m.siteCollectionId = value
}
// SetSiteId sets the siteId property value. The siteId property
func (m *GroupsItemOnenoteSectionGroupsItemSectionsItemCopyToNotebookPostRequestBody) SetSiteId(value *string)() {
    m.siteId = value
}
