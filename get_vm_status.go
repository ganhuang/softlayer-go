/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"

	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
)

func main() {
	sess := session.New() // default endpoint

	sess.Debug = true

    doExecuteRemoteScriptTest(sess)
	//doError(sess)
}

func doExecuteRemoteScriptTest(sess *session.Session) {
	// Get the VirtualGuest service
	service := services.GetVirtualGuestService(sess)

	// Execute the remote script
	fmt.Println("Error executing remote script on VM:")
	rep, _ := service.Id(79888255).Mask("name").GetPowerState()
    if *rep.Name == "Running" {
        fmt.Println("Status: %s", *rep.Name)
    }
}
