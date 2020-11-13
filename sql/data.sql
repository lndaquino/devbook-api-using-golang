INSERT INTO users (name, nick, email, password)
values
("User1", "user1", "user1@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa"),
("User2", "user2", "user2@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa"),
("User3", "user3", "user3@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa"),
("User4", "user4", "user4@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa"),
("User5", "user5", "user5@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa"),
("User6", "user6", "user6@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa"),
("User7", "user7", "user7@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa"),
("User8", "user8", "user8@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa"),
("User9", "user9", "user9@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa"),
("User10", "user10", "user10@gmail.com", "$2a$10$G3ZczxeMBx2fKHB/1jWeEeESmp4AaRBSnvchP8bx40IJkKBCy8yMa");

INSERT INTO followers (userID, followerID)
values
(15, 19),
(16, 19),
(17, 19),
(18, 19);

insert into posts (title, content, userID)
values
("User 15 Post", "It´s an user´s 15 post! Uhu!", 15),
("User 16 Post", "It´s an user´s 16 post! Uhu!", 16),
("User 17 Post", "It´s an user´s 17 post! Uhu!", 17),
("User 18 Post", "It´s an user´s 18 post! Uhu!", 18),
("User 19 Post", "It´s an user´s 19 post! Uhu!", 19);