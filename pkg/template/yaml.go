package template

// AppYaml template ...
const AppYaml = `log: 
  app_id: "{{.ProjectName}}"
  debug: true
  env: dev
  local: true
database: 
  read: 
    debug: false
    host: "127.0.0.1"
    name: ""
    password: ""
    port: 3306
    type: mysql
    user: ""
  write: 
    debug: false
    host: ""
    name: ""
    password: ""
    port: 3306
    type: mysql
    user: ""
grpc: 
  mode: debug
  port: ":8081"
http: 
  mode: debug
  port: ":8080"
`
