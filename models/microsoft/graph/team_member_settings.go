package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamMemberSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If set to true, members can add and remove apps.
    allowAddRemoveApps *bool;
    // If set to true, members can add and update private channels.
    allowCreatePrivateChannels *bool;
    // If set to true, members can add and update channels.
    allowCreateUpdateChannels *bool;
    // If set to true, members can add, update, and remove connectors.
    allowCreateUpdateRemoveConnectors *bool;
    // If set to true, members can add, update, and remove tabs.
    allowCreateUpdateRemoveTabs *bool;
    // If set to true, members can delete channels.
    allowDeleteChannels *bool;
}
// Instantiates a new teamMemberSettings and sets the default values.
func NewTeamMemberSettings()(*TeamMemberSettings) {
    m := &TeamMemberSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamMemberSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowAddRemoveApps property value. If set to true, members can add and remove apps.
func (m *TeamMemberSettings) GetAllowAddRemoveApps()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAddRemoveApps
    }
}
// Gets the allowCreatePrivateChannels property value. If set to true, members can add and update private channels.
func (m *TeamMemberSettings) GetAllowCreatePrivateChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreatePrivateChannels
    }
}
// Gets the allowCreateUpdateChannels property value. If set to true, members can add and update channels.
func (m *TeamMemberSettings) GetAllowCreateUpdateChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreateUpdateChannels
    }
}
// Gets the allowCreateUpdateRemoveConnectors property value. If set to true, members can add, update, and remove connectors.
func (m *TeamMemberSettings) GetAllowCreateUpdateRemoveConnectors()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreateUpdateRemoveConnectors
    }
}
// Gets the allowCreateUpdateRemoveTabs property value. If set to true, members can add, update, and remove tabs.
func (m *TeamMemberSettings) GetAllowCreateUpdateRemoveTabs()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowCreateUpdateRemoveTabs
    }
}
// Gets the allowDeleteChannels property value. If set to true, members can delete channels.
func (m *TeamMemberSettings) GetAllowDeleteChannels()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowDeleteChannels
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TeamMemberSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowAddRemoveApps property value. If set to true, members can add and remove apps.
// Parameters:
//  - value : Value to set for the allowAddRemoveApps property.
func (m *TeamMemberSettings) SetAllowAddRemoveApps(value *bool)() {
    m.allowAddRemoveApps = value
}
// Sets the allowCreatePrivateChannels property value. If set to true, members can add and update private channels.
// Parameters:
//  - value : Value to set for the allowCreatePrivateChannels property.
func (m *TeamMemberSettings) SetAllowCreatePrivateChannels(value *bool)() {
    m.allowCreatePrivateChannels = value
}
// Sets the allowCreateUpdateChannels property value. If set to true, members can add and update channels.
// Parameters:
//  - value : Value to set for the allowCreateUpdateChannels property.
func (m *TeamMemberSettings) SetAllowCreateUpdateChannels(value *bool)() {
    m.allowCreateUpdateChannels = value
}
// Sets the allowCreateUpdateRemoveConnectors property value. If set to true, members can add, update, and remove connectors.
// Parameters:
//  - value : Value to set for the allowCreateUpdateRemoveConnectors property.
func (m *TeamMemberSettings) SetAllowCreateUpdateRemoveConnectors(value *bool)() {
    m.allowCreateUpdateRemoveConnectors = value
}
// Sets the allowCreateUpdateRemoveTabs property value. If set to true, members can add, update, and remove tabs.
// Parameters:
//  - value : Value to set for the allowCreateUpdateRemoveTabs property.
func (m *TeamMemberSettings) SetAllowCreateUpdateRemoveTabs(value *bool)() {
    m.allowCreateUpdateRemoveTabs = value
}
// Sets the allowDeleteChannels property value. If set to true, members can delete channels.
// Parameters:
//  - value : Value to set for the allowDeleteChannels property.
func (m *TeamMemberSettings) SetAllowDeleteChannels(value *bool)() {
    m.allowDeleteChannels = value
}
