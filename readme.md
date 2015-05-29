# Rivet
API provider for Machine

Rivet provides a thin glue layer between provisioning and Docker Machine.  This
allows for any infrastructure provider to implement Rivet and be immediately
available to Docker Machine for provisioning using the Rivet driver in Machine.

Rivet uses [Pluginhook](https://github.com/progrium/pluginhook) for
provisioning.  You can use any scripts (shell, python, binaries) for the backend
implementation.

# Usage
Rivet is a small Go application that provides a JSON API.  The requests are
sent to pluginhook which then runs the hooks in your custom plugins.  You also
need to build Docker Machine with the
[rivet](https://github.com/ehazlett/machine/tree/driver-rivet) driver.  This
enables Docker Machine to work with any Rivet endpoint and run the custom
provisioning hooks.

# Auth
By default there is no authentication.  However, there is a simple token based
authentication method available.  Specify your auth token with the `--auth-token` flag
and send it in the `X-Auth-Token` header when making requests.

Start Rivet API with Token:

```shell
rivet -p /path/to/plugins --auth-token mysecrettoken
```

Pass header in requests:

```shell
curl -H 'X-Auth-Token:mysecrettoken' http://myhost:8080
```

# Example
See the [Example](https://github.com/ehazlett/rivet/tree/master/rivet/hooks)
plugin for the hook definitions.

> Note: the example plugin does not work with Machine.  It is simply an example
on how to create your own custom plugins.

# Demo
[![Rivet Demo](http://img.youtube.com/vi/kDqW1wEMRw4/0.jpg)](http://www.youtube.com/watch?v=kDqW1wEMRw4)
