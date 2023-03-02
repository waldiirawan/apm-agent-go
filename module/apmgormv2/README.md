# apmgormv2

Package apmgormv2 provides drivers to gorm.io/gorm
for tracing database operations as spans of a transaction traced
by Elastic APM.

## Usage

Swap `gorm.io/driver/*` to `github.com/waldiirawan/apm-agent-go/module/apmgormv2/driver/*`

Example :-

```golang
import (
    mysql "github.com/waldiirawan/apm-agent-go/module/apmgormv2/driver/mysql"
    "gorm.io/gorm"
)	

db, err := gorm.Open(mysql.Open("dsn"), &gorm.Config{})
```