# orbit

`orbit` is a free serverless chat service developed for deployment in AWS.

# modules

| module | description|
| --- | --- |
| [account](./account) | account creation and management |
| [chat](./chat) | getting, posting, or subscribing to a chat |

# terminology

This is super high level terminology for `orbit`.

For communication, a `chat` is the highest level object. They will contain information about the user permissions and
link to the individual `messages`.

A `message` is a simple record in a `chat` containing a timestamp, body, and author.

Finally, `users` and `accounts`. `users` contain mostly information for connecting to other `users`
and `chats`. Whereas `accounts` are more for settings and external links like emails.
