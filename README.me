# Post Service

Manage posts creations and publish events related to posts.

## API Reference

#### Create post

```http
  POST /api/v1/posts
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Content-Type` | `string` | **Required**. application/json |

**Request Body:**

| Field | Type     | Description                       |
| :---- | :------- | :-------------------------------- |
| `id`  | `string` | **Required**. Unique identifier for the post |
| `author_id` | `string` | **Required**. ID of the post author |
| `contents` | `array` | **Required**. Array of content objects |
| `contents[].type` | `string` | **Required**. Type of content (e.g., "text") |
| `contents[].text` | `string` | Text content |
| `published_at` | `string` | **Required**. ISO 8601 timestamp |
| `created_at` | `string` | **Required**. ISO 8601 timestamp |
| `updated_at` | `string` | **Required**. ISO 8601 timestamp |

#### Get post by ID

```http
  GET /api/v1/posts/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. ID of post to fetch |

#### Get multiple posts by IDs

```http
  GET /api/v1/posts?ids=${comma_separated_ids}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `ids`     | `string` | **Required**. Comma-separated list of post IDs (URL encoded) |

## How to Run?


In order to run you have to clone:
- `uala-posts-service`
- `uala-followers-service`
- `uala-timeline-service`

In the `uala-posts-service` repository, you'll find the `docker-compose` file used to start all services.

### Steps:


1. After cloning each repository, run the following command inside each one:

   ```bash
   make build
   ```

2. Once all images are built, navigate to the uala-posts-service repository and run:

   ```bash
   docker-compose up -d
   ```

### Service Ports:
- posts-service: 8080
- timeline-service: 8081
- followers-service: 8082


## Usage/Examples

### Create Post

```bash
curl --location 'localhost:8080/api/v1/posts' \
--header 'Content-Type: application/json' \
--data '{
    "id": "c7e7fbeb-1a7b-4966-8a59-bd2178244af8",
    "author_id": "123abc",
    "contents": [
        {
            "type": "text",
            "text": "Lorem test text"
        }
    ],
    "published_at": "2025-05-14T16:21:45.959431-03:00",
    "created_at": "2025-05-14T16:27:28.280206-03:00",
    "updated_at": "2025-05-14T16:27:28.280206-03:00"
}'
```

