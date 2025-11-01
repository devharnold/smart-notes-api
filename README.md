# smart-notes-api
A secure REST API where users can register, log in, and manage personal notes. Notes can include text or markdown, and users can upload attachments (images, PDFs) that get stored in AWS S3


- **Features:**
    - User authentication with **JWT**.
    - CRUD operations for notes.
    - File uploads (store links in DB, files in S3).
    - Basic search (by title or content).
    - Deployment with **Docker**.
- **Tech stack:**
    - Go (**Gin**)
    - **PostgreSQL + pgx**
    - AWS S3 (storage)
    - Docker

- **sync dependencies:**
    - go mod tidy


### Cloning the repo

```
git clone https://github.com/devharnold/smart-notes-api.git
```

### Sync Dependancies

```
go mod tidy
```

Then make sure that you have an environment variables file where you are going to have the Database URL. I did this in mind that I am running on localhost, on the Dockerfile you will see that I have baked the .env file in it.

Then when you build your docker container, build it with:

```
docker build -t "name of your container image"
```

If that works, proceed to run it with:

```
docker run -d -p "your_port:your_port" --env-file .env --name "a name to tag the container" "name of your container image"
```

Here, we use a name to tag the container image to avoid the struggle of trying to find the container with the IDs.