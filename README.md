# notifier

## What it is

**notifier** is a general purpose notification service to inform you via SMS, email, or telegram. It can be used by other services directly via REST API or indirectly via provided client. Currently support for only a few russian SMS services is available out of the box, but adding your own channel is trivial. Or you can drop me a line as well, and we'll see if support for your favorite SMS service can be added.  

## Features

- web SMS (beeline.ru, smsc.ru, websms.ru)
- telegram bot
- REST API
- command-line client
- easily extensible

## Example setup

### What we need

**notifyd** is written in Go, so you need to have some basic knowledge how to build Go programs. It shold build cleanly accorging to standard Go procedures once you install all the needed Go dependencies. There are not many of them.

**notifyd** is a web-application, hence it can be run on its own in *http* mode. But I'd recommend using it in *https* mode. For that we need a proxy (e.g **nginx**). Maybe sometimes I'll add *https* and other usefull stuff to **notifyd**, but right now we rely on proxy. A local mail server (e.g. **postfix**) is used to send out emails. We also need to get in touch with the **botfather** to register telegram bot.

I assume you know what an ssl sertificate is and how to handle it.

### nginx setup

We'll try to keep things as simple as possible. So here's the **nginx** configuration file *nginx.conf*:
```
events {
    worker_connections  1024;
}

error_log syslog:server=unix:/var/run/log;

http {

    map $ssl_client_s_dn $ssl_client_s_dn_cn {
        default "should_not_happen";
        ~/CN=(?<CN>[^/]+) $CN;
    }

    server {
        listen <--YOUR_IP_HERE-->:80;
        server_name <--SERVER_NAME_HERE-->;
        return 301 https://<--SERVER_NAME_HERE-->$request_uri;
    }

    server {

        access_log syslog:server=unix:/var/run/log;
        listen <--YOUR_IP_HERE-->:443 ssl;
        ssl_certificate     /etc/nginx/<--CERTIFICATE_FILE-->;
        ssl_certificate_key /etc/nginx/<--KEY_FILE-->;
        ssl_protocols       TLSv1 TLSv1.1 TLSv1.2;
        ssl_ciphers         HIGH:!aNULL:!MD5;

        server_name         <--SERVER_NAME_HERE-->;

        location / {
            proxy_pass http://localhost:8084/;
            proxy_set_header   X-Real-IP $remote_addr;
            proxy_set_header   Host $http_host;
            proxy_set_header   X-Forwarded-For $proxy_add_x_forwarded_for;
        }
    }
}
```
**nginx** speaks https to the outside world an communicates with **notifyd** via http on a special port.
It also sets the **X-Forwarded-For** header, just so **notifyd** knows the real address of the calling client.


### postfix setup

For **postfix** this minimal *main.cf* will do:
```
#
# LOCAL PATHNAME INFORMATION
#

sendmail_path = /usr/sbin/sendmail
newaliases_path = /usr/bin/newaliases
mailq_path = /usr/bin/mailq

html_directory = no
manpage_directory = /usr/share/man
readme_directory = no
queue_directory = /var/spool/postfix
command_directory = /usr/sbin
daemon_directory = /usr/libexec/postfix
data_directory = /var/run/postfix

#
# QUEUE AND PROCESS OWNERSHIP
#

mail_owner = postfix
setgid_group = postdrop

#
# NETWORK DETAILS
#

inet_protocols = ipv4
inet_interfaces = localhost

#
# LOCAL DELIVERY
#

# real users

home_mailbox = Maildir/
unknown_local_recipient_reject_code = 550
alias_maps = hash:/etc/postfix/aliases

# virtual users

virtual_mailbox_base=/
virtual_uid_maps=static:1002
virtual_gid_maps=static:1002

#
# TRUST AND RELAY CONTROL
#

#mynetworks = 168.100.189.0/28, 127.0.0.0/8
smtpd_recipient_restrictions = permit_mynetworks, permit_sasl_authenticated, reject_unauth_destination
strict_rfc821_envelopes = yes
```

This will launch **postfix** on 127.0.0.1, just what we need, because we only want so send out.

### telegram bot setup

#### get the token
Contact **BotFather** on telegram, the rest will look like this:
```
What can this bot do?
BotFather is the one bot to rule them all. Use it to create new bot accounts and manage your existing bots.

About Telegram bots:
https://core.telegram.org/bots
Bot API manual:
https://core.telegram.org/bots/api

Contact @BotSupport if you have questions about the Bot API.


**Alex**
/start


**BotFather**
I can help you create and manage Telegram bots. If you're new to the Bot API, please see the manual.

You can control me by sending these commands:

/newbot - create a new bot
...


**Alex**
/newbot


**BotFather**
Alright, a new bot. How are we going to call it? Please choose a name for your bot.
Alex
TelixNotifier
BotFather
Good. Now let's choose a username for your bot. It must end in `bot`. Like this, for example: TetrisBot or tetris_bot.
Alex
telix_notifier_bot
BotFather
Done! Congratulations on your new bot. You will find it at t.me/telix_notifier_bot. You can now add a description, about section and profile picture for your bot, see /help for a list of commands. By the way, when you've finished creating your cool bot, ping our Bot Support if you want a better username for it. Just make sure the bot is fully operational before you do this.

Use this token to access the HTTP API:
<--YOUR_TOKEN_HERE-->

For a description of the Bot API, see this page: https://core.telegram.org/bots/api
```
Store the token somewhere, you'll need to put it in the **notifyd** configuration file later.
#### place your new bot into some group
Create a **telegram group** and add your bot into it. It's ok if your bot "has no access to messages". Now we need the **ChatID** of the group you just added your bot to. Write something to your shiny-new bot (never mind that **notifyd** is not up yet):
```
/test@your_bot_user_name
````
and then go to the following URL:
```
https://api.telegram.org/bot<--YOUR_TOKEN_HERE-->/getUpdates
```
You'll get a JSON. Look up the `"chat":{"id":` part and store the following (negative) number somewhere, we'll need it for the **notifyd** config.

### notifyd setup
**notifyd** uses */etc/notifyd.conf* as default configuration file. The configuration file is in the TOML format:
```
[beeline]
#Your beeline login
login = "foo"
#Your beeline password
password = "bar"
#The addressee will see this string as the name of the message sender.
#The name is usually pre-configured on the operator's side.
sender = "baz"

[websms]
login = "foo"
password = "bar"
sender = "baz"

[smsc]
login = "foo"
password = "bar"
sender = "baz"

[telegram]
token = "<--YOUR_TOKEN_HERE-->"
#You can use almost whatever you like instead of `test` to name the group locally.
#The `_chaiid` suffix is mandatory.
test_chatid = <---->
```
Pretty self-explanatory. **email** has no configuration file sections.

## Invocation

### notifyd
**notifyd** is invoked as follows:
```
Usage:
  notifyd [flags]

Flags:
  -a, --address string   address and port to bind to (default: 127.0.0.1:8084) (default "127.0.0.1:8084")
  -c, --config string    configuration file (default: /etc/notifyd.conf) (default "/etc/notifyd.conf")
  -d, --daemonize        run as a daemon (default: no)
  -h, --help             help for notifyd
  -p, --pidfile string   PID file (default: /var/run/notifyd.pid) (default "/var/run/notifyd.pid")

```

### notify
**notify** is a command-line utility to send messages via **notifyd**. It wats the message to be piped and is invoked as follows:
```
Usage:
  echo "MESSAGE" | notify [beeline|smsc|websms] [flags]

Flags:
  -r, --recipients string   comma-delimited recipient list
  -u, --url string          notifyd url to query

Usage:
  echo "MESSAGE" | notify email [flags]

Flags:
  -r, --recipients string       comma-delimited recipient list
  -a, --sender-address string   sender address
  -n, --sender-name string      sender name
  -s, --subject string          email subject
  -u, --url string              notifyd url to query

Usage:
  notify telegram [flags]

Flags:
  -r, --group string   telegram group name in the configuration file
  -u, --url string     notifyd url to query

```

## API

