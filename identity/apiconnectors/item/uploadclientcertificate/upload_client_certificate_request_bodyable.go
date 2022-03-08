package uploadclientcertificate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UploadClientCertificateRequestBodyable 
type UploadClientCertificateRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetPassword()(*string)
    GetPkcs12Value()(*string)
    SetPassword(value *string)()
    SetPkcs12Value(value *string)()
}
