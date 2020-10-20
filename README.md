# DiscordClocks
Clocks for your Discord servers. Updates every five minutes due to API restrictions.

![Screenshot](https://i.imgur.com/HaCt4I4.png)

## Installation
Make sure you have a bot token for the Discord API.

Download the program and Go v1.15+, install the dependencies, and run `go build` to create the executable.  Alternatively, pick up the (possibly outdated, check the commit hash) pre-built Windows binary from [here](https://github.com/karashiiro/DiscordClocks/releases/latest).

Set the environment variable `DISCLOCKS_BOT_TOKEN` to your bot token.

## Configuration
Set `mod_roles` in the configuration file generated on the first run to an array of strings including role IDs authorized
to modify server clocks.

## Usage
`^addclock <channel ID> <tz timezone> [custom abbreviation]`: Creates a clock on voice channel `<channel ID>`. You can get
channel IDs by enabling Developer Mode in `Settings->Appearance` and then right-clicking on a channel to "Copy ID".

This should only be used on voice channels, since they can have spaces and capital letters in their names.

Example: `^addclock 637737139022462987 America/Los_Angeles`

This creates a clock on channel `637737139022462987` with a time in PST or PDT.

Example: `^addclock 637737139022462987 America/Los_Angeles PDT`

This creates a clock on channel `637737139022462987` with a time in PST or PDT, but it will have the label PDT no
matter what, even if it is incorrect.

`^removeclock <channel ID>`: Removes a clock from a voice channel.
