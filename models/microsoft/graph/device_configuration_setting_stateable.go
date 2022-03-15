package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceConfigurationSettingStateable 
type DeviceConfigurationSettingStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCurrentValue()(*string)
    GetErrorCode()(*int32)
    GetErrorDescription()(*string)
    GetInstanceDisplayName()(*string)
    GetSetting()(*string)
    GetSettingName()(*string)
    GetSources()([]SettingSourceable)
    GetState()(*ComplianceStatus)
    GetUserEmail()(*string)
    GetUserId()(*string)
    GetUserName()(*string)
    GetUserPrincipalName()(*string)
    SetCurrentValue(value *string)()
    SetErrorCode(value *int32)()
    SetErrorDescription(value *string)()
    SetInstanceDisplayName(value *string)()
    SetSetting(value *string)()
    SetSettingName(value *string)()
    SetSources(value []SettingSourceable)()
    SetState(value *ComplianceStatus)()
    SetUserEmail(value *string)()
    SetUserId(value *string)()
    SetUserName(value *string)()
    SetUserPrincipalName(value *string)()
}
