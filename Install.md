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

### Installatie + Configuratie Database
Vraag aan de begeleiders hoe dit moet.

In de database "mariadb" voer de volgende regels uit:
> show databases;
create database golang default character set utf8 default collate utf8_bin;
GRANT ALL PRIVILEGES ON golang.* to golanguser@'%' IDENTIFIED BY 'golang';
GRANT ALL PRIVILEGES ON golang.* to golanguser@'localhost' IDENTIFIED BY 'golang';
USE golang;
CREATE TABLE events(id INT NOT NULL AUTO_INCREMENT, Eventid INT NOT NULL, Eventname VARCHAR(255), PRIMARY KEY(id));
INSERT INTO events(eventid, eventname) VALUES(1, 'Start');

### Testen script
1. Als deze stappen zijn voldaan kan de golang-api alsvolgt uitgevoerd worden:
2. Open een terminal
3. Type "go run ~/GoProjects/Project1/src/golang-api/main.go"
4. Open een internet browser en surf naar "localhost:8080/api"
