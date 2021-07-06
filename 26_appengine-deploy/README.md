# Buying a domain

[https://domains.google/#/](https://domains.google/#/)

# Deploying to Google Cloud

- [install google appengine](https://domains.google/#/)
  - configure environment PATH variables
- google cloud developer console
  - create a project
  - get the project ID
- set your go version in the app.yaml file

```yaml
runtime: go113
handler:
- url: /.*
  scripe: auto
  secure: always
```

- deploy to the project. update --project with your project-id
  
```bash
gcloud app deploy app.yaml --project=<YOUR_PROJECT_ID> -v 1
my example:
gcloud app deploy --project temp-137512
```

- view your project
  - http://YOUR_PROJECT_ID.appspot.com/

example http://temp-145415.appspot.com/

change DNS info at google domains point your domain to your appengine project.

