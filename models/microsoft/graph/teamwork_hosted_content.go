package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkHostedContent 
type TeamworkHostedContent struct {
    Entity
    // Write only. Bytes for the hosted content (such as images).
    contentBytes []byte;
    // Write only. Content type. sicj as image/png, image/jpg.
    contentType *string;
}
// NewTeamworkHostedContent instantiates a new teamworkHostedContent and sets the default values.
func NewTeamworkHostedContent()(*TeamworkHostedContent) {
    m := &TeamworkHostedContent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkHostedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkHostedContentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamworkHostedContent(), nil
}
// GetContentBytes gets the contentBytes property value. Write only. Bytes for the hosted content (such as images).
func (m *TeamworkHostedContent) GetContentBytes()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.contentBytes
    }
}
// GetContentType gets the contentType property value. Write only. Content type. sicj as image/png, image/jpg.
func (m *TeamworkHostedContent) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkHostedContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentBytes(val)
        }
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TeamworkHostedContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("contentBytes", m.GetContentBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentBytes sets the contentBytes property value. Write only. Bytes for the hosted content (such as images).
func (m *TeamworkHostedContent) SetContentBytes(value []byte)() {
    if m != nil {
        m.contentBytes = value
    }
}
// SetContentType sets the contentType property value. Write only. Content type. sicj as image/png, image/jpg.
func (m *TeamworkHostedContent) SetContentType(value *string)() {
    if m != nil {
        m.contentType = value
    }
}
