package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ListItemable 
type ListItemable interface {
    BaseItemable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAnalytics()(ItemAnalyticsable)
    GetContentType()(ContentTypeInfoable)
    GetDriveItem()(DriveItemable)
    GetFields()(FieldValueSetable)
    GetSharepointIds()(SharepointIdsable)
    GetVersions()([]ListItemVersionable)
    SetAnalytics(value ItemAnalyticsable)()
    SetContentType(value ContentTypeInfoable)()
    SetDriveItem(value DriveItemable)()
    SetFields(value FieldValueSetable)()
    SetSharepointIds(value SharepointIdsable)()
    SetVersions(value []ListItemVersionable)()
}
