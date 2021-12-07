package copytosection

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CopyToSectionRequestBody 
type CopyToSectionRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    groupId *string;
    // 
    id *string;
    // 
    siteCollectionId *string;
    // 
    siteId *string;
}
// NewCopyToSectionRequestBody instantiates a new copyToSectionRequestBody and sets the default values.
func NewCopyToSectionRequestBody()(*CopyToSectionRequestBody) {
    m := &CopyToSectionRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyToSectionRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetGroupId gets the groupId property value. 
func (m *CopyToSectionRequestBody) GetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupId
    }
}
// GetId gets the id property value. 
func (m *CopyToSectionRequestBody) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetSiteCollectionId gets the siteCollectionId property value. 
func (m *CopyToSectionRequestBody) GetSiteCollectionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteCollectionId
    }
}
// GetSiteId gets the siteId property value. 
func (m *CopyToSectionRequestBody) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CopyToSectionRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *CopyToSectionRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CopyToSectionRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *CopyToSectionRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupId sets the groupId property value. 
func (m *CopyToSectionRequestBody) SetGroupId(value *string)() {
    if m != nil {
        m.groupId = value
    }
}
// SetId sets the id property value. 
func (m *CopyToSectionRequestBody) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetSiteCollectionId sets the siteCollectionId property value. 
func (m *CopyToSectionRequestBody) SetSiteCollectionId(value *string)() {
    if m != nil {
        m.siteCollectionId = value
    }
}
// SetSiteId sets the siteId property value. 
func (m *CopyToSectionRequestBody) SetSiteId(value *string)() {
    if m != nil {
        m.siteId = value
    }
}
