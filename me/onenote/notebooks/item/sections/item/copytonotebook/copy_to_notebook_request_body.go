package copytonotebook

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CopyToNotebookRequestBody provides operations to call the copyToNotebook method.
type CopyToNotebookRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    groupId *string;
    // 
    id *string;
    // 
    renameAs *string;
    // 
    siteCollectionId *string;
    // 
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
func CreateCopyToNotebookRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
func (m *CopyToNotebookRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["groupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["renameAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenameAs(val)
        }
        return nil
    }
    res["siteCollectionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCollectionId(val)
        }
        return nil
    }
    res["siteId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *CopyToNotebookRequestBody) GetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupId
    }
}
// GetId gets the id property value. 
func (m *CopyToNotebookRequestBody) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetRenameAs gets the renameAs property value. 
func (m *CopyToNotebookRequestBody) GetRenameAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.renameAs
    }
}
// GetSiteCollectionId gets the siteCollectionId property value. 
func (m *CopyToNotebookRequestBody) GetSiteCollectionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteCollectionId
    }
}
// GetSiteId gets the siteId property value. 
func (m *CopyToNotebookRequestBody) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
func (m *CopyToNotebookRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CopyToNotebookRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetGroupId sets the groupId property value. 
func (m *CopyToNotebookRequestBody) SetGroupId(value *string)() {
    if m != nil {
        m.groupId = value
    }
}
// SetId sets the id property value. 
func (m *CopyToNotebookRequestBody) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetRenameAs sets the renameAs property value. 
func (m *CopyToNotebookRequestBody) SetRenameAs(value *string)() {
    if m != nil {
        m.renameAs = value
    }
}
// SetSiteCollectionId sets the siteCollectionId property value. 
func (m *CopyToNotebookRequestBody) SetSiteCollectionId(value *string)() {
    if m != nil {
        m.siteCollectionId = value
    }
}
// SetSiteId sets the siteId property value. 
func (m *CopyToNotebookRequestBody) SetSiteId(value *string)() {
    if m != nil {
        m.siteId = value
    }
}
