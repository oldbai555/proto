server:
  name: "{{.ServerName}}"
  port: 8003
mysql:
  addr: "db"
  port: 3306 # 使用对内的mysql接口
  username: "oldbai"
  password: "123456"
zipkin:
  addr: "http://127.0.0.1:9411/api/v2/spans"
redis:
  database: 0
  host: 'redis' # 使用对内的redis接口
  port: 6379
storage:
  type: ""
  secret_id: ""
  secret_key: ""
  bucket_url: ""
chatGpt:
  api_key: