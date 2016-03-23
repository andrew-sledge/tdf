# tdf

Like tail -f, but for the really lazy.

Tails the log of the file specified, prints to a websocket.

```
tdf /var/log/nginx.log
```

Then connect to your host. Example index.html included, but you'll need to update the websocket server. By default listens on port 8080.
