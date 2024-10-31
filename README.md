# Gator 🐊

### Table of Contents

- [About](#about)
  - [Commands](#commands)
    - [Users](#users)
    - [Feeds](#feeds)
    - [Feed Follows](#feed-follows)
    - [Posts](#posts)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Database Setup](#database-setup)

## About

This is a guided project from [Boot.dev](https://www.boot.dev/tracks/backend). It's an RSS feed aggregator and has an easy-to-use CLI.
Tools we used:

- PostgreSQL database
- sqlc for generating typesafe Go code from SQL queries
- goose for handling database migrations

### Commands

#### Users

`gator register <name>`

- creates a new user and sets current user

`gator login <name>`

- sets the current user

`gator users`

- lists all users

`gator reset`

- deletes all users from the database

#### Feeds

`gator addfeed <feed_name> <url>`

- adds a new feed to the database and follows it

`gator feeds`

- lists all feeds

#### Feed Follows

`gator following`

- lists all feeds that the current user is following

`gator follow <url>`

- gets a feed from the database and follows it

`gator unfollow <url>`

- deletes a feed from the user's following

#### Posts

`gator browse {<limit>}` limit defaults to 2 if not provided

- displays all posts that the current user is following

`gator agg <time><units>` e.g. 5s, 3m, 1h

- continually fetches feeds every given interval and creates posts in the database

## Getting Started

### Prerequisites

- Go - [official download page](https://go.dev/doc/install)
- Postgres
  - Mac OS with brew:
    ```
    brew install postgresql@16
    ```
  - Linux/WSL:
    ```
    sudo apt update
    sudo apt install postgresql postgresql-contrib
    ```
- Goose
  ```
  go install github.com/pressly/goose/v3/cmd/goose@latest
  ```

### Installation

1. Clone the repository to your local machine:

```
git clone https://github.com/mshortcodes/gator.git
```

2. Change into the project directory:

```
cd gator
```

3. Compile and place the binary into your $GOBIN:

```
go install .
```

### Database setup

1. Start the Postgres server.

Mac:

```
brew services start postgresql
```

Linux:

```
sudo service postgresql start
```

2. Create a new database. You may be prompted to enter a Postgres password if it hasn't been setup.

Mac:

```
psql postgres
```

Linux:

```
sudo -u postgres psql
```

In psql:

```
CREATE DATABASE gator;
```

3. Create a config file. Gator will look in the $HOME environment variable for this:

```
touch ~/.gatorconfig.json
```

The config file should have this structure:

```
{
  "db_url": "<connection_string>",
  "current_user_name": "<username>"
}
```

Replace <connection_string> with your database connection string:

protocol://username:password@host:port/database

The username will be set via the login command and register commands.

4. Run database migrations.

```
cd sql/schema
goose postgres <connection_string> up
```
