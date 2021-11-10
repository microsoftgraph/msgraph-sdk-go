package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceComplianceActionItem struct {
    Entity
    // What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification.
    actionType *DeviceComplianceActionType;
    // Number of hours to wait till the action will be enforced. Valid values 0 to 8760
    gracePeriodHours *int32;
    // A list of group IDs to speicify who to CC this notification message to.
    notificationMessageCCList []string;
    // What notification Message template to use
    notificationTemplateId *string;
}
// Instantiates a new deviceComplianceActionItem and sets the default values.
func NewDeviceComplianceActionItem()(*DeviceComplianceActionItem) {
    m := &DeviceComplianceActionItem{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the actionType property value. What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification.
func (m *DeviceComplianceActionItem) GetActionType()(*DeviceComplianceActionType) {
    if m == nil {
        return nil
    } else {
        return m.actionType
    }
}
// Gets the gracePeriodHours property value. Number of hours to wait till the action will be enforced. Valid values 0 to 8760
func (m *DeviceComplianceActionItem) GetGracePeriodHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.gracePeriodHours
    }
}
// Gets the notificationMessageCCList property value. A list of group IDs to speicify who to CC this notification message to.
func (m *DeviceComplianceActionItem) GetNotificationMessageCCList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notificationMessageCCList
    }
}
// Gets the notificationTemplateId property value. What notification Message template to use
func (m *DeviceComplianceActionItem) GetNotificationTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTemplateId
    }
}
// The deserialization information for the current model
func (m *DeviceComplianceActionItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceActionType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceComplianceActionType)
            m.SetActionType(&cast)
        }
        return nil
    }
    res["gracePeriodHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodHours(val)
        }
        return nil
    }
    res["notificationMessageCCList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotificationMessageCCList(res)
        }
        return nil
    }
    res["notificationTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTemplateId(val)
        }
        return nil
    }
    return res
}
func (m *DeviceComplianceActionItem) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceComplianceActionItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionType() != nil {
        cast := m.GetActionType().String()
        err = writer.WriteStringValue("actionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("gracePeriodHours", m.GetGracePeriodHours())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("notificationMessageCCList", m.GetNotificationMessageCCList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationTemplateId", m.GetNotificationTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the actionType property value. What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification.
// Parameters:
//  - value : Value to set for the actionType property.
func (m *DeviceComplianceActionItem) SetActionType(value *DeviceComplianceActionType)() {
    m.actionType = value
}
// Sets the gracePeriodHours property value. Number of hours to wait till the action will be enforced. Valid values 0 to 8760
// Parameters:
//  - value : Value to set for the gracePeriodHours property.
func (m *DeviceComplianceActionItem) SetGracePeriodHours(value *int32)() {
    m.gracePeriodHours = value
}
// Sets the notificationMessageCCList property value. A list of group IDs to speicify who to CC this notification message to.
// Parameters:
//  - value : Value to set for the notificationMessageCCList property.
func (m *DeviceComplianceActionItem) SetNotificationMessageCCList(value []string)() {
    m.notificationMessageCCList = value
}
// Sets the notificationTemplateId property value. What notification Message template to use
// Parameters:
//  - value : Value to set for the notificationTemplateId property.
func (m *DeviceComplianceActionItem) SetNotificationTemplateId(value *string)() {
    m.notificationTemplateId = value
}
