ttvbot(8)                   System Manager's Manual                  ttvbot(8)



NAME
       ttvbot - a bot integration for danirod's Twitch chat


SYNOPSIS
       ttvbot


DESCRIPTION
       The ttvbot tool is used to connect to a Twitch chat and expose a set of
       integrations, services and chat replies so  that  it  can  enhance  the
       viewer  experience  of  a  channel  and make the platform safer for the
       viewers.

       When ttvbot is invoked, it will spawn a chatbot that will connect to  a
       designated  Twitch channel. Once connected, it will start listening for
       events happening in the chat and reacting to them. Some of these events
       are:


       •  Join  and  part  events,  to detect on-demand when someone enters or
          leaves the chat, which is used for stuff like the antibot engine  or
          to detect regular viewers.

       •  Incoming  chat  messages,  which  is  mostly used to detect commands
          which the bot may reply to, as a  way  to  provide  common  answers,
          links to social media or just have fun.

       •  Interact  with  the  Twitch  API  via Helix, for things like banning
          users, changing the stream metadata (title, tags...) or changing the
          list of VIP users.


   HTTP server
       To  interact  with the application using an RPC system, the application
       will spawn an HTTP server on boot. This HTTP server can be used to  in‐
       teract  with  the  modules  and to request information from the system,
       such as interact with the Twitch API via proxy.

       The list of endpoints are:


       /health
              Healthcheck endpoint. Returns HTTP 200 if all the system  compo‐
              nents  are OK, and HTTP 400 if at least one of the system compo‐
              nents is not working properly. Also sends as body  the  self-re‐
              port  diagnostic,  where  each component is able to speak by it‐
              self.


ENVIRONMENT
       TTVBOT_USERNAME
              The username of the bot. This is the username  of  the  bot  ac‐
              count,  it  is the one who interacts with the Twitch API, and it
              is the username that will be seen near messages in the chat gen‐
              erated by the bot itself.


       TTVBOT_TARGET
              The  channel  that  the  bot  will join and monitor. This is the
              channel where the bot will be monitoring for  chat  events,  and
              most  of  the time it will be the channel that will be used when
              using the Helix API to interact with Twitch  (such  as  the  one
              where the title would be changed).


       TTVBOT_IRC_TOKEN
              The IRC token used to connect to the Twitch chat using authenti‐
              cation.  This value has the format oauth:XXX, where XXX  is  the
              authentication token.


       TTVBOT_ALLOWED_BOTS
              A comma separated list of usernames that should not be banned by
              the antibot system, even if their usernames come in  the  Twitch
              bots  list.  Usually bots like streamlabs, moobot or own3d would
              be on this list.


       TTVBOT_DATABASE
              The SQLite database used to store persistent data by the  appli‐
              cation and some modules.


       TTVBOT_HTTP_BIND
              The  host  and  port where the HTTP server should be mapped. For
              instance, localhost:6060 would bind the server to localhost  and
              serve  from the port 6060, and 0.0.0.0:4000 will bind the server
              on all interfaces and serve from the port 4000.


TO-DO TASKS
       •  The HTTP server has no authentication mechanism, which means that it
          is  very  dangerous  to bind this to public interfaces that are con‐
          nected to the internet.


AUTHORS
       ttvbot is created and maintained by Dani Rodriguez.


COPYRIGHT
       ttvbot is licensed under the terms of the zlib License. The source code
       is    located    at   <https://git.danirod.es/danirod/ttvbot>   or   at
       <https://github.com/danirod/ttvbot> as a secondary mirror.

       As with any other free software project whose code is available in  the
       wild, you are free to read, study and fork the source code for your own
       needs, but keep in mind that I won't be accepting patches or  pull  re‐
       quests  for  this  project, since it is not a collaborative project but
       rather a personal toy. Thank you for your understanding.



                                                                     ttvbot(8)
