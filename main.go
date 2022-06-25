//Copyright 2022 Maria Petrova marycool674@gmail.com
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package main

import (
	"context"
	"cryptocore/config"
	"cryptocore/infrastructure/grpc"
	"cryptocore/infrastructure/log"
)

func main() {

	cfg := config.InitConfig()

	// New logger
	appLogger, err := log.NewLogger(cfg.Log.Level)
	if err != nil {
		panic(err)
	}

	// TODO: add contexts everywhere
	// Creating the context
	_, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		appLogger.Info("The app has stopped due to the context closure")
	}()

	appLogger.Info("The app is starting...")
	fatalErrCh := make(chan error, 2)

	go func() {

		// Establishing the gRPC connection
		err = grpc.NewGrpcConnection(cfg.Grpc, appLogger, fatalErrCh)
		if err != nil {
			appLogger.Fatal("The app could not establish the gRPC connection. Exiting")
			panic(err)
		}

		appLogger.Info("The app has started")

	}()

	select {
	case err := <-fatalErrCh:
		appLogger.Errorw("Received error from functional unit", "err", err)
	}
}
