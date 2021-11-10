package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DriveItemUploadableProperties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Provides a user-visible description of the item. Read-write. Only on OneDrive Personal.
    description *string;
    // Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal.
    fileSize *int64;
    // File system information on client. Read-write.
    fileSystemInfo *FileSystemInfo;
    // The name of the item (filename and extension). Read-write.
    name *string;
}
// Instantiates a new driveItemUploadableProperties and sets the default values.
func NewDriveItemUploadableProperties()(*DriveItemUploadableProperties) {
    m := &DriveItemUploadableProperties{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveItemUploadableProperties) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. Provides a user-visible description of the item. Read-write. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the fileSize property value. Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) GetFileSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileSize
    }
}
// Gets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItemUploadableProperties) GetFileSystemInfo()(*FileSystemInfo) {
    if m == nil {
        return nil
    } else {
        return m.fileSystemInfo
    }
}
// Gets the name property value. The name of the item (filename and extension). Read-write.
func (m *DriveItemUploadableProperties) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// The deserialization information for the current model
func (m *DriveItemUploadableProperties) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["fileSize"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSize(val)
        }
        return nil
    }
    res["fileSystemInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileSystemInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileSystemInfo(val.(*FileSystemInfo))
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
func (m *DriveItemUploadableProperties) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DriveItemUploadableProperties) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("fileSize", m.GetFileSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fileSystemInfo", m.GetFileSystemInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DriveItemUploadableProperties) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. Provides a user-visible description of the item. Read-write. Only on OneDrive Personal.
// Parameters:
//  - value : Value to set for the description property.
func (m *DriveItemUploadableProperties) SetDescription(value *string)() {
    m.description = value
}
// Sets the fileSize property value. Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal.
// Parameters:
//  - value : Value to set for the fileSize property.
func (m *DriveItemUploadableProperties) SetFileSize(value *int64)() {
    m.fileSize = value
}
// Sets the fileSystemInfo property value. File system information on client. Read-write.
// Parameters:
//  - value : Value to set for the fileSystemInfo property.
func (m *DriveItemUploadableProperties) SetFileSystemInfo(value *FileSystemInfo)() {
    m.fileSystemInfo = value
}
// Sets the name property value. The name of the item (filename and extension). Read-write.
// Parameters:
//  - value : Value to set for the name property.
func (m *DriveItemUploadableProperties) SetName(value *string)() {
    m.name = value
}
