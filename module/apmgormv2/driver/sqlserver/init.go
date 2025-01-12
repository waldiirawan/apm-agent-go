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

// Package apmsqlserver imports the gorm sqlserver dialect package,
// and also registers the sqlserver driver with apmsql.
package apmsqlserver // import "github.com/waldiirawan/apm-agent-go/module/apmgormv2/v2/driver/sqlserver"

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"github.com/waldiirawan/apm-agent-go/module/apmsql/v2"

	_ "github.com/waldiirawan/apm-agent-go/module/apmsql/v2/sqlserver" // register sqlserver with apmsql
)

// Open creates a dialect with apmsql
func Open(dsn string) gorm.Dialector {

	dialect := &sqlserver.Dialector{
		Config: &sqlserver.Config{
			DriverName: apmsql.DriverPrefix + sqlserver.Dialector{}.Name(),
			DSN:        dsn,
		},
	}

	return dialect
}
