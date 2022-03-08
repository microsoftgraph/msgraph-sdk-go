package copytosection

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CopyToSectionRequestBodyable 
type CopyToSectionRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetGroupId()(*string)
    GetId()(*string)
    GetSiteCollectionId()(*string)
    GetSiteId()(*string)
    SetGroupId(value *string)()
    SetId(value *string)()
    SetSiteCollectionId(value *string)()
    SetSiteId(value *string)()
}
