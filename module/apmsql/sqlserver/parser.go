// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package apmsqlserver // import "github.com/waldiirawan/apm-agent-go/module/apmsql/v2/sqlserver"

import (
	"github.com/denisenkom/go-mssqldb/msdsn"

	"github.com/waldiirawan/apm-agent-go/module/apmsql/v2"
)

// ParseDSN parses the given denisenkom/go-mssqldb datasource name.
func ParseDSN(name string) apmsql.DSNInfo {
	cfg, _, err := msdsn.Parse(name)
	if err != nil {
		// sqlserver.Open will fail with the same error,
		// so just return a zero value.
		return apmsql.DSNInfo{}
	}
	return apmsql.DSNInfo{
		Address:  cfg.Host,
		Port:     int(cfg.Port),
		Database: cfg.Database,
		User:     cfg.User,
	}
}
