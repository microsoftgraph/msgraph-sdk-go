package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChatMessageHostedContent provides operations to manage the collection of chat entities.
type ChatMessageHostedContent struct {
    TeamworkHostedContent
}
// NewChatMessageHostedContent instantiates a new chatMessageHostedContent and sets the default values.
func NewChatMessageHostedContent()(*ChatMessageHostedContent) {
    m := &ChatMessageHostedContent{
        TeamworkHostedContent: *NewTeamworkHostedContent(),
    }
    return m
}
// CreateChatMessageHostedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageHostedContentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewChatMessageHostedContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageHostedContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.TeamworkHostedContent.GetFieldDeserializers()
    return res
}
func (m *ChatMessageHostedContent) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChatMessageHostedContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.TeamworkHostedContent.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
