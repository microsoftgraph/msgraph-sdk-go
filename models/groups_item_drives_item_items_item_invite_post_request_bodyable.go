package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupsItemDrivesItemItemsItemInvitePostRequestBodyable 
type GroupsItemDrivesItemItemsItemInvitePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*string)
    GetMessage()(*string)
    GetPassword()(*string)
    GetRecipients()([]DriveRecipientable)
    GetRequireSignIn()(*bool)
    GetRetainInheritedPermissions()(*bool)
    GetRoles()([]string)
    GetSendInvitation()(*bool)
    SetExpirationDateTime(value *string)()
    SetMessage(value *string)()
    SetPassword(value *string)()
    SetRecipients(value []DriveRecipientable)()
    SetRequireSignIn(value *bool)()
    SetRetainInheritedPermissions(value *bool)()
    SetRoles(value []string)()
    SetSendInvitation(value *bool)()
}
