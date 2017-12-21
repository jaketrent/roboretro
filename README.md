## Deploy

To heroku:

One time:

```sh
heroku create myroboretro
heroku addons:create heroku-postgresql:hobby-dev
heroku buildpacks:set heroku/nodejs
heroku buildpacks:add heroku/go
```

Each commit:

```sh
git push heroku master
```
