![simplecontainer manager](.github/resources/repository.jpg)

# Simplecontainer client
> [!WARNING]
> The project is not stable yet. Use it on your own responsibility.

Simplecontainer is based on the server-client architecture. You can interact with the remote/local client via mTLS.

This repository is holding client for the simplecontainer daemon.

Installation of the client
--------------------------

Client CLI is used for communication to the simplecontainer over network using mTLS.
It is secured by mutual verification and encryption.

To install client just download it from releases:

https://github.com/simplecontainer/client/releases

Example for installing latest version:

```bash
export VERSION=$(curl -s https://raw.githubusercontent.com/simplecontainer/client/main/version)
export PLATFORM=linux-amd64
curl -Lo client https://github.com/simplecontainer/client/releases/download/$VERSION/client-$PLATFORM
chmod +x client
sudo mv client /usr/local/bin/smr
smr context connect https://localhost:1443 $HOME/.ssh/simplecontainer/root.pem --context localhost
{"level":"info","ts":1720694421.2032707,"caller":"context/Connect.go:40","msg":"authenticated against the smr-agent"}
smr ps
GROUP  NAME  DOCKER NAME  IMAGE  IP  PORTS  DEPS  DOCKER STATE  SMR STATE
```
Afterward access to control plane of the simple container is configured.
# License
This project is licensed under the GNU General Public License v3.0. See more in LICENSE file.