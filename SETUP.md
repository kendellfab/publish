Setup
=====

1. Clone repo
2. Build publish binary
3. Setup mysql database with user and password
	- CREATE DATABASE (name);
	- GRANT ALL PRIVILEGES ON (name).* To '(user)'@'hostname' IDENTIFIED BY '(password)';
4. Copy or alter config file
	1. Admin theme directory (Clone directory)/static/admin
	2. Theme directory (Clone directory)/static/(name)
	3. Upload directory
	4. Mysql connection information
	5. Email
		1. Setup in config file
		2. Or setup via environment variables
5. Go to /setup & create an account for yourself