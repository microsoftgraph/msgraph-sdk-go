package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SharepointIds struct {
    additionalData map[string]interface{};
    listId *string;
    listItemId *string;
    listItemUniqueId *string;
    siteId *string;
    siteUrl *string;
    tenantId *string;
    webId *string;
}
func NewSharepointIds()(*SharepointIds) {
    m := &SharepointIds{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SharepointIds) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SharepointIds) GetListId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.listId
    }
}
func (m *SharepointIds) GetListItemId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.listItemId
    }
}
func (m *SharepointIds) GetListItemUniqueId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.listItemUniqueId
    }
}
func (m *SharepointIds) GetSiteId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteId
    }
}
func (m *SharepointIds) GetSiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.siteUrl
    }
}
func (m *SharepointIds) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *SharepointIds) GetWebId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webId
    }
}
func (m *SharepointIds) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["listId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetListId(val)
        return nil
    }
    res["listItemId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetListItemId(val)
        return nil
    }
    res["listItemUniqueId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetListItemUniqueId(val)
        return nil
    }
    res["siteId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSiteId(val)
        return nil
    }
    res["siteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSiteUrl(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["webId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebId(val)
        return nil
    }
    return res
}
func (m *SharepointIds) IsNil()(bool) {
    return m == nil
}
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
func (m *SharepointIds) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SharepointIds) SetListId(value *string)() {
    m.listId = value
}
func (m *SharepointIds) SetListItemId(value *string)() {
    m.listItemId = value
}
func (m *SharepointIds) SetListItemUniqueId(value *string)() {
    m.listItemUniqueId = value
}
func (m *SharepointIds) SetSiteId(value *string)() {
    m.siteId = value
}
func (m *SharepointIds) SetSiteUrl(value *string)() {
    m.siteUrl = value
}
func (m *SharepointIds) SetTenantId(value *string)() {
    m.tenantId = value
}
func (m *SharepointIds) SetWebId(value *string)() {
    m.webId = value
}
