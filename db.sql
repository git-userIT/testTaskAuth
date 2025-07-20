CREATE TABLE users(
	id_user SERIAL PRIMARY KEY,
	username VARCHAR(150) NOT NULL,
	pass VARCHAR(100) NOT NULL,
	email VARCHAR(150) NOT NULL
);

CREATE UNIQUE INDEX idx_username_uniq ON users (username);

INSERT INTO users (username, pass, email) VALUES ('Стартовый пользователь', 'Без пароля', 'Без почты');

CREATE TABLE logs (
 	id_log SERIAL PRIMARY KEY,
 	action_name VARCHAR(50) NOT NULL,
 	table_name_ VARCHAR(150) NOT NULL,
 	username_modify VARCHAR(150) NOT NULL,
 	username_add VARCHAR(150) NOT NULL,
 	time_add timestamp with time zone NOT NULL
);

CREATE OR REPLACE FUNCTION audit_add_new_user_fun() RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO logs (action_name, table_name_, username_modify, username_add, time_add)
  VALUES ('INSERT', TG_TABLE_NAME, current_user, NEW.username, current_timestamp);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER audit_add_new_user
AFTER INSERT ON users
FOR EACH ROW EXECUTE FUNCTION audit_add_new_user_fun();

CREATE OR REPLACE FUNCTION user_exists(userN VARCHAR, p VARCHAR) 
RETURNS BOOLEAN AS $$
BEGIN
    RETURN EXISTS(SELECT 1 FROM users WHERE username = userN AND pass = p);
END;
$$ LANGUAGE plpgsql;