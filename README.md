# Lyceum

Lyceum is an open source eBook management system written in Go.

## Overview

The idea for Lyceum came when an arbitrary file size limit was hit while
attempting to upload a newly purchased eBook to a popular cloud based eBook
library service. The desire was born to create a hosted version that can be
deployed on a private network, with the content under full control of the owner.
``
## Development

Start the web server:

```
revel run myapp
```

Go to http://localhost:9000/ and you'll see:

    "It works"

### Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites

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
