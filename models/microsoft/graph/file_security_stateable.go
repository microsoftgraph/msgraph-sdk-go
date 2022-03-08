package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FileSecurityStateable 
type FileSecurityStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetFileHash()(FileHashable)
    GetName()(*string)
    GetPath()(*string)
    GetRiskScore()(*string)
    SetFileHash(value FileHashable)()
    SetName(value *string)()
    SetPath(value *string)()
    SetRiskScore(value *string)()
}
