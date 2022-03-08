package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamMemberSettingsable 
type TeamMemberSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowAddRemoveApps()(*bool)
    GetAllowCreatePrivateChannels()(*bool)
    GetAllowCreateUpdateChannels()(*bool)
    GetAllowCreateUpdateRemoveConnectors()(*bool)
    GetAllowCreateUpdateRemoveTabs()(*bool)
    GetAllowDeleteChannels()(*bool)
    SetAllowAddRemoveApps(value *bool)()
    SetAllowCreatePrivateChannels(value *bool)()
    SetAllowCreateUpdateChannels(value *bool)()
    SetAllowCreateUpdateRemoveConnectors(value *bool)()
    SetAllowCreateUpdateRemoveTabs(value *bool)()
    SetAllowDeleteChannels(value *bool)()
}
