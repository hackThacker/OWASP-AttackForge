-- Create the schemas for the vulnerable applications
CREATE DATABASE IF NOT EXISTS dvwa;
CREATE DATABASE IF NOT EXISTS xvwa;
CREATE DATABASE IF NOT EXISTS mutillidae;
CREATE DATABASE IF NOT EXISTS bWAPP;
CREATE DATABASE IF NOT EXISTS vwa;

-- Grant global privileges to the custom user from any host on the network
-- (The MARIADB_USER environment variable automatically creates the user,
-- but we grant it broad permissions over all databases)
GRANT ALL PRIVILEGES ON *.* TO 'hackthacker'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
