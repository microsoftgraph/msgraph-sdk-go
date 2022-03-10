package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Listable 
type Listable interface {
    BaseItemable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetColumns()([]ColumnDefinitionable)
    GetContentTypes()([]ContentTypeable)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetItems()([]ListItemable)
    GetList()(ListInfoable)
    GetSharepointIds()(SharepointIdsable)
    GetSubscriptions()([]Subscriptionable)
    GetSystem()(SystemFacetable)
    SetColumns(value []ColumnDefinitionable)()
    SetContentTypes(value []ContentTypeable)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetItems(value []ListItemable)()
    SetList(value ListInfoable)()
    SetSharepointIds(value SharepointIdsable)()
    SetSubscriptions(value []Subscriptionable)()
    SetSystem(value SystemFacetable)()
}
