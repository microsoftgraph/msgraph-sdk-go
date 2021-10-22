package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamMemberSettings struct {
    additionalData map[string]interface{};
    allowAddRemoveApps *bool;
    allowCreatePrivateChannels *bool;
    allowCreateUpdateChannels *bool;
    allowCreateUpdateRemoveConnectors *bool;
    allowCreateUpdateRemoveTabs *bool;
    allowDeleteChannels *bool;
}
func NewTeamMemberSettings()(*TeamMemberSettings) {
    m := &TeamMemberSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TeamMemberSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TeamMemberSettings) GetAllowAddRemoveApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAddRemoveApps
    }
}
func (m *TeamMemberSettings) GetAllowCreatePrivateChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreatePrivateChannels
    }
}
func (m *TeamMemberSettings) GetAllowCreateUpdateChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreateUpdateChannels
    }
}
func (m *TeamMemberSettings) GetAllowCreateUpdateRemoveConnectors()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreateUpdateRemoveConnectors
    }
}
func (m *TeamMemberSettings) GetAllowCreateUpdateRemoveTabs()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreateUpdateRemoveTabs
    }
}
func (m *TeamMemberSettings) GetAllowDeleteChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteChannels
    }
}
func (m *TeamMemberSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowAddRemoveApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowAddRemoveApps(val)
        return nil
    }
    res["allowCreatePrivateChannels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowCreatePrivateChannels(val)
        return nil
    }
    res["allowCreateUpdateChannels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowCreateUpdateChannels(val)
        return nil
    }
    res["allowCreateUpdateRemoveConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowCreateUpdateRemoveConnectors(val)
        return nil
    }
    res["allowCreateUpdateRemoveTabs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowCreateUpdateRemoveTabs(val)
        return nil
    }
    res["allowDeleteChannels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowDeleteChannels(val)
        return nil
    }
    return res
}
func (m *TeamMemberSettings) IsNil()(bool) {
    return m == nil
}
func (m *TeamMemberSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowAddRemoveApps", m.GetAllowAddRemoveApps())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowCreatePrivateChannels", m.GetAllowCreatePrivateChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowCreateUpdateChannels", m.GetAllowCreateUpdateChannels())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowCreateUpdateRemoveConnectors", m.GetAllowCreateUpdateRemoveConnectors())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowCreateUpdateRemoveTabs", m.GetAllowCreateUpdateRemoveTabs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowDeleteChannels", m.GetAllowDeleteChannels())
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
func (m *TeamMemberSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TeamMemberSettings) SetAllowAddRemoveApps(value *bool)() {
    m.allowAddRemoveApps = value
}
func (m *TeamMemberSettings) SetAllowCreatePrivateChannels(value *bool)() {
    m.allowCreatePrivateChannels = value
}
func (m *TeamMemberSettings) SetAllowCreateUpdateChannels(value *bool)() {
    m.allowCreateUpdateChannels = value
}
func (m *TeamMemberSettings) SetAllowCreateUpdateRemoveConnectors(value *bool)() {
    m.allowCreateUpdateRemoveConnectors = value
}
func (m *TeamMemberSettings) SetAllowCreateUpdateRemoveTabs(value *bool)() {
    m.allowCreateUpdateRemoveTabs = value
}
func (m *TeamMemberSettings) SetAllowDeleteChannels(value *bool)() {
    m.allowDeleteChannels = value
}
