package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DriveItemUploadablePropertiesable 
type DriveItemUploadablePropertiesable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetDescription()(*string)
    GetFileSize()(*int64)
    GetFileSystemInfo()(FileSystemInfoable)
    GetName()(*string)
    SetDescription(value *string)()
    SetFileSize(value *int64)()
    SetFileSystemInfo(value FileSystemInfoable)()
    SetName(value *string)()
}
