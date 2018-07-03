# lyceum

Lyceum is an open source eBook library server written in Go.

## Overview

The idea for Lyceum came when an arbitrary file size limit was hit while
attempting to upload a newly purchased eBook to a popular cloud based eBook
library service. The desire was born to create a hosted version that can be
deployed on a private network, with the content under full control of the owner.

## Development

Check out the source code and navigate to the project root directory. Ensure the
project dependencies are present.

```
dep ensure
```

Build the application binaries

```
./hack/build
```

Run the application binary that you are interested in.

```
lyceum-api
lyceum-scan
```

#### API

Run the `api` application from the command line.

```
lyceum-api
```

From another window, you can issue curl requests against the api server that is
running on localhost.

```
curl -v localhost:4778/items
```

#### Scan

Run the `scan` application from the command line.

```
lyceum-scan
```

## FAQ

__Where did you get the name for this project?__

As any software engineer out there will probably tell you, the first and
sometimes almost hardest problem to solve is what will we name the project? So,
according to Wikipedia, the [Lyceum](https://en.wikipedia.org/wiki/Lyceum_(Classical))
was a temple dedicated to Apollo Lyceus. It was best known for the Peripatetic
school of philosophy founded there by Aristotle in 334 / 335 BCE. Sounds pretty
good for a library of eBooks.

## License

Lyceum is under the Apache 2.0 license. See the [LICENSE][license_file] file for details.

[license_file]:./LICENSE
