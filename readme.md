# Rivet
API provider for Machine

Rivet provides a thin glue layer between provisioning and Docker Machine.  This
allows for any infrastructure provider to implement Rivet and be immediately
available to Docker Machine for provisioning using the Rivet provider in Machine.

Rivet uses [Pluginhook](https://github.com/progrium/pluginhook) for
provisioning.  You can use any scripts (shell, python, binaries) for the backend
implementation.
