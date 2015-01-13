/*
 * MumbleDJ
 * By Matthieu Grieger
 * strings.go
 * Copyright (c) 2014 Matthieu Grieger (MIT License)
 */

package main

// Message shown to users when they do not have permission to execute a command.
const NO_PERMISSION_MSG = "You do not have permission to execute that command."

// Message shown to users when they try to add a playlist to the queue and do not have permission to do so.
const NO_PLAYLIST_PERMISSION_MSG = "You do not have permission to add playlists to the queue."

// Message shown to users when they try to execute a command that doesn't exist.
const COMMAND_DOESNT_EXIST_MSG = "The command you entered does not exist."

// Message shown to users when they try to move the bot to a non-existant channel.
const CHANNEL_DOES_NOT_EXIST_MSG = "The channel you specified does not exist."

// Message shown to users when they attempt to add an invalid URL to the queue.
const INVALID_URL_MSG = "The URL you submitted does not match the required format."

// Message shown to users when they attempt to perform an action on a song when
// no song is playing.
const NO_MUSIC_PLAYING_MSG = "There is no music playing at the moment."

// Message shown to users when they attempt to skip a playlist when there is no playlist playing.
const NO_PLAYLIST_PLAYING_MSG = "There is no playlist playing at the moment."

// Message shown to users when they attempt to use the nextsong command when there is no song coming up.
const NO_SONG_NEXT_MSG = "There are no songs queued at the moment."

// Message shown to users when they issue a command that requires an argument and one was not supplied.
const NO_ARGUMENT_MSG = "The command you issued requires an argument and you did not provide one."

// Message shown to users when they try to change the volume to a value outside the volume range.
const NOT_IN_VOLUME_RANGE_MSG = "Out of range. The volume must be between %f and %f."

// Message shown to user when a successful configuration reload finishes.
const CONFIG_RELOAD_SUCCESS_MSG = "The configuration has been successfully reloaded."

// Message shown to users when an admin skips a song.
const ADMIN_SONG_SKIP_MSG = "An admin has decided to skip the current song."

// Message shown to users when an admin skips a playlist.
const ADMIN_PLAYLIST_SKIP_MSG = "An admin has decided to skip the current playlist."

// Message shown to users when the audio for a video could not be downloaded.
const AUDIO_FAIL_MSG = "The audio download for this video failed. YouTube has likely not generated the audio files for this video yet."

// Message shown to a channel when a new song starts playing.
const NOW_PLAYING_HTML = `
	<table>
		<tr>
			<td align="center"><img src="%s" width=150 /></td>
		</tr>
		<tr>
			<td align="center"><b><a href="http://youtu.be/%s">%s</a> (%s)</b></td>
		</tr>
		<tr>
			<td align="center">Added by %s</td>
		</tr>
	</table>
`

// Message shown to channel when a song is added to the queue by a user.
const SONG_ADDED_HTML = `
	<b>%s</b> has added "%s" to the queue.
`

// Message shown to channel when a playlist is added to the queue by a user.
const PLAYLIST_ADDED_HTML = `
	<b>%s</b> has added the playlist "%s" to the queue.
`

// Message shown to channel when a song has been skipped.
const SONG_SKIPPED_HTML = `
	The number of votes required for a skip has been met. <b>Skipping song!</b>
`

// Message shown to channel when a playlist has been skipped.
const PLAYLIST_SKIPPED_HTML = `
	The number of votes required for a skip has been met. <b>Skipping playlist!</b>
`

// Message shown to display bot commands.
const HELP_HTML = `<br/>
	<b>User Commands:</b>
	<p><b>!help</b> - Displays this help.</p>
	<p><b>!add </b>- Adds songs to queue.</p>
	<p><b>!skip</b> - Casts a vote to skip the current song</p>
	<p><b>!numsongs</b> - Shows how many songs are in queue.</p>
	<p> <b>!skipplaylist</b> - Casts a vote to skip over the current playlist.</p>
	<p style="-qt-paragraph-type:empty"><br/></p>
	<p><b>Admin Commands:</b></p>
	<p><b>!reset</b> - An admin command that resets the song queue. </p>
	<p><b>!forceskip</b> - An admin command that forces a song skip. </p>
	<p><b>!forceskipplaylist</b> - An admin command that forces a playlist skip. </p>
	<p><b>!move </b>- Moves MumbleDJ into channel if it exists.</p>
	<p><b>!reload</b> - Reloads mumbledj.gcfg configuration settings.</p>
	<p><b>!kill</b> - Safely cleans the bot environment and disconnects from the server.</p>
`

// Message shown to users when they ask for the current volume (volume command without argument)
const CUR_VOLUME_HTML = `
	The current volume is <b>%.2f</b>.
`

// Message shown to users when another user votes to skip the current song.
const SKIP_ADDED_HTML = `
	<b>%s</b> has voted to skip the current song.
`

// Message shown to users when another user votes to skip the current playlist.
const PLAYLIST_SKIP_ADDED_HTML = `
	<b>%s</b> has voted to skip the current playlist.
`

// Message shown to users when they successfully change the volume.
const VOLUME_SUCCESS_HTML = `
	<b>%s</b> has changed the volume to <b>%.2f</b>.
`

// Message shown to users when a user successfully resets the SongQueue.
const QUEUE_RESET_HTML = `
	<b>%s</b> has cleared the song queue.
`

// Message shown to users when a user asks how many songs are in the queue.
const NUM_SONGS_HTML = `
	There are currently <b>%d</b> song(s) in the queue.
`

// Message shown to users when they issue the nextsong command.
const NEXT_SONG_HTML = `
	The next song in the queue is "%s", added by <b>%s</b>.
`
