package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterBaseable 
type PrinterBaseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetCapabilities()(PrinterCapabilitiesable)
    GetDefaults()(PrinterDefaultsable)
    GetDisplayName()(*string)
    GetIsAcceptingJobs()(*bool)
    GetJobs()([]PrintJobable)
    GetLocation()(PrinterLocationable)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetStatus()(PrinterStatusable)
    SetCapabilities(value PrinterCapabilitiesable)()
    SetDefaults(value PrinterDefaultsable)()
    SetDisplayName(value *string)()
    SetIsAcceptingJobs(value *bool)()
    SetJobs(value []PrintJobable)()
    SetLocation(value PrinterLocationable)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetStatus(value PrinterStatusable)()
}
