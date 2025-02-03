package cli

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/commands/cli/context"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/manager"
	"os"
)

func Context() contracts.Command {
	return command.Command{
		Name: "context",
		Condition: func(*manager.Manager) bool {
			return true
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				if len(os.Args) < 3 {
					if mgr.Context != nil {
						fmt.Println(mgr.Context.Name)
					} else {
						fmt.Println("no active context found - please add least one context")
					}
				} else {
					switch os.Args[2] {
					case "connect":
						if len(os.Args) > 4 {
							err := context.Connect(os.Args[3], os.Args[4], mgr.Configuration.Environment.ClientDirectory)

							if err != nil {
								fmt.Println(err.Error())
								os.Exit(1)
							}

							fmt.Println("connected to the simplecontainer agent")
						} else {
							fmt.Println("Try this: smr context connect https://API_URL:1443 PATH_TO_CERT.PEM --context NAME_YOU_WANT")
						}
						break
					case "switch":
						contextName := ""
						if len(os.Args) > 3 {
							contextName = os.Args[3]
						}

						context.Switch(contextName, mgr.Context)
						break
					case "export":
						contextName := ""
						if len(os.Args) > 3 {
							contextName = os.Args[3]
						}

						API := ""
						_, err := fmt.Scan(&API)

						if err != nil {
							fmt.Println("failed to read API URL, please specify API url")
							os.Exit(1)
						}

						context.Export(contextName, mgr.Context, mgr.Configuration.Environment.ClientDirectory, API)
					case "import":
						encrypted := ""
						if len(os.Args) > 3 {
							encrypted = os.Args[3]
						}

						key := ""
						_, err := fmt.Scan(&key)

						if err != nil {
							fmt.Println("failed to read decryption key, please specify key in stdin")
							os.Exit(1)
						}

						context.Import(encrypted, mgr.Context, mgr.Configuration.Environment.ClientDirectory, key)
					case "fetch":
						block, _ := pem.Decode(mgr.Context.PrivateKey.Bytes())
						PrivateKeyTmp, err := x509.ParsePKCS8PrivateKey(block.Bytes)

						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}

						PrivateKey := PrivateKeyTmp.(*ecdsa.PrivateKey)

						var bytes []byte
						bytes, err = x509.MarshalPKCS8PrivateKey(PrivateKey)

						if err != nil {
							fmt.Println(err)
							os.Exit(1)
						}

						key := hex.EncodeToString(bytes[:32])
						context.ImportCertificates(mgr.Context, mgr.Configuration.Environment.ClientDirectory, key)
					default:
						fmt.Println("Available commands are: connect, switch")
					}
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
			},
		},
	}
}
