package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamsTabConfiguration struct {
    additionalData map[string]interface{};
    contentUrl *string;
    entityId *string;
    removeUrl *string;
    websiteUrl *string;
}
func NewTeamsTabConfiguration()(*TeamsTabConfiguration) {
    m := &TeamsTabConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TeamsTabConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TeamsTabConfiguration) GetContentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentUrl
    }
}
func (m *TeamsTabConfiguration) GetEntityId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.entityId
    }
}
func (m *TeamsTabConfiguration) GetRemoveUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.removeUrl
    }
}
func (m *TeamsTabConfiguration) GetWebsiteUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.websiteUrl
    }
}
func (m *TeamsTabConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentUrl(val)
        return nil
    }
    res["entityId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEntityId(val)
        return nil
    }
    res["removeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemoveUrl(val)
        return nil
    }
    res["websiteUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebsiteUrl(val)
        return nil
    }
    return res
}
func (m *TeamsTabConfiguration) IsNil()(bool) {
    return m == nil
}
func (m *TeamsTabConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("entityId", m.GetEntityId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("removeUrl", m.GetRemoveUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("websiteUrl", m.GetWebsiteUrl())
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
func (m *TeamsTabConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TeamsTabConfiguration) SetContentUrl(value *string)() {
    m.contentUrl = value
}
func (m *TeamsTabConfiguration) SetEntityId(value *string)() {
    m.entityId = value
}
func (m *TeamsTabConfiguration) SetRemoveUrl(value *string)() {
    m.removeUrl = value
}
func (m *TeamsTabConfiguration) SetWebsiteUrl(value *string)() {
    m.websiteUrl = value
}
