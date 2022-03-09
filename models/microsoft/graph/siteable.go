package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/termstore"
)

// Siteable 
type Siteable interface {
    BaseItemable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAnalytics()(ItemAnalyticsable)
    GetColumns()([]ColumnDefinitionable)
    GetContentTypes()([]ContentTypeable)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetDrives()([]Driveable)
    GetError()(PublicErrorable)
    GetExternalColumns()([]ColumnDefinitionable)
    GetItems()([]BaseItemable)
    GetLists()([]Listable)
    GetOnenote()(Onenoteable)
    GetPermissions()([]Permissionable)
    GetRoot()(Rootable)
    GetSharepointIds()(SharepointIdsable)
    GetSiteCollection()(SiteCollectionable)
    GetSites()([]Siteable)
    GetTermStore()(id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Storeable)
    GetTermStores()([]id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Storeable)
    SetAnalytics(value ItemAnalyticsable)()
    SetColumns(value []ColumnDefinitionable)()
    SetContentTypes(value []ContentTypeable)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetDrives(value []Driveable)()
    SetError(value PublicErrorable)()
    SetExternalColumns(value []ColumnDefinitionable)()
    SetItems(value []BaseItemable)()
    SetLists(value []Listable)()
    SetOnenote(value Onenoteable)()
    SetPermissions(value []Permissionable)()
    SetRoot(value Rootable)()
    SetSharepointIds(value SharepointIdsable)()
    SetSiteCollection(value SiteCollectionable)()
    SetSites(value []Siteable)()
    SetTermStore(value id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Storeable)()
    SetTermStores(value []id62b8df0892707d421d6e0a5aefa589248c11f95794bf4122483a0ef812fad7d.Storeable)()
}
