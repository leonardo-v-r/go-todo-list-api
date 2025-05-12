CREATE TABLE task_tb (
    task_id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_title VARCHAR(255) NOT NULL,
    task_description VARCHAR(255),
    task_status TINYINT NOT NULL,
    created_at DATETIME NOT NULL
);