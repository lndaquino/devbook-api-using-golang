## Devbook - API using GO
&nbsp;

### Description
Devbook is a social media devoted to developers. It has basic features related to Users and Posts management.
You can follow/unfollow users and like/dislike their posts
&nbsp;
&nbsp;

### How to run
- go run main.go
&nbsp;
- or generating the binary file (go build) and running it
&nbsp;
&nbsp;


### API Setup
Before running API it´s necessary to run MySQL server and configure it in .env file. Run the scripts on file /sql/sql.sql to create database and tables. Also configure API´s secret key and the port it will be accessible.
&nbsp;
&nbsp;


## Routes
To access authenticated routes you need to send a Bearer token (jwt) in request header field "Authorization". Body fiels are send in JSON format under header field "Content-Type" = application/json
&nbsp;

### User routes
- creating an user: /users
  - method: POST
  - body: {
    "name": string, "nick": string, "email": string, "password": string
  }
- login: /login
  - method: POST
  - body: {
    "email": string, "password": string
  }
  - response: jwt token
- getting all user: /users
  - method: GET
  - needs authentication
- getting an user by ID: /users/{userID}
  - method: GET
  - needs authentication
- updating an user by ID: /users/{userID}
  - method: PUT
  - needs authentication
  - body: {
    "name": string, "nick": string, "email": string
  }
- deleting an user by ID: /users/{userID}
  - method: DELETE
  - needs authentication
- following an user: /users/{userID}/follow
  - method: POST
  - needs authentication
  - obs: the follower user is got from jwt token and the userID params is the user to be followed
- unfollowing an user: /users/{userID}/unfollow
  - method: POST
  - needs authentication
  - obs: the follower user is got from jwt token and the userID params is the user to be unfollowed
- getting user followers: /users/{userID}/followers
  - method: GET
  - needs authentication
- getting who user is following: /users/{userID}/following
  - method: GET
  - needs authentication
- updating password: /users/{userID}/updatePassword
  - method: POST
  - needs authentication
  - body: {
    "password" : string, "newPassword" : string
  }

&nbsp;
&nbsp;
### Post routes
- creating a post: /posts
  - method: POST
  - body: {
    "title": string, "content": string
  }
  - needs authentication
- getting user posts and from who he´s following: /posts
  - method: GET
  - needs authentication
- getting a specific post: /posts/{postID}
  - method: GET
  - needs authentication
- updating a post: /posts/{postID}
  - method: PUT
  - body: {
    "title": string, "content": string
  }
  - needs authentication
- deleting a post: /posts/{postID}
  - method: DELETE
  - needs authentication
- getting posts from a specific user: /users/{userID}/posts
  - method: GET
  - needs authentication
- to like a post: /posts/{postID}/like
  - method: POST
  - needs authentication
- to dislike a post: /posts/{postID}/dislike
  - method: POST
  - needs authentication