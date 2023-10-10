// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package main

import (
	"crypto/tls"
	"github.com/mattermost/mattermost-server/v6/plugin"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	plugin.ClientMain(&Plugin{})
}
