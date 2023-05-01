# 🌊 Hydrometer

Realtime data from [Duero basin](https://es.wikipedia.org/wiki/Cuenca_hidrogr%C3%A1fica_del_Duero) water reservoirs.

> Data is scraped from [SAIH](https://www.saihduero.es/).

![Duero river basin](https://user-images.githubusercontent.com/57636993/235498359-499e6060-281e-43a7-93ea-42084441c428.png)

## How to use

To clone and run this application, you'll need [Git](https://git-scm.com) and [Go](https://go.dev) installed on your computer.
From your command line:

```bash
$ git clone https://github/lewinkoon/hydrometer
$ cd hydrometer
```

Execute the application

```bash
$ go run cmd/*
```

Now the server is up and running on `http://localhost:4000`

## Credits

This software uses the following open source packages:

- [Colly](https://github.com/gocolly/colly) : elegant scraper and crawler framework 
