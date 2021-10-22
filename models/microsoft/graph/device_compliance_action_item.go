package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceComplianceActionItem struct {
    Entity
    actionType *DeviceComplianceActionType;
    gracePeriodHours *int32;
    notificationMessageCCList []string;
    notificationTemplateId *string;
}
func NewDeviceComplianceActionItem()(*DeviceComplianceActionItem) {
    m := &DeviceComplianceActionItem{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceComplianceActionItem) GetActionType()(*DeviceComplianceActionType) {
    if m == nil {
        return nil
    } else {
        return m.actionType
    }
}
func (m *DeviceComplianceActionItem) GetGracePeriodHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.gracePeriodHours
    }
}
func (m *DeviceComplianceActionItem) GetNotificationMessageCCList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notificationMessageCCList
    }
}
func (m *DeviceComplianceActionItem) GetNotificationTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTemplateId
    }
}
func (m *DeviceComplianceActionItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceActionType)
        if err != nil {
            return err
        }
        cast := val.(DeviceComplianceActionType)
        m.SetActionType(&cast)
        return nil
    }
    res["gracePeriodHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetGracePeriodHours(val)
        return nil
    }
    res["notificationMessageCCList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetNotificationMessageCCList(res)
        return nil
    }
    res["notificationTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationTemplateId(val)
        return nil
    }
    return res
}
func (m *DeviceComplianceActionItem) IsNil()(bool) {
    return m == nil
}
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
func (m *DeviceComplianceActionItem) SetActionType(value *DeviceComplianceActionType)() {
    m.actionType = value
}
func (m *DeviceComplianceActionItem) SetGracePeriodHours(value *int32)() {
    m.gracePeriodHours = value
}
func (m *DeviceComplianceActionItem) SetNotificationMessageCCList(value []string)() {
    m.notificationMessageCCList = value
}
func (m *DeviceComplianceActionItem) SetNotificationTemplateId(value *string)() {
    m.notificationTemplateId = value
}
