// Copyright (c) 2016-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package commands

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var WebhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Management of webhooks",
}

var WebhookListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List webhooks",
	Long:    "list all webhooks",
	Example: "  webhook list myteam",
	RunE:    listWebhookCmdF,
}

var WebhookCreateIncomingCmd = &cobra.Command{
	Use:     "create-incoming",
	Short:   "Create incoming webhook",
	Long:    "create incoming webhook which allows external posting of messages to specific channel",
	Example: "  webhook create-incoming --channel [channelID] --user [userID] --display-name [displayName] --description [webhookDescription] --lock-to-channel --icon [iconURL]",
	RunE:    createIncomingWebhookCmdF,
}

var WebhookModifyIncomingCmd = &cobra.Command{
	Use:     "modify-incoming",
	Short:   "Modify incoming webhook",
	Long:    "Modify existing incoming webhook by changing its title, description, channel or icon url",
	Example: "  webhook modify-incoming [webhookID] --channel [channelID] --display-name [displayName] --description [webhookDescription] --lock-to-channel --icon [iconURL]",
	RunE:    modifyIncomingWebhookCmdF,
}

var WebhookCreateOutgoingCmd = &cobra.Command{
	Use:   "create-outgoing",
	Short: "Create outgoing webhook",
	Long:  "create outgoing webhook which allows external posting of messages from a specific channel",
	Example: `  webhook create-outgoing --team myteam --user myusername --display-name mywebhook --trigger-words "build\ntest" --urls http://localhost:8000/my-webhook-handler
	webhook create-outgoing --team myteam --channel mychannel --user myusername --display-name mywebhook --description "My cool webhook" --trigger-when 1 --trigger-words "build\ntest" --icon http://localhost:8000/my-slash-handler-bot-icon.png --urls http://localhost:8000/my-webhook-handler --content-type "application/json"`,
	RunE: createOutgoingWebhookCmdF,
}

func listWebhookCmdF(command *cobra.Command, args []string) error {
	app, err := InitDBCommandContextCobra(command)
	if err != nil {
		return err
	}
	defer app.Shutdown()

	var teams []*model.Team
	if len(args) < 1 {
		var getErr *model.AppError
		// If no team is specified, list all teams
		teams, getErr = app.GetAllTeams()
		if getErr != nil {
			return getErr
		}
	} else {
		teams = getTeamsFromTeamArgs(app, args)
	}

	for i, team := range teams {
		if team == nil {
			CommandPrintErrorln("Unable to find team '" + args[i] + "'")
			continue
		}

		// Fetch all hooks with a very large limit so we get them all.
		incomingResult := app.Srv.Store.Webhook().GetIncomingByTeam(team.Id, 0, 100000000)
		outgoingResult := app.Srv.Store.Webhook().GetOutgoingByTeam(team.Id, 0, 100000000)

		if result := <-incomingResult; result.Err == nil {
			CommandPrettyPrintln(fmt.Sprintf("Incoming webhooks for %s (%s):", team.DisplayName, team.Name))
			hooks := result.Data.([]*model.IncomingWebhook)
			for _, hook := range hooks {
				CommandPrettyPrintln("\t" + hook.DisplayName + " (" + hook.Id + ")")
			}
		} else {
			CommandPrintErrorln("Unable to list incoming webhooks for '" + args[i] + "'")
		}

		if result := <-outgoingResult; result.Err == nil {
			hooks := result.Data.([]*model.OutgoingWebhook)
			CommandPrettyPrintln(fmt.Sprintf("Outgoing webhooks for %s (%s):", team.DisplayName, team.Name))
			for _, hook := range hooks {
				CommandPrettyPrintln("\t" + hook.DisplayName + " (" + hook.Id + ")")
			}
		} else {
			CommandPrintErrorln("Unable to list outgoing webhooks for '" + args[i] + "'")
		}
	}
	return nil
}

func createIncomingWebhookCmdF(command *cobra.Command, args []string) error {
	app, err := InitDBCommandContextCobra(command)
	if err != nil {
		return err
	}
	defer app.Shutdown()

	channelArg, _ := command.Flags().GetString("channel")
	channel := getChannelFromChannelArg(app, channelArg)
	if channel == nil {
		return errors.New("Unable to find channel '" + channelArg + "'")
	}

	userArg, _ := command.Flags().GetString("user")
	user := getUserFromUserArg(app, userArg)
	displayName, _ := command.Flags().GetString("display-name")
	description, _ := command.Flags().GetString("description")
	iconUrl, _ := command.Flags().GetString("icon")
	channelLocked, _ := command.Flags().GetBool("lock-to-channel")

	incomingWebhook := &model.IncomingWebhook{
		ChannelId:     channel.Id,
		DisplayName:   displayName,
		Description:   description,
		IconURL:       iconUrl,
		ChannelLocked: channelLocked,
	}

	if _, err := app.CreateIncomingWebhookForChannel(user.Id, channel, incomingWebhook); err != nil {
		return err
	}

	return nil
}

func modifyIncomingWebhookCmdF(command *cobra.Command, args []string) error {
	app, err := InitDBCommandContextCobra(command)
	if err != nil {
		return err
	}
	defer app.Shutdown()

	if len(args) < 1 {
		return errors.New("WebhookID is not specified")
	}

	webhookArg := args[0]
	oldHook, getErr := app.GetIncomingWebhook(webhookArg)
	if getErr != nil {
		return errors.New("Unable to find webhook '" + webhookArg + "'")
	}

	updatedHook := oldHook

	channelArg, _ := command.Flags().GetString("channel")
	if channelArg != "" {
		channel := getChannelFromChannelArg(app, channelArg)
		if channel == nil {
			return errors.New("Unable to find channel '" + channelArg + "'")
		}
		updatedHook.ChannelId = channel.Id
	}

	displayName, _ := command.Flags().GetString("display-name")
	if displayName != "" {
		updatedHook.DisplayName = displayName
	}
	description, _ := command.Flags().GetString("description")
	if description != "" {
		updatedHook.Description = description
	}
	iconUrl, _ := command.Flags().GetString("icon")
	if iconUrl != "" {
		updatedHook.IconURL = iconUrl
	}
	channelLocked, _ := command.Flags().GetBool("lock-to-channel")
	updatedHook.ChannelLocked = channelLocked

	if _, err := app.UpdateIncomingWebhook(oldHook, updatedHook); err != nil {
		return err
	}

	return nil
}

func createOutgoingWebhookCmdF(command *cobra.Command, args []string) error {
	app, err := InitDBCommandContextCobra(command)
	if err != nil {
		return err
	}
	defer app.Shutdown()

	teamArg, errTeam := command.Flags().GetString("team")
	if errTeam != nil || teamArg == "" {
		return errors.New("Team is required")
	}
	team := getTeamFromTeamArg(app, teamArg)
	if team == nil {
		return errors.New("Unable to find team: " + teamArg)
	}

	userArg, errUser := command.Flags().GetString("user")
	if errUser != nil || userArg == "" {
		return errors.New("User is required")
	}
	user := getUserFromUserArg(app, userArg)
	if user == nil {
		return errors.New("Unable to find user: " + userArg)
	}

	displayName, errName := command.Flags().GetString("display-name")
	if errName != nil || displayName == "" {
		return errors.New("Display name is required")
	}

	triggerWordsString, errWords := command.Flags().GetString("trigger-words")
	if errWords != nil || triggerWordsString == "" {
		return errors.New("Trigger word or words required")
	}
	triggerWords := strings.Split(triggerWordsString, "\n")

	callbackURLsString, errURL := command.Flags().GetString("urls")
	if errURL != nil || callbackURLsString == "" {
		return errors.New("Callback URL or URLs required")
	}
	callbackURLs := strings.Split(callbackURLsString, "\n")

	triggerWhen, _ := command.Flags().GetInt("trigger-when")
	description, _ := command.Flags().GetString("description")
	contentType, _ := command.Flags().GetString("content-type")
	iconURL, _ := command.Flags().GetString("icon")

	outgoingWebhook := &model.OutgoingWebhook{
		CreatorId:    user.Id,
		Username:     user.Username,
		TeamId:       team.Id,
		TriggerWords: triggerWords,
		TriggerWhen:  triggerWhen,
		CallbackURLs: callbackURLs,
		DisplayName:  displayName,
		Description:  description,
		ContentType:  contentType,
		IconURL:      iconURL,
	}

	channelArg, _ := command.Flags().GetString("channel")
	if channelArg != "" {
		channel := getChannelFromChannelArg(app, channelArg)
		if channel != nil {
			outgoingWebhook.ChannelId = channel.Id
		}
	}

	if _, err := app.CreateOutgoingWebhook(outgoingWebhook); err != nil {
		return err
	}

	return nil
}

func init() {
	WebhookCreateIncomingCmd.Flags().String("channel", "", "Channel ID")
	WebhookCreateIncomingCmd.Flags().String("user", "", "User ID")
	WebhookCreateIncomingCmd.Flags().String("display-name", "", "Incoming webhook display name")
	WebhookCreateIncomingCmd.Flags().String("description", "", "Incoming webhook description")
	WebhookCreateIncomingCmd.Flags().String("icon", "", "Icon URL")
	WebhookCreateIncomingCmd.Flags().Bool("lock-to-channel", false, "Lock to channel")

	WebhookModifyIncomingCmd.Flags().String("channel", "", "Channel ID")
	WebhookModifyIncomingCmd.Flags().String("display-name", "", "Incoming webhook display name")
	WebhookModifyIncomingCmd.Flags().String("description", "", "Incoming webhook description")
	WebhookModifyIncomingCmd.Flags().String("icon", "", "Icon URL")
	WebhookModifyIncomingCmd.Flags().Bool("lock-to-channel", false, "Lock to channel")

	WebhookCreateOutgoingCmd.Flags().String("team", "", "Team name or ID (required)")
	WebhookCreateOutgoingCmd.Flags().String("channel", "", "Channel name or ID")
	WebhookCreateOutgoingCmd.Flags().String("user", "", "User username, email, or ID (required)")
	WebhookCreateOutgoingCmd.Flags().String("display-name", "", "Outgoing webhook display name (required)")
	WebhookCreateOutgoingCmd.Flags().String("description", "", "Outgoing webhook description")
	WebhookCreateOutgoingCmd.Flags().String("trigger-words", "", "Words to trigger webhook (word1\nword2) (required)")
	WebhookCreateOutgoingCmd.Flags().Int("trigger-when", 0, "When to trigger webhook (either when trigger word is first (enter 1) or when it's anywhere (enter 0))")
	WebhookCreateOutgoingCmd.Flags().String("icon", "", "Icon URL")
	WebhookCreateOutgoingCmd.Flags().String("urls", "", "Callback URLs (url1\nurl2) (required)")
	WebhookCreateOutgoingCmd.Flags().String("content-type", "", "Content-type")

	WebhookCmd.AddCommand(
		WebhookListCmd,
		WebhookCreateIncomingCmd,
		WebhookModifyIncomingCmd,
		WebhookCreateOutgoingCmd,
	)

	RootCmd.AddCommand(WebhookCmd)
}
