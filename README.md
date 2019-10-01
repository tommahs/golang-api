# golang-api
> API geschreven in de programmeertaal Go met gebruik van de standaard libraries aangeleverd door Go.

Auteur: Tom van Hamersveld


## Introductie
### Functionele werking API:
>![alt text](https://raw.githubusercontent.com/tommahs/golang-api/master/functioneel-api.png?token=AFMV6OLIHODYLHMQCXXUIUC5TR7AS)
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
6. Vraagt aan de database
7. Probeert uniek 1 uit events te lezen
8. Stuurt resultaat naar API
9. Stuurt resultaat naar Operating System
10. Stuurt resultaat naar Internet Browser
11. Toont resultaat op het scherm




## Setup

## Assignment
