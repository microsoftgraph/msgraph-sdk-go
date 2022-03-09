package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnenoteResource provides operations to manage the educationRoot singleton.
type OnenoteResource struct {
    OnenoteEntityBaseModel
    // The content stream
    content []byte;
    // The URL for downloading the content
    contentUrl *string;
}
// NewOnenoteResource instantiates a new onenoteResource and sets the default values.
func NewOnenoteResource()(*OnenoteResource) {
    m := &OnenoteResource{
        OnenoteEntityBaseModel: *NewOnenoteEntityBaseModel(),
    }
    return m
}
// CreateOnenoteResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteResourceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOnenoteResource(), nil
}
// GetContent gets the content property value. The content stream
func (m *OnenoteResource) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// GetContentUrl gets the contentUrl property value. The URL for downloading the content
func (m *OnenoteResource) GetContentUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteResource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OnenoteEntityBaseModel.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentUrl(val)
        }
        return nil
    }
    return res
}
func (m *OnenoteResource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OnenoteResource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OnenoteEntityBaseModel.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content stream
func (m *OnenoteResource) SetContent(value []byte)() {
    if m != nil {
        m.content = value
    }
}
// SetContentUrl sets the contentUrl property value. The URL for downloading the content
func (m *OnenoteResource) SetContentUrl(value *string)() {
    if m != nil {
        m.contentUrl = value
    }
}
