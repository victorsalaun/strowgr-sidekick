/*
 *  Copyright (C) 2016 VSCT
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */
package sidekick

import ()

type Daemon struct {
	Properties *Config
}

func (daemon *Daemon) IsMaster() (bool, error) {
	return (daemon.Properties.Status == "master"), nil

	//	resp, err := http.Get(fmt.Sprintf("http://%s:%d/real-ip", daemon.Properties.Vip, daemon.Properties.Port))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer resp.Body.Close()
	//	body, err := ioutil.ReadAll(resp.Body)
	//	log.Printf("master ip: %s", body)
	//	return string(body) == daemon.Properties.IpAddr, nil
}

func (daemon *Daemon) IsSlave() (bool, error) {
	return (daemon.Properties.Status == "slave"), nil
	//	isMaster, err := daemon.IsMaster()
	//	return !isMaster, err
}

func (daemon *Daemon) Is(target string) (bool, error) {
	return (daemon.Properties.Status == target), nil
}

func NewDaemon(properties *Config) *Daemon {
	return &Daemon{
		Properties: properties,
	}
}
