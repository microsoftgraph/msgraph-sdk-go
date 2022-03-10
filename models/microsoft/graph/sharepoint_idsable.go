package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharepointIdsable 
type SharepointIdsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetListId()(*string)
    GetListItemId()(*string)
    GetListItemUniqueId()(*string)
    GetSiteId()(*string)
    GetSiteUrl()(*string)
    GetTenantId()(*string)
    GetWebId()(*string)
    SetListId(value *string)()
    SetListItemId(value *string)()
    SetListItemUniqueId(value *string)()
    SetSiteId(value *string)()
    SetSiteUrl(value *string)()
    SetTenantId(value *string)()
    SetWebId(value *string)()
}
