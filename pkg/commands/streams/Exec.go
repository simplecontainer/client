package streams

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/gorilla/websocket"
	"github.com/simplecontainer/client/pkg/command"
	"github.com/simplecontainer/client/pkg/contracts"
	"github.com/simplecontainer/client/pkg/helpers"
	"github.com/simplecontainer/client/pkg/logger"
	"github.com/simplecontainer/client/pkg/manager"
	"github.com/simplecontainer/smr/pkg/kinds/containers/platforms/types"
	"github.com/simplecontainer/smr/pkg/network/wss"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/term"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func Exec() contracts.Command {
	return command.Command{
		Name: "exec",
		Condition: func(mgr *manager.Manager) bool {
			return mgr.Context.ConnectionTest()
		},
		Functions: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				format, err := helpers.BuildFormat(helpers.GrabArg(2), mgr.Configuration.Startup.G)
				if err != nil {
					logger.Log.Error("error building format:", zap.Error(err))
					os.Exit(1)
				}

				url := fmt.Sprintf("%s/api/v1/exec/%s/%s/%s", mgr.Context.ApiURL, format.ToString(), strconv.FormatBool(viper.GetBool("it")), viper.GetString("c"))

				conn, err := wss.Request(mgr.Context.Client, url)

				if err != nil {
					logger.Log.Error("error connecting to WebSocket", zap.Error(err))
					os.Exit(1)
				}

				defer conn.Close()

				var done = make(chan bool)

				if !viper.GetBool("it") {
					go func() {
						for {
							t, msg, err := conn.ReadMessage()

							if err != nil {
								logger.Log.Error("error reading WebSocket message", zap.Error(err))
								break
							}

							if t == websocket.CloseMessage {
								fmt.Println(string(msg))
							} else {
								var result types.ExecResult

								if err := json.Unmarshal(msg, &result); err != nil {
									fmt.Println(err)
									os.Exit(1)
								}

								if result.Exit == 0 {
									fmt.Print(result.Stdout)
								} else {
									fmt.Print(result.Stderr)
								}

								done <- true
								return
							}
						}
					}()

					select {
					case <-done:
						return
					}
				} else {
					fd := os.Stdin.Fd()
					oldState, err := term.MakeRaw(int(fd))

					if err != nil {
						log.Fatalf("Error setting terminal to raw mode: %v", err)
					}

					defer term.Restore(int(fd), oldState)

					if terminal.IsTerminal(int(fd)) {
						go func() {
							for {
								consoleReader := bufio.NewReaderSize(os.Stdin, 1)
								input, _ := consoleReader.ReadByte() // Ctrl-C = 3
								if input == 3 {
									done <- true
									return
								}

								conn.WriteMessage(websocket.BinaryMessage, []byte{input})
							}
						}()
					}

					go func() {
						for {
							_, msg, err := conn.ReadMessage()

							if err != nil {
								done <- true
								return
							}

							var outBuf, errBuf bytes.Buffer
							outputDone := make(chan error)

							go func() {
								_, err = stdcopy.StdCopy(&outBuf, &errBuf, bytes.NewBuffer(msg))
								outputDone <- err
							}()

							select {
							case err := <-outputDone:
								if err != nil {
									return
								}
								break
							}

							stdout, err := ioutil.ReadAll(&outBuf)
							if err != nil {
								return
							}
							stderr, err := ioutil.ReadAll(&errBuf)
							if err != nil {
								return
							}

							_, err = io.Copy(os.Stdout, bytes.NewReader(stdout))

							if err != nil {
								fmt.Println(err)
								return
							}

							_, err = io.Copy(os.Stderr, bytes.NewReader(stderr))

							if err != nil {
								fmt.Println(err)
								return
							}
						}
					}()

					select {
					case <-done:
						return
					}
				}
			},
		},
		DependsOn: []func(*manager.Manager, []string){
			func(mgr *manager.Manager, args []string) {
				// Additional dependencies if necessary
			},
		},
	}
}
