package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ServiceHealthIssuePost struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The published time of the post.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The content of the service issue post.
    description *ItemBody;
    // The post type of the service issue historical post. Possible values are: regular, quick, strategic, unknownFutureValue.
    postType *PostType;
}
// Instantiates a new serviceHealthIssuePost and sets the default values.
func NewServiceHealthIssuePost()(*ServiceHealthIssuePost) {
    m := &ServiceHealthIssuePost{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceHealthIssuePost) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the createdDateTime property value. The published time of the post.
func (m *ServiceHealthIssuePost) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The content of the service issue post.
func (m *ServiceHealthIssuePost) GetDescription()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the postType property value. The post type of the service issue historical post. Possible values are: regular, quick, strategic, unknownFutureValue.
func (m *ServiceHealthIssuePost) GetPostType()(*PostType) {
    if m == nil {
        return nil
    } else {
        return m.postType
    }
}
// The deserialization information for the current model
func (m *ServiceHealthIssuePost) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val.(*ItemBody))
        }
        return nil
    }
    res["postType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePostType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PostType)
            m.SetPostType(&cast)
        }
        return nil
    }
    return res
}
func (m *ServiceHealthIssuePost) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ServiceHealthIssuePost) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetPostType() != nil {
        cast := m.GetPostType().String()
        err := writer.WriteStringValue("postType", &cast)
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
func (m *ServiceHealthIssuePost) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the createdDateTime property value. The published time of the post.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ServiceHealthIssuePost) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The content of the service issue post.
// Parameters:
//  - value : Value to set for the description property.
func (m *ServiceHealthIssuePost) SetDescription(value *ItemBody)() {
    m.description = value
}
// Sets the postType property value. The post type of the service issue historical post. Possible values are: regular, quick, strategic, unknownFutureValue.
// Parameters:
//  - value : Value to set for the postType property.
func (m *ServiceHealthIssuePost) SetPostType(value *PostType)() {
    m.postType = value
}
