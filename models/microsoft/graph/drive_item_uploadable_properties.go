package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DriveItemUploadableProperties 
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
// NewDriveItemUploadableProperties instantiates a new driveItemUploadableProperties and sets the default values.
func NewDriveItemUploadableProperties()(*DriveItemUploadableProperties) {
    m := &DriveItemUploadableProperties{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveItemUploadableProperties) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. Provides a user-visible description of the item. Read-write. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFileSize gets the fileSize property value. Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) GetFileSize()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.fileSize
    }
}
// GetFileSystemInfo gets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItemUploadableProperties) GetFileSystemInfo()(*FileSystemInfo) {
    if m == nil {
        return nil
    } else {
        return m.fileSystemInfo
    }
}
// GetName gets the name property value. The name of the item (filename and extension). Read-write.
func (m *DriveItemUploadableProperties) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveItemUploadableProperties) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. Provides a user-visible description of the item. Read-write. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetFileSize sets the fileSize property value. Provides an expected file size to perform a quota check prior to upload. Only on OneDrive Personal.
func (m *DriveItemUploadableProperties) SetFileSize(value *int64)() {
    if m != nil {
        m.fileSize = value
    }
}
// SetFileSystemInfo sets the fileSystemInfo property value. File system information on client. Read-write.
func (m *DriveItemUploadableProperties) SetFileSystemInfo(value *FileSystemInfo)() {
    if m != nil {
        m.fileSystemInfo = value
    }
}
// SetName sets the name property value. The name of the item (filename and extension). Read-write.
func (m *DriveItemUploadableProperties) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
