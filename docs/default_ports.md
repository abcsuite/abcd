While abcd is highly configurable when it comes to the network configuration,
the following is intended to be a quick reference for the default ports used so
port forwarding can be configured as required.

abcd provides a `--upnp` flag which can be used to automatically map the Aero
peer-to-peer listening port if your router supports UPnP.  If your router does
not support UPnP, or you don't wish to use it, please note that only the Aero
peer-to-peer port should be forwarded unless you specifically want to allow RPC
access to your abcd from external sources such as in more advanced network
configurations.

|Name|Port|
|----|----|
|Default Aero peer-to-peer port|TCP 9527|
|Default RPC port|TCP 9528|
