# smart-notes-api
A secure REST API where users can register, log in, and manage personal notes. Notes can include text or markdown, and users can upload attachments (images, PDFs) that get stored in AWS S3


- **Features:**
    - User authentication with **JWT**.
    - CRUD operations for notes.
    - File uploads (store links in DB, files in S3).
    - Basic search (by title or content).
    - Deployment with **Docker + AWS ECS/Fargate**.
- **Tech stack:**
    - Go (**Gin**)
    - **PostgreSQL + pgx**
    - AWS S3 (storage)
    - Docker