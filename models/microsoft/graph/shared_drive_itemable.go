package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedDriveItemable 
type SharedDriveItemable interface {
    BaseItemable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDriveItem()(DriveItemable)
    GetItems()([]DriveItemable)
    GetList()(Listable)
    GetListItem()(ListItemable)
    GetOwner()(IdentitySetable)
    GetPermission()(Permissionable)
    GetRoot()(DriveItemable)
    GetSite()(Siteable)
    SetDriveItem(value DriveItemable)()
    SetItems(value []DriveItemable)()
    SetList(value Listable)()
    SetListItem(value ListItemable)()
    SetOwner(value IdentitySetable)()
    SetPermission(value Permissionable)()
    SetRoot(value DriveItemable)()
    SetSite(value Siteable)()
}
