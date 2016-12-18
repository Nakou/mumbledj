/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/add_secret.go
 * Copyright (c) 2016 Captain Nakou (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/interfaces"
	"github.com/spf13/viper"
)

// AddSecretCommand is a command that adds an audio track associated with a supported
// URL to the queue.
type AddSecretCommand struct{}

// Aliases returns the current aliases for the command.
func (c *AddSecretCommand) Aliases() []string {
	return viper.GetStringSlice("commands.addsecret.aliases")
}

// Description returns the description for the command.
func (c *AddSecretCommand) Description() string {
	return viper.GetString("commands.addsecret.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *AddSecretCommand) IsAdminCommand() bool {
	return viper.GetBool("commands.addsecret.is_admin")
}

// Execute executes the command with the given user and arguments.
// Return value descriptions:
//    string: A message to be returned to the user upon successful execution.
//    bool:   Whether the message should be private or not. true = private,
//            false = public (sent to whole channel).
//    error:  An error message to be returned upon unsuccessful execution.
//            If no error has occurred, pass nil instead.
// Example return statement:
//    return "This is a private message!", true, nil
func (c *AddSecretCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	var (
		allTracks      []interfaces.Track
		tracks         []interfaces.Track
		service        interfaces.Service
		err            error
		lastTrackAdded interfaces.Track
	)

	if len(args) == 0 {
		return "", true, errors.New(viper.GetString("commands.addsecret.messages.no_url_error"))
	}

	for _, arg := range args {
		if service, err = DJ.GetService(arg); err == nil {
			tracks, err = service.GetTracks(arg, user)
			if err == nil {
				allTracks = append(allTracks, tracks...)
			}
		}
	}

	if len(allTracks) == 0 {
		return "", true, errors.New(viper.GetString("commands.addsecret.messages.no_valid_tracks_error"))
	}

	numTooLong := 0
	numAdded := 0
	for _, track := range allTracks {
		if err = DJ.Queue.AppendTrack(track); err != nil {
			numTooLong++
		} else {
			numAdded++
			lastTrackAdded = track
		}
	}

	if numAdded == 0 {
		return "", true, errors.New(viper.GetString("commands.addsecret.messages.tracks_too_long_error"))
	} else if numAdded == 1 {
		return fmt.Sprintf(viper.GetString("commands.addsecret.messages.one_track_added"),
			user.Name, lastTrackAdded.GetTitle(), lastTrackAdded.GetService()), false, nil
	}

	retString := fmt.Sprintf(viper.GetString("commands.addsecret.messages.one_track_added_secretly"), user.Name, numAdded)
	if numTooLong != 0 {
		retString += fmt.Sprintf(viper.GetString("commands.addsecret.messages.num_tracks_too_long"), numTooLong)
	}
	return retString, false, nil
}
