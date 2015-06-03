// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.
//
// Version: 0.1

// STIX:Indicator

package indicator

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/jordan2175/freestix/libcybox/cybox"
	"github.com/jordan2175/freestix/libstix/common"
	"github.com/jordan2175/freestix/libstix/defs"
	"time"
)

type CompositeIndicatorExpressionType struct {
	Operator  string         `json:"operator,omitempty"`
	Indicator *IndicatorType `json:"indicator,omitempty"`
}

type RelatedCampaignReferencesType struct {
	Scope            string                              `json:"scope,omitempty"`
	Related_Campaign common.RelatedCampaignReferenceType `json:"relatedCampaign,omitempty"`
}

// Support for IndicatorType version 2.1.1

type IndicatorType struct {
	common.IndicatorBaseType
	Version                      string                               `json:"version,omitempty"`
	Negate                       bool                                 `json:"negate,omitempty"`
	Title                        string                               `json:"title,omitempty"`
	Types                        map[string][]string                  `json:"type,omitempty"`
	AlternativeIDs               []string                             `json:"alternative_ids,omitempty"`
	Descriptions                 []map[string]string                  `json:"descriptions,omitempty"`
	ShortDescriptions            []map[string]string                  `json:"short_descriptions,omitempty"`
	ValidTimePositions           []ValidTimeType                      `json:"valid_time_positions,omitempty"`
	Observable                   *cybox.ObservableType                `json:"observable,omitempty"`
	CompositeIndicatorExpression *CompositeIndicatorExpressionType    `json:"composite_indicator_expression,omitempty"`
	IndicatedTTP                 []common.RelatedTTPType              `json:"indicated_ttps,omitempty"`
	KillChainPhases              []common.KillChainPhaseReferenceType `json:"kill_chain_phases,omitempty"`
	TestMechanisms               []TestMechanismType                  `json:"test_mechanisms,omitempty"`
	LikelyImpact                 *common.StatementType                `json:"likely_impact,omitempty"`
	SuggestedCOAs                []SuggestedCOAsType                  `json:"suggested_coas,omitempty"`
	Handling                     []common.MarkingSpecificationType    `json:"handling,omitempty"`
	Confidence                   *common.ConfidenceType               `json:"confidence,omitempty"`
	Sightings                    *SightingsType                       `json:"sightings,omitempty"`
	RelatedIndicators            *RelatedIndicatorsType               `json:"related_indicators,omitempty"`
	RelatedCampaigns             *RelatedCampaignReferencesType       `json:"related_campaigns,omitempty"`
	RelatedPackages              []common.RelatedPackageRefType       `json:"related_packages,omitempty"`
	Producer                     *common.InformationSourceType        `json:"producer,omitempty"`
}

// TODO need to test this, this may not work

type SuggestedCOAsType struct {
	Scope        string                            `json:"scope,omitempty"`
	SuggestedCOA *common.RelatedCourseOfActionType `json:"suggestedCOA,omitempty"`
}

type TestMechanismType struct {
	Id       string                        `json:"id,omitempty"`
	IdRef    string                        `json:"idref,omitempty"`
	Efficacy *common.StatementType         `json:"efficacy,omitempty"`
	Producer *common.InformationSourceType `json:"producer,omitempty"`
}

type ValidTimeType struct {
	StartTime      string `json:"start_time,omitempty"`
	StartPrecision string `json:"start_precision,omitempty"`
	EndTime        string `json:"end_time,omitempty"`
	EndPrecision   string `json:"end_precision,omitempty"`
}

// ----------------------------------------------------------------------
// Methods IndicatorType
// ----------------------------------------------------------------------

func (this *IndicatorType) CreateId() {
	this.Id = defs.COMPANY + ":indicator-" + uuid.New()
}

func (this *IndicatorType) AddIdRef(idref string) {
	this.IdRef = idref
}

func (this *IndicatorType) CreateTimeStamp() {
	this.Timestamp = time.Now().Format(time.RFC3339)
}

func (this *IndicatorType) AddTimeStamp(t string) {
	// TODO Need to format the string in to ISO 8601 format or check that it is in the right format
	this.Timestamp = t
}

func (this *IndicatorType) AddVersion(ver string) {
	this.Version = ver
}

func (this *IndicatorType) SetNegate(b bool) {
	this.Negate = b
}

func (this *IndicatorType) AddTitle(t string) {
	this.Title = t
}

func (this *IndicatorType) AddType(vocab, value string) {
	if this.Types == nil {
		m := make(map[string][]string)
		this.Types = m
	}

	this.Types[vocab] = append(this.Types[vocab], value)
}

func (this *IndicatorType) AddStandardType(value string) {
	this.AddType(defs.INDICATOR_TYPE_VOCAB, value)
}

func (this *IndicatorType) AddAlternativeID(value string) {
	if this.AlternativeIDs == nil {
		a := make([]string, 0)
		this.AlternativeIDs = a
	}
	this.AlternativeIDs = append(this.AlternativeIDs, value)
}

func (this *IndicatorType) AddDescription(format, value string) {
	m := make(map[string]string)
	m[format] = value

	this.Descriptions = append(this.Descriptions, m)
}

func (this *IndicatorType) AddShortDescription(format, value string) {
	m := make(map[string]string)
	m[format] = value

	this.ShortDescriptions = append(this.ShortDescriptions, m)
}

func (this *IndicatorType) initValidTimePosition() {
	if this.ValidTimePositions == nil {
		a := make([]ValidTimeType, 0)
		this.ValidTimePositions = a
	}
}

func (this *IndicatorType) AddValidTimePosition(start, end string) {
	this.initValidTimePosition()

	tp := ValidTimeType{
		StartTime:      start,
		StartPrecision: "Second",
		EndTime:        end,
		EndPrecision:   "Second",
	}
	this.ValidTimePositions = append(this.ValidTimePositions, tp)
}

// AddObservable adds a relevant cyber observable for this Indicator.
//
// This support 0..1
func (this *IndicatorType) AddObservable(o cybox.ObservableType) {
	this.Observable = &o
}

// TODO Composite_Indicator_Expression

// TODO Indicated_TTP

// AddKillChain adds a relevant kill chain phases indicated by this Indicator.
// This method would require that you set the PhaseId and KillChainId manually.
// See AddKillChainPhaseAndChain()
//
// This support 0..n objects
func (this *IndicatorType) AddKillChain(k common.KillChainPhaseReferenceType) {
	if this.KillChainPhases == nil {
		a := make([]common.KillChainPhaseReferenceType, 0)
		this.KillChainPhases = a
	}
	this.KillChainPhases = append(this.KillChainPhases, k)
}

// AddKillChainPhaseAndChain adds a relevant kill chain phases indicated by this
// Indicator.
//
// This support 0..n objects
func (this *IndicatorType) AddKillChainPhaseAndChain(phase, chain string) {
	if this.KillChainPhases == nil {
		a := make([]common.KillChainPhaseReferenceType, 0)
		this.KillChainPhases = a
	}
	data := common.KillChainPhaseReferenceType{
		PhaseId:     phase,
		KillChainId: chain,
	}
	this.KillChainPhases = append(this.KillChainPhases, data)
}

// TODO Test_Mechanisms

// AddLikelyImpact the likely potential impact within the relevant context if
// this Indicator were to occur. This is typically local to an Indicator consumer
// and not typically shared. This field includes a Description of the likely
// potential impact within the relevant context if this Indicator were to
// occur and a Confidence held in the accuracy of this assertion. NOTE: This
// structure potentially still needs to be fleshed out more for structured
// characterization of impact.
//
// This support 0..1 objects
func (this *IndicatorType) AddLikelyImpact(s common.StatementType) {
	this.LikelyImpact = &s
}

// AddHandling adds the relevant handling guidance for this Indicator. The valid
// marking scope is the nearest IndicatorBaseType ancestor of this Handling
// element and all its descendants.
//
// This supports 0..n objects
func (this *IndicatorType) AddHandling(m common.MarkingSpecificationType) {
	if this.Handling == nil {
		a := make([]common.MarkingSpecificationType, 0)
		this.Handling = a
	}
	this.Handling = append(this.Handling, m)
}

// AddConfidence adds a level of confidence held in the accuracy of this
// Indicator.
//
// This supports 0..1 objects
func (this *IndicatorType) AddConfidence(c common.ConfidenceType) {
	this.Confidence = &c
}

// AddSightings adds a set of sighting reports for this Indicator.
//
// This supports 0..1 objects
func (this *IndicatorType) AddSightings(s SightingsType) {
	this.Sightings = &s
}

// AddRelatedIndicators add an optional related indicator and enables content
// producers to express a relationship between the enclosing indicator (this.e.,
// the subject of the relationship) and a disparate indicator (this.e., the object
// side of the relationship).
//
// This supports 0..1 objects
func (this *IndicatorType) AddRelatedIndicators(r RelatedIndicatorsType) {
	this.RelatedIndicators = &r
}

// TODO Related_Campaigns

// AddRelatedPackages identifies or characterizes relationships to a set of
// related Packages.
//
// This supports 0..n
func (this *IndicatorType) AddRelatedPackage(p common.RelatedPackageRefType) {
	if this.RelatedPackages == nil {
		a := make([]common.RelatedPackageRefType, 0)
		this.RelatedPackages = a
	}
	this.RelatedPackages = append(this.RelatedPackages, p)
}

// AddProducer adds details for the source of this indicator.
//
// This supports 0..1 objects
func (this *IndicatorType) AddProducer(source common.InformationSourceType) {
	this.Producer = &source
}