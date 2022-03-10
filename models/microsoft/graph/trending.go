package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Trending provides operations to manage the collection of drive entities.
type Trending struct {
    Entity
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Used for navigating to the trending document.
    resource Entityable;
    // Reference properties of the trending document, such as the url and type of the document.
    resourceReference ResourceReferenceable;
    // Properties that you can use to visualize the document in your experience.
    resourceVisualization ResourceVisualizationable;
    // Value indicating how much the document is currently trending. The larger the number, the more the document is currently trending around the user (the more relevant it is). Returned documents are sorted by this value.
    weight *float64;
}
// NewTrending instantiates a new trending and sets the default values.
func NewTrending()(*Trending) {
    m := &Trending{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrendingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrendingFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTrending(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Trending) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(Entityable))
        }
        return nil
    }
    res["resourceReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceReference(val.(ResourceReferenceable))
        }
        return nil
    }
    res["resourceVisualization"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceVisualizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceVisualization(val.(ResourceVisualizationable))
        }
        return nil
    }
    res["weight"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeight(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *Trending) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetResource gets the resource property value. Used for navigating to the trending document.
func (m *Trending) GetResource()(Entityable) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetResourceReference gets the resourceReference property value. Reference properties of the trending document, such as the url and type of the document.
func (m *Trending) GetResourceReference()(ResourceReferenceable) {
    if m == nil {
        return nil
    } else {
        return m.resourceReference
    }
}
// GetResourceVisualization gets the resourceVisualization property value. Properties that you can use to visualize the document in your experience.
func (m *Trending) GetResourceVisualization()(ResourceVisualizationable) {
    if m == nil {
        return nil
    } else {
        return m.resourceVisualization
    }
}
// GetWeight gets the weight property value. Value indicating how much the document is currently trending. The larger the number, the more the document is currently trending around the user (the more relevant it is). Returned documents are sorted by this value.
func (m *Trending) GetWeight()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.weight
    }
}
func (m *Trending) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Trending) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resourceReference", m.GetResourceReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resourceVisualization", m.GetResourceVisualization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("weight", m.GetWeight())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *Trending) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetResource sets the resource property value. Used for navigating to the trending document.
func (m *Trending) SetResource(value Entityable)() {
    if m != nil {
        m.resource = value
    }
}
// SetResourceReference sets the resourceReference property value. Reference properties of the trending document, such as the url and type of the document.
func (m *Trending) SetResourceReference(value ResourceReferenceable)() {
    if m != nil {
        m.resourceReference = value
    }
}
// SetResourceVisualization sets the resourceVisualization property value. Properties that you can use to visualize the document in your experience.
func (m *Trending) SetResourceVisualization(value ResourceVisualizationable)() {
    if m != nil {
        m.resourceVisualization = value
    }
}
// SetWeight sets the weight property value. Value indicating how much the document is currently trending. The larger the number, the more the document is currently trending around the user (the more relevant it is). Returned documents are sorted by this value.
func (m *Trending) SetWeight(value *float64)() {
    if m != nil {
        m.weight = value
    }
}
