# mini-wallet 

# requirement
1. Go 1.16
2. PostgreSQL
3. Redis

# How to install
- Prepare DB
    1. Make sure PostgreSQL is running on your local
    2. Open app.env and replace the `DB_SOURCE` value with your local PostgreSQL credential
    3. Create DB by executing mini_wallet.sql in your local postgreSQL
    4. Inside your psql command line, execute this line `REASSIGN OWNED BY ferdinand to <your_username>;`
- Prepare Redis
    1. Make sure redis is running on your local
    2. Open app.env and replace the `REDIS_HOST` value with your local Redis config (usually it's the same)

# How to run
- go build && ./main-wallet
- mini-wallet API should be running in localhost:7001