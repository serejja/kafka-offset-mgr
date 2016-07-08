kafka-offset-mgr
================

`kafka-offset-mgr` is a simple command line tool to manage Apache Kafka consumer offsets.

Installation
------------

1. Install Golang [http://golang.org/doc/install](http://golang.org/doc/install)
2. `go get github.com/serejja/kafka-offset-mgr`

Alternatively, you may get the latest release from [releases page](https://github.com/serejja/kafka-offset-mgr/releases).

Usage
-----

```
~> ./kafka-offset-mgr help
NAME:
   kafka-offset-mgr - Apache Kafka consumer offset manager

USAGE:
   kafka-offset-mgr command [command options] [arguments...]
   
VERSION:
   0.1.0
   
COMMANDS:
     view     View consumer offsets
     set      Set consumer offsets
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

Set command
-----------

Set command is used to override an existing consumer offset value. This command accepts 5 flags:

```
~> ./kafka-offset-mgr help set
NAME:
   kafka-offset-mgr set - Set consumer offsets

USAGE:
   kafka-offset-mgr set [command options] [arguments...]

OPTIONS:
   --brokers <host>:<port>, -b <host>:<port>  Comma separated list of Kafka brokers in <host>:<port> form.
   --group value, -g value                    Kafka consumer group to set offset for.
   --topic value, -t value                    Kafka topic to set offset for.
   --partition value, -p value                Kafka topic partition to set offset for. (default: 0)
   --offset value, -o value                   New offset value. (default: 0)
```

Example usage would be:

```
~> ./kafka-offset-mgr set -b localhost:9092 -g groupname -t topicname -p 3 -o 932
```

View command
------------

Not yet implemented.