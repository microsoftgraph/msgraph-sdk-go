package copytonotebook

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopyToNotebookRequestBody provides operations to call the copyToNotebook method.
type CopyToNotebookRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The groupId property
    groupId *string;
    // The id property
    id *string;
    // The renameAs property
    renameAs *string;
    // The siteCollectionId property
    siteCollectionId *string;
    // The siteId property
    siteId *string;
}
// NewCopyToNotebookRequestBody instantiates a new copyToNotebookRequestBody and sets the default values.
func NewCopyToNotebookRequestBody()(*CopyToNotebookRequestBody) {
    m := &CopyToNotebookRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCopyToNotebookRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCopyToNotebookRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopyToNotebookRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyToNotebookRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CopyToNotebookRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["renameAs"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenameAs(val)
        }
        return nil
    }
    res["siteCollectionId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCollectionId(val)
        }
        return nil
    }
    res["siteId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetGroupId gets the groupId property value. The groupId property
func (m *CopyToNotebookRequestBody) GetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupId
    }
}
// GetId gets the id property value. The id property
func (m *CopyToNotebookRequestBody) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetRenameAs gets the renameAs property value. The renameAs property
func (m *CopyToNotebookRequestBody) GetRenameAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.renameAs
    }
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *CopyToNotebookRequestBody) GetSiteCollectionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteCollectionId
    }
}
// GetSiteId gets the siteId property value. The siteId property
func (m *CopyToNotebookRequestBody) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
// Serialize serializes information the current object
func (m *CopyToNotebookRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CopyToNotebookRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupId sets the groupId property value. The groupId property
func (m *CopyToNotebookRequestBody) SetGroupId(value *string)() {
    if m != nil {
        m.groupId = value
    }
}
// SetId sets the id property value. The id property
func (m *CopyToNotebookRequestBody) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetRenameAs sets the renameAs property value. The renameAs property
func (m *CopyToNotebookRequestBody) SetRenameAs(value *string)() {
    if m != nil {
        m.renameAs = value
    }
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *CopyToNotebookRequestBody) SetSiteCollectionId(value *string)() {
    if m != nil {
        m.siteCollectionId = value
    }
}
// SetSiteId sets the siteId property value. The siteId property
func (m *CopyToNotebookRequestBody) SetSiteId(value *string)() {
    if m != nil {
        m.siteId = value
    }
}
