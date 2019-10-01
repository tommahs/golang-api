# golang-api
Auteur: Tom van Hamersveld
## Introductie

Een Web API is Application programming interface bereikbaar over het netwerk. De Sue student edition heeft als thema "data"  en deze API zorgt ervoor dat de data opgehaald en weergegeven wordt. Je zou een API kunnen zien als een doorgeef luik.

In de install.md staat een uitwerking zoals wij denken dat een WebAPI zou moeten werken. Aan de lezer de uitdaging om dit na te bouwen en/of uit te breiden.

<b>Mocht je problemen hebben, probeer te lokaliseren waar het probleem zit door de volgende code toe te voegen aan het gedeelte waar je mee bezig bent :</b>
> fmt.Printf("%T", "ZET HIER DE VARIABELE NEER WAAR JE DENKT DAT HET FOUT GAAT")

> fmt.Printf("%v", "ZET HIR DE VARIABELE NEER WAAR JE DENKT DAT HET FOUT GAAT")

### Begrippen:
API:
>https://computerworld.nl/development/74796-wat-is-een-api

Database:
Verzameling tabellen beschikbaar gemaakt in software. Een simpele tabel in een database is vergelijkbaar met een excel-sheet
>http://www.webdesign-gids.nl/wat_is_een_database



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
7. Database probeert uniek id 1 uit events te lezen
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

Om de inhoud van een bestand te zien:
> "cat bestandsnaam"
## Challenge
Wil je graag Go leren programmeren? Denk je dat jij de code die wij geschreven hebben kan verbeteren?
Wil je graag gebruik maken van deze code om iets voor jezelf te maken? Dat kan!
Wij staan open voor ideeen!

Mogelijkheden:
>- Maak een API die lijkt op het functioneel ontwerp

>- Maak modules en voeg deze toe aan de al geschreven code api -> install.md bevat installatie handleiding

>- (hobby / huiswerk) Zoek op youtube naar "Sentdex", deze man heeft playlists voor programmeren in zowel python als go.
