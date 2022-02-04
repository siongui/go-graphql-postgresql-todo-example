==================================
Go GraphQL PostgreSQL Todo Example
==================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/go-graphql-postgresql-todo-example?status.svg
   :target: https://godoc.org/github.com/siongui/go-graphql-postgresql-todo-example

.. image:: https://github.com/siongui/go-graphql-postgresql-todo-example/workflows/ci/badge.svg
    :target: https://github.com/siongui/go-graphql-postgresql-todo-example/blob/master/.github/workflows/ci.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/go-graphql-postgresql-todo-example
   :target: https://goreportcard.com/report/github.com/siongui/go-graphql-postgresql-todo-example

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/go-graphql-postgresql-todo-example/blob/master/UNLICENSE


`Go kit`_ + GraphQL_ + PostgreSQL_ Todo example for Go_.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17`_

.. contents:: Table of Contents


Usage
+++++

After git clone this repo, generate and run server:

.. code-block:: bash

  $ cd /path/to/this/repo/
  # Install necessary dependencies
  $ go mod download
  # go generate server code
  $ make schema_generate
  # database migrations
  $ make database_migrations_local_development
  # Run server
  $ make


PostgreSQL
++++++++++

Install Docker_ and `Docker Compose`_

.. code-block:: bash

  $ cd /path/to/this/repo/
  $ docker-compose up

psql_ to connect to the database:

.. code-block:: bash

  # Default password: changeme. Can be changed in docker-compose.yml
  $ psql -h localhost -p 5432 -U postgres -d todo_db
  Password for user postgres:
  psql (12.9 (Ubuntu 12.9-0ubuntu0.20.04.1), server 14.1)
  WARNING: psql major version 12, server major version 14.
           Some psql features might not work.
  Type "help" for help.

  todo_db=#


GraphQL Example
+++++++++++++++

.. code-block:: graphql

  {
    getTodo(id: 1) {
      content_code
      content_name
    }
  }

.. code-block:: graphql

  mutation {
    createTodo(input: {content_code: "TD001"}) {
      id
      content_code
      created_date
      updated_date
    }
  }


GraphQL Schema Linter
+++++++++++++++++++++

Use graphql-schema-linter_ for schema linting. See
`.graphql-schema-linterrc <.graphql-schema-linterrc>`_ for linting config.

To run the linter:

.. code-block:: bash

  $ make graphql_schema_lint


Database Migrations
+++++++++++++++++++

golang-migrate_ is used to apply database migrations.

To create migrations, install `golang-migrate CLI`_ first.

.. code-block:: bash

  $ cd /path/to/this/repo/
  $ migrate create -ext sql -dir migrations/ create_todo_table
  migrations/20220202204515_create_todo_table.up.sql
  migrations/20220202204515_create_todo_table.down.sql

Edit the ``up.sql`` and ``down.sql`` accordingly. After finish, set
**POSTGRESQL_URL** to tell migrate CLI where the database is:

.. code-block:: bash

  $ export POSTGRESQL_URL='postgres://postgres:changeme@localhost:5432/todo_db?sslmode=disable'

Now we apply the migrations to the database:

.. code-block:: bash

  $ migrate -database ${POSTGRESQL_URL} -path migrations/ up

Check if the migrations is correctly applied:

.. code-block:: bash

  # Default password: changeme. Can be changed in docker-compose.yml
  $ psql -h localhost -p 5432 -U postgres -d todo_db
  Password for user postgres:
  psql (12.9 (Ubuntu 12.9-0ubuntu0.20.04.1), server 14.1)
  WARNING: psql major version 12, server major version 14.
           Some psql features might not work.
  Type "help" for help.

  todo_db=# \dt+
                              List of relations
   Schema |       Name        | Type  |  Owner   |    Size    | Description
  --------+-------------------+-------+----------+------------+-------------
   public | schema_migrations | table | postgres | 8192 bytes |
   public | todos             | table | postgres | 8192 bytes |
  (2 rows)

  todo_db=# TABLE todos;
   id | content_code | created_at | updated_at | deleted_at | content_name | description | start_date | end_date | status | created_by | updated_by
  ----+--------------+------------+------------+------------+--------------+-------------+------------+----------+--------+------------+------------
  (0 rows)

Migrations with gorm.Model
--------------------------

gorm_ is ORM library for Go. The migration SQL for gorm.Model_:

.. code-block:: go

  // gorm.Model embedded in MyType
  type MyType struct {
  	gorm.Model
  }

The table name is *my_types*

.. code-block:: sql

  CREATE TABLE "my_types"
  (
      "id"         bigserial,
      "created_at" timestamptz NOT NULL,
      "updated_at" timestamptz NOT NULL,
      "deleted_at" timestamptz,
      PRIMARY KEY ("id")
  );

  CREATE INDEX "idx_my_type_deleted_at" ON "my_types" ("deleted_at")


Code Structure
++++++++++++++

- `config/ <config/>`_: application configuration
- `graph/ <graph/>`_: GraphQL schema
- `todo/ <todo/>`_: Go micro service - *todo*
- `tools/tools.go <tools/tools.go>`_: Track tool dependencies for a module.
  See [2]_


Issues
++++++

- generating core failed: comment the ``autobind`` in https://gqlgen.com/config.
  See `generating core failed: unable to load example/graph/model in v0.16 <https://github.com/99designs/gqlgen/issues/1860>`_


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `github.com/99designs/gqlgen <https://github.com/99designs/gqlgen>`_
.. [2] | `gqlgen Quick start <https://github.com/99designs/gqlgen#quick-start>`_
       | `How can I track tool dependencies for a module? <https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module>`_
.. [3] `github.com/siongui/go-kit-url-shortener-micro-service <https://github.com/siongui/go-kit-url-shortener-micro-service>`_
.. [4] `Building a GraphQL Server with Go Backend Tutorial | Getting Started <https://www.howtographql.com/graphql-go/1-getting-started/>`_
.. [5] `How To Remove Docker Images, Containers, and Volumes | DigitalOcean <https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes>`_
.. [6] | `go kit graphql <https://www.google.com/search?q=go+kit+graphql>`_
       | `GraphQL support · Issue #636 · go-kit/kit · GitHub <https://github.com/go-kit/kit/issues/636>`_
       | `Add initial GraphQL support by sagikazarmark · Pull Request #81 · sagikazarmark/modern-go-application · GitHub <https://github.com/sagikazarmark/modern-go-application/pull/81>`_
.. [7] `jinzhu/configor: Golang Configuration tool that support YAML, JSON, TOML, Shell Environment <https://github.com/jinzhu/configor>`_

.. _Go: https://golang.org/
.. _Go kit: https://gokit.io/
.. _GraphQL: https://graphql.org/
.. _PostgreSQL: https://www.postgresql.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17: https://golang.org/dl/
.. _Docker: https://docs.docker.com/engine/install/
.. _Docker Compose: https://docs.docker.com/compose/install/
.. _psql: https://www.postgresguide.com/utilities/psql/
.. _graphql-schema-linter: https://github.com/cjoudrey/graphql-schema-linter
.. _golang-migrate: https://github.com/golang-migrate/migrate
.. _golang-migrate CLI: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
.. _gorm: https://gorm.io/
.. _gorm.Model: https://gorm.io/docs/models.html#gorm-Model
.. _UNLICENSE: https://unlicense.org/
