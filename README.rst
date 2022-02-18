==================================
Go GraphQL PostgreSQL Todo Example
==================================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/go-kit-gqlgen-postgres-todo-example?status.svg
   :target: https://godoc.org/github.com/siongui/go-kit-gqlgen-postgres-todo-example

.. image:: https://github.com/siongui/go-kit-gqlgen-postgres-todo-example/workflows/ci/badge.svg
    :target: https://github.com/siongui/go-kit-gqlgen-postgres-todo-example/blob/master/.github/workflows/ci.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/go-kit-gqlgen-postgres-todo-example
   :target: https://goreportcard.com/report/github.com/siongui/go-kit-gqlgen-postgres-todo-example

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/go-kit-gqlgen-postgres-todo-example/blob/master/UNLICENSE


`Go kit`_ + GraphQL_ (gqlgen_) + PostgreSQL_ + gorm_ Todo example for Go_.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17`_

.. contents:: Table of Contents


Usage
+++++

After git clone this repo, generate and run server:

.. code-block:: bash

  $ cd /path/to/this/repo/
  # go generate GraphQL code
  $ make schema_generate

See next section to run dockerized PostgreSQL. After PostgreSQL is running:

.. code-block:: bash

  # database migrations
  $ make database_migrations
  # Run server
  $ make

After the server is running, click on the link of log message. Now you can send
GraphQL requests via web interface.


Run Dockerized PostgreSQL
+++++++++++++++++++++++++

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

*getTodo* example #1:

.. code-block:: graphql

  {
    getTodo(id: "1") {
      id
      content_code
      created_date
      updated_date
      content_name
      description
      start_date
      end_date
      status
      created_by
      updated_by
    }
  }

*createTodo* example #1:

.. code-block:: graphql

  mutation {
    createTodo(
      input: {
        content_code: "TD001",
        content_name: "my name",
        description: "test",
        start_date: "2022-01-23T14:20:50.52+08:00",
        end_date: "2022-01-30T14:20:50.52+08:00",
        status: Active}
    ) {
      id
      content_code
      created_date
      updated_date
    }
  }

*createTodo* example #2:

.. code-block:: graphql

  mutation {
    createTodo(
      input: {
        content_code: "TD003",
        content_name: "my content name",
        description: "my content description",
        start_date: "2022-01-23T14:20:50.52+08:00",
        end_date: "2022-01-30T14:20:50.52+08:00",
        status: Active}
    ) {
      id
      content_code
      created_date
      updated_date
      content_name
      description
      start_date
      end_date
      status
      created_by
      updated_by
    }
  }

*TodoPages* example:

.. code-block:: graphql

  {
    TodoPages(paginationInput: {count: 5, page: 1}) {
      pagination_info {
        total_count
        total_pages
        current_page
      }
      todos {
        id
        content_code
        created_date
        updated_date
        content_name
        description
        start_date
        end_date
        status
        created_by
        updated_by
      }
    }
  }

*updateTodo* example #1:

.. code-block:: graphql

  mutation {
    updateTodo(
      id: "1000"
      input: {
        content_code: "TD0031",
        content_name: "my updated content name",
        description: "my updated content description",
        start_date: "2022-02-01T14:20:50.52+08:00",
        end_date: "2022-02-02T14:20:50.52+08:00",
        status: Inactive}
    ) {
      id
      content_code
      created_date
      updated_date
      content_name
      description
      start_date
      end_date
      status
      created_by
      updated_by
    }
  }

*updateTodo* example #2:

.. code-block:: graphql

  mutation {
    updateTodo(
      id: "5"
      input: {
        content_code: "TD007",
        content_name: "my updated content name2",
        description: "my updated content description7",
        status: Inactive}
    ) {
      id
      content_code
      created_date
      updated_date
      content_name
      description
      start_date
      end_date
      status
      created_by
      updated_by
    }
  }

*TodoSearch* example #1:

.. code-block:: graphql

  query {
    TodoSearch(
      paginationInput: {count: 5, page: 1}
      input:{
        content_code: "00"
      }
    ) {
      pagination_info {
        total_count
        total_pages
        current_page
      }
      todos {
        id
        content_code
        created_date
        updated_date
        content_name
        description
        start_date
        end_date
        status
        created_by
        updated_by
      }
    }
  }

*TodoSearch* example #2:

.. code-block:: graphql

  query {
    TodoSearch(
      paginationInput: {count: 5, page: 1}
      input:{
        start_date: "2022-02-06T07:11:18+08:00"
        status: Inactive
      }
    ) {
      pagination_info {
        total_count
        total_pages
        current_page
      }
      todos {
        id
        content_code
        created_date
        updated_date
        content_name
        description
        start_date
        end_date
        status
        created_by
        updated_by
      }
    }
  }


Linter
++++++

Two linters are used. graphql-schema-linter_ and golangci-lint_.

GraphQL Schema Linter
---------------------

Use graphql-schema-linter_ for GraphQL schema linting. See
`.graphql-schema-linterrc <.graphql-schema-linterrc>`_ for linting config.

To run the linter:

.. code-block:: bash

  $ make graphql_schema_lint

golangci-lint
-------------

Use golangci-lint_ for Go code linting. See
`.golangci.yml <.golangci.yml>`_ for linter config.

To install golangci-lint:

.. code-block:: bash

  $ make install_golangci_lint

To run golangci-lint:

.. code-block:: bash

  $ make golangci_lint


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


Authentication and Permission
+++++++++++++++++++++++++++++

Steps to implement:

1. Get RSA Public key from http endpoint of identity provider.
2. Use gin or chi to to get the token from header, and passed the token to
   gqlgen resolver.
3. Verify and extract user info and permissons from the token.

- | `go - How to propagate context values from Gin middleware to gqlgen resolvers? - Stack Overflow <https://stackoverflow.com/questions/67267065/how-to-propagate-context-values-from-gin-middleware-to-gqlgen-resolvers>`_
  | `Providing authentication details through context — gqlgen <https://gqlgen.com/recipes/authentication/>`_
  | `Using Gin to setup HTTP handlers — gqlgen <https://gqlgen.com/recipes/gin/>`_

- `GitHub - ghiden/go-kit-stringsvc2-with-jwt: Go kit: stringsvc2 with JWT <https://github.com/ghiden/go-kit-stringsvc2-with-jwt>`_
- `Go-kit微服务| JWT身份认证 <https://liu-yt.github.io/2019/06/23/Go-kit%E5%BE%AE%E6%9C%8D%E5%8A%A1-JWT%E8%BA%AB%E4%BB%BD%E8%AE%A4%E8%AF%81/>`_
- `go-kit 微服务 身份认证 （JWT） | hwholiday <https://www.hwholiday.com/2020/go_kit_v3/>`_
- `go-kit authentication - Google search <https://www.google.com/search?q=go-kit+authentication>`_
- `Get HTTP headers in directives/mutations · Issue #262 · 99designs/gqlgen · GitHub <https://github.com/99designs/gqlgen/issues/262>`_
- `laisky-blog: [Golang] 使用 gqlgen 编写 GraphQL 后端 <https://blog.laisky.com/p/gqlgen/>`_
- `Using schema directives to implement permission checks — gqlgen <https://gqlgen.com/reference/directives/>`_
- `gqlgen data validation · GitHub <https://gist.github.com/david-yappeter/3b9c1d68588dc3e1fb4a4a6efc086d07>`_
- `Mapping GraphQL scalar types to Go types — gqlgen <https://gqlgen.com/reference/scalars/>`_
- `使用 GraphQL Gateway 串接多個 Data Schema - 小惡魔 - AppleBOY <https://blog.wu-boy.com/2021/02/graphql-gateway-in-golang/>`_
- `GraphQL with Golang: Building a GraphQL Server with GO and MySQL | by Vishal Jain | Towards Dev <https://towardsdev.com/graphql-with-golang-building-a-graphql-server-with-go-and-mysql-b931e8c3e3d3>`_
- `golang jwt parse rsa from url - Google search <https://www.google.com/search?q=golang+jwt+parse+rsa+from+url>`_


Code Structure
++++++++++++++

app config
----------

- `config/ <config/>`_: application configuration

GraphQL
-------

- `gqlgen.yml <gqlgen.yml>`_: gqlgen config file.
- `graph/ <graph/>`_: GraphQL schema, resolvers, and custom scalar.

business logic
--------------

- `todo/ <todo/>`_: Go micro service - *todo*
- `todo/tododb/ <todo/tododb/>`_: database library for *todo* service.

database migrations
-------------------

- `tools/migrate/ <tools/migrate/>`_: command line tool for database migrations.
- `migrations/ <migrations/>`_: database migrations SQL files.

dependency tracking
-------------------

- `tools/tools.go <tools/tools.go>`_: Track tool dependencies for a module.
  See [2]_


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
.. [8] generating core failed: comment the ``autobind`` in https://gqlgen.com/config.
       See `generating core failed: unable to load example/graph/model in v0.16 <https://github.com/99designs/gqlgen/issues/1860>`_
.. [9] | `golang migrate err no change - Google search <https://www.google.com/search?q=golang+migrate+err+no+change>`_
       | `go - golang-migrate no change error on initial migration - Stack Overflow <https://stackoverflow.com/questions/67910574/golang-migrate-no-change-error-on-initial-migration>`_
       | `Migrate.Up() errors out if the latest schema is in use · Issue #100 · golang-migrate/migrate · GitHub <https://github.com/golang-migrate/migrate/issues/100>`_
.. [10] | `Using Gin to setup HTTP handlers — gqlgen <https://gqlgen.com/recipes/gin/>`_
        | `How can i use gin with standard handlers? · Issue #57 · gin-gonic/gin · GitHub <https://github.com/gin-gonic/gin/issues/57>`_
        | `Using http.Handler? · Issue #293 · gin-gonic/gin · GitHub <https://github.com/gin-gonic/gin/issues/293>`_


.. _Go: https://golang.org/
.. _Go kit: https://gokit.io/
.. _GraphQL: https://graphql.org/
.. _gqlgen: https://github.com/99designs/gqlgen
.. _PostgreSQL: https://www.postgresql.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17: https://golang.org/dl/
.. _Docker: https://docs.docker.com/engine/install/
.. _Docker Compose: https://docs.docker.com/compose/install/
.. _psql: https://www.postgresguide.com/utilities/psql/
.. _graphql-schema-linter: https://github.com/cjoudrey/graphql-schema-linter
.. _golangci-lint: https://golangci-lint.run/
.. _golang-migrate: https://github.com/golang-migrate/migrate
.. _golang-migrate CLI: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
.. _gorm: https://gorm.io/
.. _gorm.Model: https://gorm.io/docs/models.html#gorm-Model
.. _UNLICENSE: https://unlicense.org/
