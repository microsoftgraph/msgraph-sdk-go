package copytonotebook

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new copyToNotebookRequestBody and sets the default values.
func NewCopyToNotebookRequestBody()(*CopyToNotebookRequestBody) {
    m := &CopyToNotebookRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyToNotebookRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the groupId property value. 
func (m *CopyToNotebookRequestBody) GetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupId
    }
}
// Gets the id property value. 
func (m *CopyToNotebookRequestBody) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the renameAs property value. 
func (m *CopyToNotebookRequestBody) GetRenameAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.renameAs
    }
}
// Gets the siteCollectionId property value. 
func (m *CopyToNotebookRequestBody) GetSiteCollectionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteCollectionId
    }
}
// Gets the siteId property value. 
func (m *CopyToNotebookRequestBody) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
// The deserialization information for the current model
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
func (m *CopyToNotebookRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CopyToNotebookRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the groupId property value. 
// Parameters:
//  - value : Value to set for the groupId property.
func (m *CopyToNotebookRequestBody) SetGroupId(value *string)() {
    m.groupId = value
}
// Sets the id property value. 
// Parameters:
//  - value : Value to set for the id property.
func (m *CopyToNotebookRequestBody) SetId(value *string)() {
    m.id = value
}
// Sets the renameAs property value. 
// Parameters:
//  - value : Value to set for the renameAs property.
func (m *CopyToNotebookRequestBody) SetRenameAs(value *string)() {
    m.renameAs = value
}
// Sets the siteCollectionId property value. 
// Parameters:
//  - value : Value to set for the siteCollectionId property.
func (m *CopyToNotebookRequestBody) SetSiteCollectionId(value *string)() {
    m.siteCollectionId = value
}
// Sets the siteId property value. 
// Parameters:
//  - value : Value to set for the siteId property.
func (m *CopyToNotebookRequestBody) SetSiteId(value *string)() {
    m.siteId = value
}
