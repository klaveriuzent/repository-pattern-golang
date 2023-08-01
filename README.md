# repository-pattern-golang

need `.env` file setup like
```
# Database credentials
DB_DRIVER="postgres" #You can use: sqlsvr | mysql | postgres
DB_HOST="set_credentials" #Your host location (localhost or IP address)
DB_USER="set_credentials" #Your database username
DB_PASSWORD="set_credentials" #Your database password
DB_NAME="set_credentials" #Your database name
DB_PORT="set_credentials" #Your database port

SERVICE_NAME="repositoryPattern"
# Authentication credentials
TOKEN_TTL="set_credentials" #Default SET "36000"
JWT_PRIVATE_KEY="set_credentials" #JWT private key
API_SECRET="set_credentials" #API secret key
TOKEN_HOUR_LIFESPAN="set_credentials" #Default SET 24 (dont set as string, must be number)
```
