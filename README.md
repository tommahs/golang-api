# Golang-api
## Introductie
Voor de Sue Golang meetup 2 is er gekozen om een API te maken.
De map tutorials bevat kleine stukjes voorbeeld code zoals in de presentatie weergegeven.

De map example bevat een werkend voorbeeld van een zelfgeschreven API met minimale externe packages. Dit example gaat ervanuit dat je een database(mariadb) lokaal draait. In de install.md staat een uitwerking zoals wij denken dat een WebAPI zou moeten werken.

### Functionele werking example API:
![alt text](https://raw.githubusercontent.com/tommahs/golang-api/master/example-api/functioneel-api.png)
### Uitleg example API:
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
7. Database probeert uniek id 1 uit events te lezen
8. Database stuurt resultaat naar API
9. API stuurt resultaat naar Operating System
10. Operating System stuurt resultaat naar Internet Browser
11. Internet Browser toont resultaat op het scherm

## Workshop
>- Maak een API die lijkt op het functioneel ontwerp
>- Maak modules en voeg deze toe aan de al geschreven code api -> install.md bevat installatie handleiding
> Eigen idee
