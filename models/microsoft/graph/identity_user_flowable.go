package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityUserFlowable 
type IdentityUserFlowable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetUserFlowType()(*UserFlowType)
    GetUserFlowTypeVersion()(*float32)
    SetUserFlowType(value *UserFlowType)()
    SetUserFlowTypeVersion(value *float32)()
}
