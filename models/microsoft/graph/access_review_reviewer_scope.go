package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewReviewerScope 
type AccessReviewReviewerScope struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The query specifying who will be the reviewer. See table for examples.
    query *string;
    // In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
    queryRoot *string;
    // The type of query. Examples include MicrosoftGraph and ARM.
    queryType *string;
}
// NewAccessReviewReviewerScope instantiates a new accessReviewReviewerScope and sets the default values.
func NewAccessReviewReviewerScope()(*AccessReviewReviewerScope) {
    m := &AccessReviewReviewerScope{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessReviewReviewerScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewReviewerScopeFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAccessReviewReviewerScope(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewReviewerScope) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewReviewerScope) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["query"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    res["queryRoot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryRoot(val)
        }
        return nil
    }
    res["queryType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryType(val)
        }
        return nil
    }
    return res
}
// GetQuery gets the query property value. The query specifying who will be the reviewer. See table for examples.
func (m *AccessReviewReviewerScope) GetQuery()(*string) {
    if m == nil {
        return nil
    } else {
        return m.query
    }
}
// GetQueryRoot gets the queryRoot property value. In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
func (m *AccessReviewReviewerScope) GetQueryRoot()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryRoot
    }
}
// GetQueryType gets the queryType property value. The type of query. Examples include MicrosoftGraph and ARM.
func (m *AccessReviewReviewerScope) GetQueryType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.queryType
    }
}
// Serialize serializes information the current object
func (m *AccessReviewReviewerScope) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryRoot", m.GetQueryRoot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryType", m.GetQueryType())
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
func (m *AccessReviewReviewerScope) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetQuery sets the query property value. The query specifying who will be the reviewer. See table for examples.
func (m *AccessReviewReviewerScope) SetQuery(value *string)() {
    if m != nil {
        m.query = value
    }
}
// SetQueryRoot sets the queryRoot property value. In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
func (m *AccessReviewReviewerScope) SetQueryRoot(value *string)() {
    if m != nil {
        m.queryRoot = value
    }
}
// SetQueryType sets the queryType property value. The type of query. Examples include MicrosoftGraph and ARM.
func (m *AccessReviewReviewerScope) SetQueryType(value *string)() {
    if m != nil {
        m.queryType = value
    }
}
