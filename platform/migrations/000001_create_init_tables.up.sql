CREATE TABLE teachers (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    email VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (email)
) ENGINE=INNODB;

CREATE TABLE students (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    email VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    is_suspended BOOLEAN DEFAULT FALSE,
    UNIQUE (email)
) ENGINE=INNODB;

CREATE TABLE teachings (
    student_id INT NOT NULL, 
    teacher_id INT NOT NULL, 
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (teacher_id) REFERENCES teachers(id)
) ENGINE=INNODB;

CREATE TABLE notifications (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    teacher_id INT NOT NULL,
    message VARCHAR(255) NOT NULL, 
    FOREIGN KEY (teacher_id) REFERENCES teachers(id)
) ENGINE=INNODB;

CREATE TABLE notifications_recipients (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    notification_id INT NOT NULL,
    recipient_id INT NOT NULL, 
    FOREIGN KEY (notification_id) REFERENCES notifications(id),
    FOREIGN KEY (recipient_id) REFERENCES students(id)
) ENGINE=INNODB;