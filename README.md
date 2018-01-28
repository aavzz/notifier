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

### nginx setup

### postfix setup

### telegram bot setup

### notifyd setup

## Invocation

### notifyd

### notify

## API

