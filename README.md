# tgme

Get notifications when commands finish.
I use this for exactly that purpose:

```sh
long running command && tgme "command finished" || tgme "THERE WAS AN ERROR"
```

## Configurations

Use this template please:

```sh
# ~/.config/tgme/config
token    = telegram bot api token
receiver = your account id
```

> You need to have at least 1 message sent to the bot from you in order for
  the bot to reach you.
