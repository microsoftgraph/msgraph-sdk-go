package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Driveable 
type Driveable interface {
    BaseItemable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBundles()([]DriveItemable)
    GetDriveType()(*string)
    GetFollowing()([]DriveItemable)
    GetItems()([]DriveItemable)
    GetList()(Listable)
    GetOwner()(IdentitySetable)
    GetQuota()(Quotaable)
    GetRoot()(DriveItemable)
    GetSharePointIds()(SharepointIdsable)
    GetSpecial()([]DriveItemable)
    GetSystem()(SystemFacetable)
    SetBundles(value []DriveItemable)()
    SetDriveType(value *string)()
    SetFollowing(value []DriveItemable)()
    SetItems(value []DriveItemable)()
    SetList(value Listable)()
    SetOwner(value IdentitySetable)()
    SetQuota(value Quotaable)()
    SetRoot(value DriveItemable)()
    SetSharePointIds(value SharepointIdsable)()
    SetSpecial(value []DriveItemable)()
    SetSystem(value SystemFacetable)()
}
