## API overview
#### For each API endpoint request should contain authorization cookie. If it does not API server can return public version of response or error if endpoint is for authorized users only

### API Errors:
#### In case of any errors API server returns
```json5
{
  "status": "error",
  "error_detail": "some error description"
}
```
#### General errors (can be returned on each request):
* internal-error

### API default success response:
```json5
{
  "status": "ok",
  "data": null
}
```

## Authorization

### GET /api/session - returns session by token in cookie
#### Response:
```json5
{
  "status": "ok",
  "data": {
    "id": 1
  }
}
```
#### Errors:
* no-auth-cookie
* invalid-auth-cookie

### POST /api/session/register - register new user
#### Body:
```json5
{
  "email": "mail@ya.ru",
  "password": "some_pass"
}
```
#### Response - default
#### Errors:
* email-exists

### POST /api/session/login - Login user in system
#### Body:
```json5
{
  "email": "mail@ya.ru",
  "password": "some_pass"
}
```
#### Response:
```json5
{
  "status": "ok",
  "data": {
    "id": 1
  }
}
```
#### Errors:
* credentials-not-found

### PATCH /api/session/user/password - change user password
#### Body:
```json5
{
  "user_id": 1,
  "password": "new_password"
}
```
#### Response - default
#### Errors:
* user-id-not-found

### POST /api/session/logout - Logout user from system
#### Response - default

## Meetings

### GET /api/meetings - returns all meetings
#### Public response:
```json5
{
  "status": "ok",
  "data": [
    {
      "id": 1,
      "admin_id": 1,
      "title": "meeting title",
      "description": "some meeting description",
      "tags": ["tag1", "tag2"],
      "latitude": 54.0,
      "longitude": 52.0
    },
    {
      "id": 2,
      "admin_id": 2,
      "title": "hello world",
      "tags": ["tag3"],
      "latitude": 54.0,
      "longitude": 52.0
    }
  ]
}
```
#### For registered user:
#### Body:
```json5
{
  "user_id": 3
}
```
#### Response:
```json5
{
  "status": "ok",
  "data": [
    {
      "id": 1,
      "admin_id": 1,
      "current_user_status": "invited", // added this field
      "date_time": "2020-01-24T10:41:21.381Z", // added this field
      "title": "meeting title",
      "description": "some meeting description",
      "tags": ["tag1", "tag2"],
      "latitude": 54.0,
      "longitude": 52.0
    },
    {
      "id": 2,
      "admin_id": 2,
      "current_user_status": "not-invited", // added this field
      "date_time": "2020-01-24T10:41:21.381Z", // added this field
      "title": "hello world",
      "tags": ["tag3"],
      "latitude": 54.0,
      "longitude": 52.0
    }
  ]
}
```
#### Errors:
* user-id-not-found

### POST /api/meeting - creates meeting
#### Body:
```json5
{
  "admin_id": 1,
  "settings": {
    "title": "meeting title",
    "date_time": "2020-01-24T10:41:21.381Z", // ISO format
    "label": "address of meeting",
    // optional
    "max_users": 10,
    "tags": ["tag1", "tag2"],
    "description": "some meeting description",
    "duration": 3, // in hours,
    "min_age": 16,
    "male": "male", // or "female",
    "request_description_required": true
  }
}
```
#### Response - default
#### Errors:
* user-id-not-found

### DELETE /api/meeting - delete meeting
#### Body:
```json5
{
  "meeting_id": 1
}
```
#### Response - default
#### Errors:
* meeting-id-not-found

### PATCH /api/meeting/settings - updates meeting settings
#### Body:
```json5
{
  "meeting_id": 1,
  "settings": {
    "title": "title",
    "date_time": "2020-01-24T10:41:21.381Z", // ISO format
    "label": "address of meeting",
    "max_users": 5,
    "tags": ["tag2"],
    "description": "hello world",
    "duration": 4, // in hours,
    "min_age": 18,
    "male": "female",
    "request_description_required": false
  }
}
```
#### Response - default
#### Errors:
* meeting-id-not-found

### POST /api/meeting/request-participation
#### Body:
```json5
{
  "user_id": 1,
  "meeting_id": 1,
  "request_description": "some description" // optional
}
```
#### Response:
```json5
{
  "too_low_rating_tags": ["tag1", "tag2"], // list of tags rating in whose too low for requested meeting
  "inappropriate_info_fields": [
    {"error_code": "max-users-count-reached", "description": "actual: 10"},
    {"error_code": "age-less-than-min", "description": "actual: 16, wanted: 18"},
    {"error_code": "wrong-gender", "description": "actual: female, wanted: male"}
  ],
  "has_near_meeting": true
}
```
#### Errors:
* user-id-not-found
* meeting-id-not-found

### POST /api/meeting/invite-user
#### Body:
```json5
{
  "user_id": 1,
  "meeting_id": 1
}
```
#### Response - default
#### Errors:
* user-id-not-found
* meeting-id-not-found

### DELETE /api/meeting/user - kick user out of meeting
#### Body:
```json5
{
  "user_id": 1,
  "meeting_id": 1
}
```
#### Response - default
#### Errors:
* user-id-not-found
* meeting-id-not-found


## Users

### GET /api/user/settings/:id - returns user's settings
#### Path parameters
* `:id` - user id
#### Response:
```json5
{
  "status": "ok",
  "data": {
    "name": "User Name",
    "nickname": "user_nickname",
    "gender": "male", // or "female",
    "age": 16,
    "avatar_url": "https://domain.com/avatar.png",
    "rating": {
      "tag1": 80,
      "tag2": 70
    }
  }
}
```
#### Errors:
* user-id-not-found

### PATCH /api/user/settings - change settings
#### Body:
```json5
{
  "user_id": 1,
  "settings": {
    "name": "User Name",
    "nickname": "user_nickname",
    "gender": "male", // or "female",
    "age": 16,
    "avatar_url": "https://domain.com/avatar.png",
  }
}
```
#### Response - default
#### Errors:
* user-id-not-found

## Chatting
### POST /api/chat/meeting - creates chat for meeting
#### Body:
```json5
{
  "meeting_id": 1
}
```
#### Response - default

### POST /api/chat/meeting/request - creates chat with meeting admin
#### Body:
```json5
{
  "meeting_id": 1
}
```
#### Response - default

### PATCH /api/chat/meeting - closes chat
#### Body:
```json5
{
  "chat_id": 1
}
```
#### Response - default
