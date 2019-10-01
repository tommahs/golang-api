# golang-api
API geschreven in de programmeertaal Go met gebruik van de standaard libraries aangeleverd door Go.
Auteur: Tom van Hamersveld


## Introductie
### Functionele werking API:
![alt text](https://raw.githubusercontent.com/tommahs/golang-api/master/functioneel-api.png?token=AFMV6OLIHODYLHMQCXXUIUC5TR7AS)
### Uitleg:
1. User “surft” naar de URL (adres)
    Localhost:8080/api/db/events/read/1
2. Browser stuurt door naar Operating
    System
3. Operating System ziet localhost dus stuurt
    door naar API op de eigen PC
4. API ontvangt URL en leest
    /api/db/events/read/1
5. API voert logica uit bij /api/db
   In dit geval, "lees" informatie uit database
   tabel "events" waarbij uniek id "1" is uit
   database
6. API vraagt aan de database
7. Database probeert uniek 1 uit events te lezen
8. Database stuurt resultaat naar API
9. API stuurt resultaat naar Operating System
10. Operating System stuurt resultaat naar Internet Browser
11. Internet Browser toont resultaat op het scherm

## Setup
In de setup moeten er mappen aangemaakt worden en genavigeerd worden naar verschillende mappen.
De volgende commando's kunnen hiervoor ingezet worden: "mkdir" "cd" en "pwd"
Stel je wilt een map aanmaken in ~/GoProjects/Project1:
> "mkdir ~/GoProjects/Project1/src"

Om van map te wisselen kan het volgende commando gebruikt worden:
>"cd ~/GoProjects/Project1/src"

of als je omhoog wilt,
> "cd .."

Om te kijken in waar in de mappenstructuur je bent:
> "pwd"

### Installatie golang-api
Surf naar "https://github.com/tommahs/golang-api"
Druk rechtsbovenin op "Clone or download"
Download de zip variant
Pak de zip uit en plaats de map "golang-api" in de map "~/GoProjects/Project1/src"

### Installatie SQL library
Surf naar "https://github.com/go-sql-driver/mysql"
Druk rechtsbovenin op "Clone or download"
Download de zip variant
Pak de zip uit en plaats de bestanden in  map "mysql" in de map "~/GoProjects/Project1/src/github.com/go-sql-driver"

### Configuratie Database
In de database "mariadb" voer de volgende regels uit:
> show databases;
create database golang default character set utf8 default collate utf8_bin;
GRANT ALL PRIVILEGES ON golang.* to golanguser@'%' IDENTIFIED BY 'golang';
GRANT ALL PRIVILEGES ON golang.* to golanguser@'localhost' IDENTIFIED BY 'golang';
USE golang;
CREATE TABLE events(id INT NOT NULL AUTO_INCREMENT, Eventid INT NOT NULL, Eventname VARCHAR(255), PRIMARY KEY(id));
INSERT INTO events(eventid, eventname) VALUES(1, 'Start');

### Testen script
1. Als deze stappen zijn voldaan kan het script alsvolgt uitgevoerd worden:
2. Open een terminal
3. Type "go run ~/GoProjects/Project1/src/golang-api/main.go"
4. Open een internet browser en surf naar "localhost:8080/api"
