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


GraphQL_ + PostgreSQL Todo example for Go_.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17`_


Usage
+++++

After git clone this repo, generate and run server:

.. code-block:: bash

  $ cd /path/to/this/repo/
  # Install necessary dependencies
  $ go mod download
  # go generate server code
  $ make schema_generate
  # Run server
  $ make


Code Structure
++++++++++++++

`tools.go <tools.go>`_: Track tool dependencies for a module. See [2]_


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

.. [2] `gqlgen Quick start <https://github.com/99designs/gqlgen#quick-start>`_

       `How can I track tool dependencies for a module? <https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module>`_

.. [3] `github.com/siongui/go-kit-url-shortener-micro-service <https://github.com/siongui/go-kit-url-shortener-micro-service>`_

.. _Go: https://golang.org/
.. _GraphQL: https://graphql.org/
.. _PostgreSQL: https://www.postgresql.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17: https://golang.org/dl/
.. _UNLICENSE: https://unlicense.org/
