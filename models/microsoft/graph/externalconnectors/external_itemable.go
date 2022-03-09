package externalconnectors

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExternalItemable 
type ExternalItemable interface {
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAcl()([]Aclable)
    GetContent()(ExternalItemContentable)
    GetProperties()(Propertiesable)
    SetAcl(value []Aclable)()
    SetContent(value ExternalItemContentable)()
    SetProperties(value Propertiesable)()
}
