package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Hashesable 
type Hashesable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCrc32Hash()(*string)
    GetQuickXorHash()(*string)
    GetSha1Hash()(*string)
    GetSha256Hash()(*string)
    SetCrc32Hash(value *string)()
    SetQuickXorHash(value *string)()
    SetSha1Hash(value *string)()
    SetSha256Hash(value *string)()
}
