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
(1, 2),
(1, 5),
(1, 8),
(1, 10),
(2, 1),
(2, 7),
(5, 8),
(6, 10);