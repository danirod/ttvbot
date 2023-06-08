# ttvbot

My Twitch bot.

## Overall design

I am using Fx as a dependency injection framework.

## Modules

Modules are application components. They are facilities that provide functionality.

Each module is in the ./modules directory and is a package that should provide a `Module` variable which is a Fx module, so that the module can be self-packaged and deployed into the main Fx app with a single line of code.

* **botlist**: This module can be used to check by username whether someone is a bot or not.
* **chatbot**: This module connects via IRC to the Twitch chat and expose functions that allow to create event listeners, send messages and all that stuff.
* **database**: This module exposes the local database where data can be persisted.

## Invokes

Application functionaly using the modules.

* **botdetector**: is used to actually maintain monitor the Twitch chat and check for bots. It listen to JOIN events and checks whether someone who enters is a bot or not.
* **roster**: stores Twitch chat events for later replace, such as moderation or testing new bot features without actually having to go live or requesting interactions from the viewers.

## Configuration

Use environment variables.

* TTVBOT_USERNAME: the username of the bot (for instance, mychannel_bot).
* TTVBOT_TARGET: the target channel that the bot has to monitor (for instance, mychannel).
* TTVBOT_IRC_TOKEN: the password for the chat, start with "oauth:".
* TTVBOT_ALLOWED_BOTS: bots that would be marked as bot by the third party bot detection API that we do not want to ban, such as Moobot, Nightbot, Streamlabs...
* TTVBOT_DATABASE: points to the SQLite database used by the application.

## Open source policy for this project

This project is open as in you can read the code and learn.

However, this is not a generic Twitch bot. It is a bot with specific features for my Twitch channel. You probably don't want to just run this code with your own channel. Therefore I am probably not accepting roadmap features or PRs for this repository.