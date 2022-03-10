package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceComplianceActionItemable 
type DeviceComplianceActionItemable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActionType()(*DeviceComplianceActionType)
    GetGracePeriodHours()(*int32)
    GetNotificationMessageCCList()([]string)
    GetNotificationTemplateId()(*string)
    SetActionType(value *DeviceComplianceActionType)()
    SetGracePeriodHours(value *int32)()
    SetNotificationMessageCCList(value []string)()
    SetNotificationTemplateId(value *string)()
}
