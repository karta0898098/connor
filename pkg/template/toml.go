package template

// AppToml template ...
const AppToml = `
[log]
env = "dev"
app_id = "{{.ProjectName}}"
debug = true
local = true

[database]
    [database.read]
    debug = false
    host = "127.0.0.1"
    user = ""
    password  = ""
    port = 3306
    name = ""
    type = "mysql"

    [database.write]
    debug = false
    host = ""
    user = "rode"
    password  = ""
    port = 3306
    name = ""
    type = "mysql"

[http]
mode = "debug"
port = ":8080"

[grpc]
mode = "debug"
port = ":8081"
`
