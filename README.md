# nikkiame

nikkiame(v2)

a blog server written by golang.

## install & run

clone this repo
``` shellsession
git clone https://github.com/eniehack/nikkiamev2.git
```

create configratiuon file, and edit it.

(database.dsn syntax is gorm's. [see also](https://gorm.io/docs/connecting_to_the_database.html))

``` shellsession
cp config.example.toml config.toml
vim config.toml
```

migrate database.

``` shellsession
cd cmd/migrator/
go build
./migrator -c ../../config.toml
```

build nikkiame.

``` shellsession
cd ../../
go build -o nikkiame
```

run server(but http server's configuration is required).

``` shellsession
./nikkiame
```

## License

copyright (c) 2022 eniehack

This software is licensed under [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0) OR [Affero GNU Public License 3.0](https://www.gnu.org/licenses/agpl-3.0.en.html).

## develop

Fork it (https://gitlab.com/eniehack/branca-cr/fork)

clone this repo

``` shellsession
git clone https://github.com/eniehack/nikkiamev2.git
```

Create your feature branch (git checkout -b my-new-feature)

create go.work

``` shellsession
go work init
```

Commit your changes (git commit -am 'Add some feature')

Push to the branch (git push origin my-new-feature)

Create a new Pull Request

## Contributors

* eniehack - creator and maintainer

