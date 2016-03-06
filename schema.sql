-- ServiceD schema

DROP DATABASE IF EXISTS serviced;
CREATE DATABASE serviced;
USE serviced;

CREATE TABLE environments (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL
) ENGINE=INNODB;

CREATE TABLE applications (
  id              INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name            VARCHAR(255) NOT NULL UNIQUE,
  type            ENUM('DOCKER', 'CMD', 'VIRTUAL') NOT NULL,

  cmdline         VARCHAR(255) NULL,

  docker_image    VARCHAR(255) NOT NULL,
  docker_args     VARCHAR(255) NULL,

  instance_count  INTEGER NOT NULL DEFAULT 1,
  network_type    ENUM('HOST', 'BRIDGED') NOT NULL
) ENGINE=INNODB;

CREATE TABLE application_dependencies (
  parent_id INTEGER NOT NULL,
  dependency_id INTEGER NOT NULL,
  is_hard BOOLEAN NOT NULL,

  PRIMARY KEY (parent_id, dependency_id),

  FOREIGN KEY (parent_id) REFERENCES applications(id) ON DELETE CASCADE,
  FOREIGN KEY (dependency_id) REFERENCES applications(id) ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE application_releases (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  release_tag VARCHAR(255) NOT NULL,
  application_id INTEGER NOT NULL,

  FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE deployments (
  id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,

  environment_id INTEGER NOT NULL,
  application_id INTEGER NOT NULL,
  release_id INTEGER NOT NULL,

  instance_count INTEGER NOT NULL DEFAULT 1,
  marathon_path VARCHAR(255) NOT NULL UNIQUE,

  FOREIGN KEY (environment_id) REFERENCES environments(id) ON DELETE CASCADE,
  FOREIGN KEY (application_id) REFERENCES applications(id) ON DELETE CASCADE,
  FOREIGN KEY (release_id) REFERENCES application_releases(id) ON DELETE CASCADE
) ENGINE=INNODB;

