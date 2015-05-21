# Hooks
Rivet uses [Pluginhook](https://github.com/progium/pluginhook).  The following
hooks are used with Pluginhook along with Machine:

- `create`: used during Machine creation
- `get_ip`: used to retrieve the instance IP
- `get_state`: used to retrieve the instance state 
- `kill`: used to kill the Machine
- `remove`: used during Machine removal
- `restart`: used to restart a Machine
- `start`: used to start a Machine
- `stop`: used to stop a Machine
