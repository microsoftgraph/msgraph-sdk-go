package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsAutopilotDeviceIdentityable 
type WindowsAutopilotDeviceIdentityable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAddressableUserName()(*string)
    GetAzureActiveDirectoryDeviceId()(*string)
    GetDisplayName()(*string)
    GetEnrollmentState()(*EnrollmentState)
    GetGroupTag()(*string)
    GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagedDeviceId()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetProductKey()(*string)
    GetPurchaseOrderIdentifier()(*string)
    GetResourceName()(*string)
    GetSerialNumber()(*string)
    GetSkuNumber()(*string)
    GetSystemFamily()(*string)
    GetUserPrincipalName()(*string)
    SetAddressableUserName(value *string)()
    SetAzureActiveDirectoryDeviceId(value *string)()
    SetDisplayName(value *string)()
    SetEnrollmentState(value *EnrollmentState)()
    SetGroupTag(value *string)()
    SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagedDeviceId(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetProductKey(value *string)()
    SetPurchaseOrderIdentifier(value *string)()
    SetResourceName(value *string)()
    SetSerialNumber(value *string)()
    SetSkuNumber(value *string)()
    SetSystemFamily(value *string)()
    SetUserPrincipalName(value *string)()
}
