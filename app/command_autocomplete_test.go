// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/stretchr/testify/assert"
)

func TestParseStaticListArgument(t *testing.T) {
	fixedArgs := model.NewStaticListArgument()
	fixedArgs.AddArgument("on", "help")

	argument := &model.Argument{
		Name:     "", //positional
		HelpText: "some_help",
		Type:     model.StaticListArgumentType,
		Data:     fixedArgs,
	}
	found, _, suggestions := parseStaticListArgument(argument, "")
	assert.True(t, found)
	assert.Equal(t, []model.Suggestion{{Hint: "on", Description: "help"}}, suggestions)

	found, _, suggestions = parseStaticListArgument(argument, "o")
	assert.True(t, found)
	assert.Equal(t, []model.Suggestion{{Hint: "on", Description: "help"}}, suggestions)

	found, changedInput, _ := parseStaticListArgument(argument, "on ")
	assert.False(t, found)
	assert.Equal(t, "", changedInput)

	found, changedInput, _ = parseStaticListArgument(argument, "on some")
	assert.False(t, found)
	assert.Equal(t, "some", changedInput)

	fixedArgs.AddArgument("off", "help")

	found, _, suggestions = parseStaticListArgument(argument, "o")
	assert.True(t, found)
	assert.Equal(t, []model.Suggestion{{Hint: "on", Description: "help"}, {Hint: "off", Description: "help"}}, suggestions)

	found, _, suggestions = parseStaticListArgument(argument, "of")
	assert.True(t, found)
	assert.Equal(t, []model.Suggestion{{Hint: "off", Description: "help"}}, suggestions)

	found, _, suggestions = parseStaticListArgument(argument, "o some")
	assert.True(t, found)
	assert.Len(t, suggestions, 0)

	found, changedInput, _ = parseStaticListArgument(argument, "off some")
	assert.False(t, found)
	assert.Equal(t, "some", changedInput)

	fixedArgs.AddArgument("onon", "help")

	found, _, suggestions = parseStaticListArgument(argument, "on")
	assert.True(t, found)
	assert.Equal(t, []model.Suggestion{{Hint: "on", Description: "help"}, {Hint: "onon", Description: "help"}}, suggestions)

	found, _, suggestions = parseStaticListArgument(argument, "ono")
	assert.True(t, found)
	assert.Equal(t, []model.Suggestion{{Hint: "onon", Description: "help"}}, suggestions)

	found, changedInput, _ = parseStaticListArgument(argument, "on some")
	assert.False(t, found)
	assert.Equal(t, "some", changedInput)

	found, changedInput, _ = parseStaticListArgument(argument, "onon some")
	assert.False(t, found)
	assert.Equal(t, "some", changedInput)
}

func TestParseInputTextArgument(t *testing.T) {
	argument := &model.Argument{
		Name:     "", //positional
		HelpText: "some_help",
		Type:     model.TextInputArgumentType,
		Data:     &model.TextInputArgument{Hint: "hint", Pattern: "pat"},
	}
	found, _, suggestion := parseInputTextArgument(argument, "")
	assert.True(t, found)
	assert.Equal(t, model.Suggestion{Hint: "hint", Description: "some_help"}, suggestion)

	found, _, suggestion = parseInputTextArgument(argument, " ")
	assert.True(t, found)
	assert.Equal(t, model.Suggestion{Hint: "hint", Description: "some_help"}, suggestion)

	found, _, suggestion = parseInputTextArgument(argument, "abc")
	assert.True(t, found)
	assert.Equal(t, model.Suggestion{Hint: "hint", Description: "some_help"}, suggestion)

	found, _, suggestion = parseInputTextArgument(argument, "\"abc dfd df ")
	assert.True(t, found)
	assert.Equal(t, model.Suggestion{Hint: "hint", Description: "some_help"}, suggestion)

	found, changedInput, _ := parseInputTextArgument(argument, "abc efg ")
	assert.False(t, found)
	assert.Equal(t, "efg ", changedInput)

	found, changedInput, _ = parseInputTextArgument(argument, "abc ")
	assert.False(t, found)
	assert.Equal(t, "", changedInput)

	found, changedInput, _ = parseInputTextArgument(argument, "\"abc def\" abc")
	assert.False(t, found)
	assert.Equal(t, "abc", changedInput)
}

func TestSuggestions(t *testing.T) {
	th := Setup(t).InitBasic()
	defer th.TearDown()

	jira := model.CreateJiraAutocompleteData()
	suggestions := th.App.GetSuggestions([]*model.AutocompleteData{jira}, "ji", model.SYSTEM_ADMIN_ROLE_ID)
	assert.Len(t, suggestions, 1)
	assert.Equal(t, jira.Trigger, suggestions[0].Hint)
	assert.Equal(t, jira.HelpText, suggestions[0].Description)

	suggestions = th.App.GetSuggestions([]*model.AutocompleteData{jira}, "jira crea", model.SYSTEM_ADMIN_ROLE_ID)
	assert.Len(t, suggestions, 1)
	assert.Equal(t, "create", suggestions[0].Hint)
	assert.Equal(t, "Create a new Issue", suggestions[0].Description)

	suggestions = th.App.GetSuggestions([]*model.AutocompleteData{jira}, "jira c", model.SYSTEM_ADMIN_ROLE_ID)
	assert.Len(t, suggestions, 2)
	assert.Equal(t, "create", suggestions[1].Hint)
	assert.Equal(t, "Create a new Issue", suggestions[1].Description)
	assert.Equal(t, "connect", suggestions[0].Hint)
	assert.Equal(t, "Connect your Mattermost account to your Jira account", suggestions[0].Description)

	suggestions = th.App.GetSuggestions([]*model.AutocompleteData{jira}, "jira create ", model.SYSTEM_ADMIN_ROLE_ID)
	assert.Len(t, suggestions, 1)
	assert.Equal(t, "[text]", suggestions[0].Hint)
	assert.Equal(t, "This text is optional, will be inserted into the description field", suggestions[0].Description)

	suggestions = th.App.GetSuggestions([]*model.AutocompleteData{jira}, "jira create some", model.SYSTEM_ADMIN_ROLE_ID)
	assert.Len(t, suggestions, 1)
	assert.Equal(t, "[text]", suggestions[0].Hint)
	assert.Equal(t, "This text is optional, will be inserted into the description field", suggestions[0].Description)

	suggestions = th.App.GetSuggestions([]*model.AutocompleteData{jira}, "jira create some text ", model.SYSTEM_ADMIN_ROLE_ID)
	assert.Len(t, suggestions, 0)

	suggestions = th.App.GetSuggestions([]*model.AutocompleteData{jira}, "invalid command", model.SYSTEM_ADMIN_ROLE_ID)
	assert.Len(t, suggestions, 0)

	suggestions = th.App.GetSuggestions([]*model.AutocompleteData{jira}, "jira settings notifications o", model.SYSTEM_ADMIN_ROLE_ID)
	assert.Len(t, suggestions, 2)
	assert.Equal(t, "on", suggestions[0].Hint)
	assert.Equal(t, "Turn notifications on", suggestions[0].Description)
	assert.Equal(t, "off", suggestions[1].Hint)
	assert.Equal(t, "Turn notifications off", suggestions[1].Description)

	suggestions = th.App.GetSuggestions([]*model.AutocompleteData{jira}, "jira ", model.SYSTEM_ADMIN_ROLE_ID)
	assert.Len(t, suggestions, 10)

	suggestions = th.App.GetSuggestions([]*model.AutocompleteData{jira}, "jira ", model.SYSTEM_USER_ROLE_ID)
	assert.Len(t, suggestions, 8)
}
