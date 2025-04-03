# go-miniserver

![Hero Image](./artifacts/general/img/hero.jpg)

A simple go server build with the `net/http` package

## Quick Links
| Link | Description | Credential |
|------|-------------|------------|
|http://localhost:8080 | Local development | - |


## Development
### Overview and local setup
This Go project is structured as follows:

```
├── go-miniserver
│   ├── bin
│   │   └── *                     # Compiled binary goes here after build
│   ├── cmd
│   │   └── api
│   │        ├── main.go          # Entry point of the application
│   │        ├── router.go        # Router setup using net/http
│   │        └── middleware.go    # Custom middleware definitions
│   ├── internal
│   │   └── handlers
│   │        ├── handlers.go      # Route handler wiring
│   │        └── healthcheck.go   # Healthcheck endpoint logic
│   └── .env                      # Environment variables for local dev
```

**Note:** 
This project uses the standard Go `net/http` package for HTTP routing.
All routes and middleware are defined in `/cmd/api`, and actual handler logic is encapsulated in `/internal/handlers`.

### First time init
To develop locally, the *./src/node_modules* and the database must be available. Please run the following command before starting the container for the first time:

```
# Always start in the root of repository project
cd ./

# Start the container
docker-compose up

# Switch to ./src directory
cd src/

# Install the node_modules
npm install
```

### Docker & NodeJS commands for development

```
# Always start in the root of repository project
cd ./

# Edit .env.sample and change name to .env

# Start projects 
docker-compose up

# Remove the container
docker-compose down

# Start GULP watcher for build (JS, SASS/LESS) (in new terminal)
$ cd src/; npm run build;
```

## Bundle & Deployment Commands
### Local Maschine

```
# Always start in the root of repository project
cd ./

# Create release bundle (*./dist) [on initial build use flag: -- --init=true]
$ cd src/; npm run bundle 

# Start deployment process (*./dist) [on initial deployment use flag: -- --init=true]
$ cd src/; npm run deploy
```

## Firewall Settings PLESK
### Production

```
1. Open Settings
Tools & Einstellungen

2. Open Firewall settings
Firewall

3. Activate Firewall

4. Add new rule
Benutzerdefinierte Regel hinzufügen

5. Name, Port and Remote Address
Name: MongoDB
Ports: 27017
Remote Address: local maschine IP (Telekom)

6. Änderung aktivieren

7. Aktivieren
```

## MongoDB Settings (if not running Docker and DB in container)
### MongoDB Add Admin User

```
1. Open DB
$ mongo

2. Use Admin Table
> use admin

3. Create Admin User
> db.createUser({ user: "hassani", pwd: passwordPrompt(), roles: [ { role: "userAdminAnyDatabase", db: "admin" }, "readWriteAnyDatabase" ] })

4. Set Password

5. Exit and change config
> exit
```

### MongoDB Enable Authentication

```
1. Open config File
$ sudo nano /etc/mongod.conf

2. Scroll down to find the commented-out security section
security:

3. Then add the authorization parameter and set it to enabled
security:
  authorization: enabled

4. Save and exit
STRG + x

5. Restart mongodb
$ sudo systemctl restart mongod

6. Check service status
$ sudo systemctl status mongod

7. Login now with credentials:
$ mongo -u hassani -p
```

### MongoDB remote Access

```
! Dont forget first PLESK Firewall Settings
1. Open config file
$ sudo nano /etc/mongod.conf

2. Change bindIP (127.0.0.1):
# network interfaces
net:
  port: 27017
  bindIp: 0.0.0.0

3. Restart mongodb
$ sudo systemctl restart mongod

4. Check service status
$ sudo systemctl status mongod

5. Login with credentials
mongodb://username:password@localhost:27017
```


