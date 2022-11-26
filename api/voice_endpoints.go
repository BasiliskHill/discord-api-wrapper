/*
 * Copyright (c) 2022. Veteran Software
 *
 * Discord API Wrapper - A custom wrapper for the Discord REST API developed for a proprietary project.
 *
 * This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
 * License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
 * warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 */

package api

import (
	"encoding/json"
	"fmt"
	"github.com/veteran-software/discord-api-wrapper/v10/logging"
	"io"
	"net/http"
)

// ListVoiceRegions - Returns an array of voice region objects that can be used when setting a voice or stage channel's rtc_region.
//
//goland:noinspection GoUnusedExportedFunction
func ListVoiceRegions() *[]VoiceRegion {
	resp, err := Rest.Request(http.MethodGet, fmt.Sprintf(listVoiceRegions, api), nil, nil)
	if err != nil {
		logging.Errorln(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var voiceRegions *[]VoiceRegion
	err = json.NewDecoder(resp.Body).Decode(&voiceRegions)
	if err != nil {
		logging.Errorln(err)
		return nil
	}

	return voiceRegions
}