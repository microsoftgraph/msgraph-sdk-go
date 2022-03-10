package create

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// CreateRequestBodyable 
type CreateRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCertificateSigningRequest()(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintCertificateSigningRequestable)
    GetConnectorId()(*string)
    GetDisplayName()(*string)
    GetHasPhysicalDevice()(*bool)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetPhysicalDeviceId()(*string)
    SetCertificateSigningRequest(value i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintCertificateSigningRequestable)()
    SetConnectorId(value *string)()
    SetDisplayName(value *string)()
    SetHasPhysicalDevice(value *bool)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetPhysicalDeviceId(value *string)()
}
