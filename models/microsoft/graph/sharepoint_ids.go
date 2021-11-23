package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// sharepointIds 
type SharepointIds struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The unique identifier (guid) for the item's list in SharePoint.
    listId *string;
    // An integer identifier for the item within the containing list.
    listItemId *string;
    // The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site.
    listItemUniqueId *string;
    // The unique identifier (guid) for the item's site collection (SPSite).
    siteId *string;
    // The SharePoint URL for the site that contains the item.
    siteUrl *string;
    // The unique identifier (guid) for the tenancy.
    tenantId *string;
    // The unique identifier (guid) for the item's site (SPWeb).
    webId *string;
}
// NewSharepointIds instantiates a new sharepointIds and sets the default values.
func NewSharepointIds()(*SharepointIds) {
    m := &SharepointIds{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharepointIds) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetListId gets the listId property value. The unique identifier (guid) for the item's list in SharePoint.
func (m *SharepointIds) GetListId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.listId
    }
}
// GetListItemId gets the listItemId property value. An integer identifier for the item within the containing list.
func (m *SharepointIds) GetListItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.listItemId
    }
}
// GetListItemUniqueId gets the listItemUniqueId property value. The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site.
func (m *SharepointIds) GetListItemUniqueId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.listItemUniqueId
    }
}
// GetSiteId gets the siteId property value. The unique identifier (guid) for the item's site collection (SPSite).
func (m *SharepointIds) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
// GetSiteUrl gets the siteUrl property value. The SharePoint URL for the site that contains the item.
func (m *SharepointIds) GetSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteUrl
    }
}
// GetTenantId gets the tenantId property value. The unique identifier (guid) for the tenancy.
func (m *SharepointIds) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetWebId gets the webId property value. The unique identifier (guid) for the item's site (SPWeb).
func (m *SharepointIds) GetWebId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharepointIds) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["listId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListId(val)
        }
        return nil
    }
    res["listItemId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItemId(val)
        }
        return nil
    }
    res["listItemUniqueId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItemUniqueId(val)
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
    res["siteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteUrl(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["webId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebId(val)
        }
        return nil
    }
    return res
}
func (m *SharepointIds) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SharepointIds) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("listId", m.GetListId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("listItemId", m.GetListItemId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("listItemUniqueId", m.GetListItemUniqueId())
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
        err := writer.WriteStringValue("siteUrl", m.GetSiteUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webId", m.GetWebId())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SharepointIds) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetListId sets the listId property value. The unique identifier (guid) for the item's list in SharePoint.
func (m *SharepointIds) SetListId(value *string)() {
    m.listId = value
}
// SetListItemId sets the listItemId property value. An integer identifier for the item within the containing list.
func (m *SharepointIds) SetListItemId(value *string)() {
    m.listItemId = value
}
// SetListItemUniqueId sets the listItemUniqueId property value. The unique identifier (guid) for the item within OneDrive for Business or a SharePoint site.
func (m *SharepointIds) SetListItemUniqueId(value *string)() {
    m.listItemUniqueId = value
}
// SetSiteId sets the siteId property value. The unique identifier (guid) for the item's site collection (SPSite).
func (m *SharepointIds) SetSiteId(value *string)() {
    m.siteId = value
}
// SetSiteUrl sets the siteUrl property value. The SharePoint URL for the site that contains the item.
func (m *SharepointIds) SetSiteUrl(value *string)() {
    m.siteUrl = value
}
// SetTenantId sets the tenantId property value. The unique identifier (guid) for the tenancy.
func (m *SharepointIds) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetWebId sets the webId property value. The unique identifier (guid) for the item's site (SPWeb).
func (m *SharepointIds) SetWebId(value *string)() {
    m.webId = value
}
