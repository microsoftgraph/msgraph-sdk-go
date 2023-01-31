package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody 
type OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // 
    groupId *string
    // 
    notebookFolder *string
    // 
    renameAs *string
    // 
    siteCollectionId *string
    // 
    siteId *string
}
// NewOnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody instantiates a new OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody and sets the default values.
func NewOnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody()(*OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) {
    m := &OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateOnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["notebookFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotebookFolder(val)
        }
        return nil
    }
    res["renameAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenameAs(val)
        }
        return nil
    }
    res["siteCollectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCollectionId(val)
        }
        return nil
    }
    res["siteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) GetGroupId()(*string) {
    return m.groupId
}
// GetNotebookFolder gets the notebookFolder property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) GetNotebookFolder()(*string) {
    return m.notebookFolder
}
// GetRenameAs gets the renameAs property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) GetRenameAs()(*string) {
    return m.renameAs
}
// GetSiteCollectionId gets the siteCollectionId property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) GetSiteCollectionId()(*string) {
    return m.siteCollectionId
}
// GetSiteId gets the siteId property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) GetSiteId()(*string) {
    return m.siteId
}
// Serialize serializes information the current object
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notebookFolder", m.GetNotebookFolder())
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
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGroupId sets the groupId property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) SetGroupId(value *string)() {
    m.groupId = value
}
// SetNotebookFolder sets the notebookFolder property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) SetNotebookFolder(value *string)() {
    m.notebookFolder = value
}
// SetRenameAs sets the renameAs property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) SetRenameAs(value *string)() {
    m.renameAs = value
}
// SetSiteCollectionId sets the siteCollectionId property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) SetSiteCollectionId(value *string)() {
    m.siteCollectionId = value
}
// SetSiteId sets the siteId property value. 
func (m *OnenoteNotebooksItemMicrosoftGraphCopyNotebookCopyNotebookPostRequestBody) SetSiteId(value *string)() {
    m.siteId = value
}
