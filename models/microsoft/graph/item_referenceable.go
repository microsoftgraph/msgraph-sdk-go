package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ItemReferenceable 
type ItemReferenceable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDriveId()(*string)
    GetDriveType()(*string)
    GetId()(*string)
    GetName()(*string)
    GetPath()(*string)
    GetShareId()(*string)
    GetSharepointIds()(SharepointIdsable)
    GetSiteId()(*string)
    SetDriveId(value *string)()
    SetDriveType(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetPath(value *string)()
    SetShareId(value *string)()
    SetSharepointIds(value SharepointIdsable)()
    SetSiteId(value *string)()
}
