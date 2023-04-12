CREATE DATABASE `todoApp` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE IF NOT EXISTS todoApp.activities (
    activity_id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    email VARCHAR(60) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    CONSTRAINT activities_PK PRIMARY KEY (activity_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE IF NOT EXISTS todoApp.todos (
    todo_id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    activity_group_id INT(11) UNSIGNED NOT NULL,
    title VARCHAR(255) NOT NULL,
    priority VARCHAR(25) NOT NULL,
    is_active BOOL NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    CONSTRAINT todos_PK PRIMARY KEY (todo_id),
    CONSTRAINT todos_activity_group_id_FK FOREIGN KEY (`activity_group_id`) REFERENCES todoApp.activities (`activity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;