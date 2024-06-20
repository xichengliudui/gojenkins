package gojenkins

import "encoding/xml"

type FlowDefinition struct {
	XMLName          xml.Name            `xml:"flow-definition" json:"XMLName,omitempty"`
	Text             string              `xml:",chardata" json:"Text,omitempty"`
	Plugin           string              `xml:"plugin,attr" json:"Plugin,omitempty"`
	Actions          *Actions            `xml:"actions" json:"Actions,omitempty"`
	Description      string              `xml:"description" json:"Description,omitempty"`
	KeepDependencies string              `xml:"keepDependencies" json:"KeepDependencies,omitempty"`
	Properties       *JobProperties      `xml:"properties" json:"Properties,omitempty"`
	Definition       *PipelineDefinition `xml:"definition" json:"Definition,omitempty"`
	Triggers         string              `xml:"triggers" json:"Triggers,omitempty"`
	Disabled         string              `xml:"disabled" json:"Disabled,omitempty"`
}

type Actions struct {
	Text                                                                                 string                    `xml:",chardata" json:"Text,omitempty"`
	OrgJenkinsciPluginsPipelineModeldefinitionActionsDeclarativeJobAction                *JobAction                `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction" json:"OrgJenkinsciPluginsPipelineModeldefinitionActionsDeclarativeJobAction,omitempty"`
	OrgJenkinsciPluginsPipelineModeldefinitionActionsDeclarativeJobPropertyTrackerAction *JobPropertyTrackerAction `xml:"org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction" json:"OrgJenkinsciPluginsPipelineModeldefinitionActionsDeclarativeJobPropertyTrackerAction,omitempty"`
}

type JobAction struct {
	Text   string `xml:",chardata" json:"Text,omitempty"`
	Plugin string `xml:"plugin,attr" json:"Plugin,omitempty"`
}

type JobPropertyTrackerAction struct {
	Text          string `xml:",chardata" json:"Text,omitempty"`
	Plugin        string `xml:"plugin,attr" json:"Plugin,omitempty"`
	JobProperties string `xml:"jobProperties" json:"JobProperties,omitempty"`
	Triggers      string `xml:"triggers" json:"Triggers,omitempty"`
	Parameters    string `xml:"parameters" json:"Parameters,omitempty"`
	Options       string `xml:"options" json:"Options,omitempty"`
}

type JobProperties struct {
	Text                                                                       string                              `xml:",chardata" json:"Text,omitempty"`
	IoJenkinsPluginsDingTalkJobProperty                                        *DingTalkJobProperty                `xml:"io.jenkins.plugins.DingTalkJobProperty" json:"IoJenkinsPluginsDingTalkJobProperty,omitempty"`
	OrgJenkinsciPluginsWorkflowJobPropertiesDisableConcurrentBuildsJobProperty *DisableConcurrentBuildsJobProperty `xml:"org.jenkinsci.plugins.workflow.job.properties.DisableConcurrentBuildsJobProperty" json:"OrgJenkinsciPluginsWorkflowJobPropertiesDisableConcurrentBuildsJobProperty,omitempty"`
	JenkinsModelBuildDiscarderProperty                                         *BuildDiscarderProperty             `xml:"jenkins.model.BuildDiscarderProperty" json:"JenkinsModelBuildDiscarderProperty,omitempty"`
	HudsonModelParametersDefinitionProperty                                    *ParametersDefinitionProperty       `xml:"hudson.model.ParametersDefinitionProperty" json:"HudsonModelParametersDefinitionProperty,omitempty"`
	OrgJenkinsciPluginsWorkflowJobPropertiesPipelineTriggersJobProperty        *PipelineTriggersJobProperty        `xml:"org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty" json:"OrgJenkinsciPluginsWorkflowJobPropertiesPipelineTriggersJobProperty,omitempty"`
}

type BuildDiscarderProperty struct {
	Text     string    `xml:",chardata" json:"Text,omitempty"`
	Strategy *Strategy `xml:"strategy" json:"Strategy,omitempty"`
}

type Strategy struct {
	Text               string `xml:",chardata" json:"Text,omitempty"`
	Class              string `xml:"class,attr" json:"Class,omitempty"`
	DaysToKeep         string `xml:"daysToKeep" json:"DaysToKeep,omitempty"`
	NumToKeep          string `xml:"numToKeep" json:"NumToKeep,omitempty"`
	ArtifactDaysToKeep string `xml:"artifactDaysToKeep" json:"ArtifactDaysToKeep,omitempty"`
	ArtifactNumToKeep  string `xml:"artifactNumToKeep" json:"ArtifactNumToKeep,omitempty"`
}

type DisableConcurrentBuildsJobProperty struct {
	Text          string         `xml:",chardata" json:"text,omitempty"`
	AbortPrevious *AbortPrevious `xml:"abortPrevious" json:"abortPrevious,omitempty"`
}

type AbortPrevious struct {
	Text string `xml:",chardata" json:"text,omitempty"`
}

type ParametersDefinitionProperty struct {
	Text                 string                `xml:",chardata" json:"Text,omitempty"`
	ParameterDefinitions *ParameterDefinitions `xml:"parameterDefinitions" json:"ParameterDefinitions,omitempty"`
}

type ParameterDefinitions struct {
	Text                                                         string                                                                             `xml:",chardata" json:"Text,omitempty"`
	HudsonModelStringParameterDefinition                         []*StringParameterDefinition                                                       `xml:"hudson.model.StringParameterDefinition" json:"HudsonModelStringParameterDefinition,omitempty"`
	HudsonModelBooleanParameterDefinition                        []*BooleanParameterDefinition                                                      `xml:"hudson.model.BooleanParameterDefinition" json:"HudsonModelBooleanParameterDefinition,omitempty"`
	ComCloudbeesPluginsCredentialsCredentialsParameterDefinition []*CredentialsParameterDefinition                                                  `xml:"com.cloudbees.plugins.credentials.CredentialsParameterDefinition" json:"ComCloudbeesPluginsCredentialsCredentialsParameterDefinition,omitempty"`
	ExtendedChoiceParameter                                      *ComCwctravelHudsonPluginsExtendedChoiceParameterExtendedChoiceParameterDefinition `xml:"com.cwctravel.hudson.plugins.extended__choice__parameter.ExtendedChoiceParameterDefinition"`
}

type PipelineDefinition struct {
	Text    string `xml:",chardata" json:"Text,omitempty"`
	Class   string `xml:"class,attr" json:"Class,omitempty"`
	Plugin  string `xml:"plugin,attr" json:"Plugin,omitempty"`
	Script  string `xml:"script" json:"Script,omitempty"`
	Sandbox string `xml:"sandbox" json:"Sandbox,omitempty"`
}

type PipelineTriggersJobProperty struct {
	Text     string        `xml:",chardata" json:"Text,omitempty"`
	Triggers *PipeTriggers `xml:"triggers" json:"Triggers,omitempty"`
}

type PipeTriggers struct {
	Text                                 string                    `xml:",chardata" json:"Text,omitempty"`
	OrgJenkinsciPluginsGwtGenericTrigger *GwtGenericTrigger        `xml:"org.jenkinsci.plugins.gwt.GenericTrigger" json:"OrgJenkinsciPluginsGwtGenericTrigger,omitempty"`
	ComCloudbeesJenkinsGitHubPushTrigger *JenkinsGitHubPushTrigger `xml:"com.cloudbees.jenkins.GitHubPushTrigger" json:"ComCloudbeesJenkinsGitHubPushTrigger,omitempty"`
}

type GwtGenericTrigger struct {
	Text                      string     `xml:",chardata" json:"Text,omitempty"`
	Plugin                    string     `xml:"plugin,attr" json:"Plugin,omitempty"`
	Spec                      string     `xml:"spec" json:"Spec,omitempty"`
	GenericVariables          *Variables `xml:"genericVariables" json:"GenericVariables,omitempty"`
	RegexpFilterText          string     `xml:"regexpFilterText" json:"RegexpFilterText,omitempty"`
	RegexpFilterExpression    string     `xml:"regexpFilterExpression" json:"RegexpFilterExpression,omitempty"`
	PrintPostContent          string     `xml:"printPostContent" json:"PrintPostContent,omitempty"`
	PrintContributedVariables string     `xml:"printContributedVariables" json:"PrintContributedVariables,omitempty"`
	CauseString               string     `xml:"causeString" json:"CauseString,omitempty"`
	Token                     string     `xml:"token" json:"Token,omitempty"`
	TokenCredentialId         string     `xml:"tokenCredentialId" json:"TokenCredentialId,omitempty"`
	SilentResponse            string     `xml:"silentResponse" json:"SilentResponse,omitempty"`
	OverrideQuietPeriod       string     `xml:"overrideQuietPeriod" json:"OverrideQuietPeriod,omitempty"`
}

type Variables struct {
	Text                                  string                `xml:",chardata" json:"Text,omitempty"`
	OrgJenkinsciPluginsGwtGenericVariable []*GwtGenericVariable `xml:"org.jenkinsci.plugins.gwt.GenericVariable" json:"OrgJenkinsciPluginsGwtGenericVariable,omitempty"`
}

type DingTalkJobProperty struct {
	Text            string                   `xml:",chardata" json:"Text,omitempty"`
	Plugin          string                   `xml:"plugin,attr" json:"Plugin,omitempty"`
	NotifierConfigs *DingTalkNotifierConfigs `xml:"notifierConfigs" json:"NotifierConfigs,omitempty"`
}

type DingTalkNotifierConfigs struct {
	Text                                   string                    `xml:",chardata" json:"Text,omitempty"`
	IoJenkinsPluginsDingTalkNotifierConfig []*DingTalkNotifierConfig `xml:"io.jenkins.plugins.DingTalkNotifierConfig" json:"IoJenkinsPluginsDingTalkNotifierConfig,omitempty"`
}

type DingTalkNotifierConfig struct {
	Text            string                   `xml:",chardata" json:"Text,omitempty"`
	Checked         string                   `xml:"checked" json:"Checked,omitempty"`
	RobotId         string                   `xml:"robotId" json:"RobotId,omitempty"`
	RobotName       string                   `xml:"robotName" json:"RobotName,omitempty"`
	AtAll           string                   `xml:"atAll" json:"AtAll,omitempty"`
	AtMobile        string                   `xml:"atMobile" json:"AtMobile,omitempty"`
	Content         string                   `xml:"content" json:"Content,omitempty"`
	NoticeOccasions *DingTalkNoticeOccasions `xml:"noticeOccasions" json:"NoticeOccasions,omitempty"`
}

type DingTalkNoticeOccasions struct {
	Text   string   `xml:",chardata" json:"Text,omitempty"`
	String []string `xml:"string" json:"String,omitempty"`
}

type GwtGenericVariable struct {
	Text           string `xml:",chardata" json:"Text,omitempty"`
	ExpressionType string `xml:"expressionType" json:"ExpressionType,omitempty"`
	Key            string `xml:"key" json:"Key,omitempty"`
	Value          string `xml:"value" json:"Value,omitempty"`
	RegexpFilter   string `xml:"regexpFilter" json:"RegexpFilter,omitempty"`
	DefaultValue   string `xml:"defaultValue" json:"DefaultValue,omitempty"`
}

type JenkinsGitHubPushTrigger struct {
	Text   string `xml:",chardata" json:"Text,omitempty"`
	Plugin string `xml:"plugin,attr" json:"Plugin,omitempty"`
	Spec   string `xml:"spec" json:"Spec,omitempty"`
}

type ComCwctravelHudsonPluginsExtendedChoiceParameterExtendedChoiceParameterDefinition struct {
	Text                    string `xml:",chardata"`
	Plugin                  string `xml:"plugin,attr"`
	Name                    string `xml:"name"`
	Description             string `xml:"description"`
	QuoteValue              string `xml:"quoteValue"`
	SaveJSONParameterToFile string `xml:"saveJSONParameterToFile"`
	VisibleItemCount        string `xml:"visibleItemCount"`
	Type                    string `xml:"type"`
	Value                   string `xml:"value"`
	MultiSelectDelimiter    string `xml:"multiSelectDelimiter"`
}

type CredentialsParameterDefinition struct {
	Text           string `xml:",chardata" json:"Text,omitempty"`
	Plugin         string `xml:"plugin,attr" json:"Plugin,omitempty"`
	Name           string `xml:"name" json:"Name,omitempty"`
	Description    string `xml:"description" json:"Description,omitempty"`
	DefaultValue   string `xml:"defaultValue" json:"DefaultValue,omitempty"`
	CredentialType string `xml:"credentialType" json:"CredentialType,omitempty"`
	Required       string `xml:"required" json:"Required,omitempty"`
}

type BooleanParameterDefinition struct {
	Text         string `xml:",chardata" json:"Text,omitempty"`
	Name         string `xml:"name" json:"Name,omitempty"`
	Description  string `xml:"description" json:"Description,omitempty"`
	DefaultValue string `xml:"defaultValue" json:"DefaultValue,omitempty"`
}

type StringParameterDefinition struct {
	Text         string `xml:",chardata" json:"Text,omitempty"`
	Name         string `xml:"name" json:"Name,omitempty"`
	Description  string `xml:"description" json:"Description,omitempty"`
	DefaultValue string `xml:"defaultValue" json:"DefaultValue,omitempty"`
	Trim         string `xml:"trim" json:"Trim,omitempty"`
}
