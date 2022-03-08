package invite

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// InviteRequestBodyable 
type InviteRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetExpirationDateTime()(*string)
    GetMessage()(*string)
    GetPassword()(*string)
    GetRecipients()([]i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveRecipientable)
    GetRequireSignIn()(*bool)
    GetRetainInheritedPermissions()(*bool)
    GetRoles()([]string)
    GetSendInvitation()(*bool)
    SetExpirationDateTime(value *string)()
    SetMessage(value *string)()
    SetPassword(value *string)()
    SetRecipients(value []i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveRecipientable)()
    SetRequireSignIn(value *bool)()
    SetRetainInheritedPermissions(value *bool)()
    SetRoles(value []string)()
    SetSendInvitation(value *bool)()
}
