# Heb12 Go

This program is the basic code which all Heb12 software can include as a library. It is meant to easily centralize all of the code that is shared between the programs. That way each program can use this as a library, but have the platform specific code stored separately.

These new tools are written in Go. This is done because Go is fast, simple, enforces good documentation, and has good support throughout different platforms. It can be used as the foundation for the destkop app, the mobile apps, the website back end, and (with WebAssembly) even the website front end if so desired.

Some code is stored in other repositories. This is because those modules may specifically be useful for other programs, and it makes sense to maintain them separately. [bref](https://code.heb12.com/Heb12/bref) is an example of this.

**Note:** This software is in its early stages and is subject to have breaking changes. Do not depend on it for anything important yet.

## TODO

- [x] Basic Bible version parser (bver)
- [x] Basic OSIS document manager (osistool)
- [ ] CLI tool (a basic CLI wrapper around everything)
- [x] Separate osistool between just OSIS parsing and actually managing the OSIS works
- [ ] Search tool (client side generation and rendering of search texts)

## Packages

This heb12 module includes several packages which have different uses. 

### bible

bible uses the modules bref, heb12/manage, heb12/config, and heb12/osis to get Bible verse.

[More documentation](https://pkg.go.dev/code.heb12.com/heb12/heb12/bible?tab=doc)

### bver

A Bible version parser. Useful for parsing human input when it comes to translations, but it does not include all possible Bible versions, and it is not meant to do so. This should not be used in most situations.

[More documentation](https://pkg.go.dev/code.heb12.com/heb12/heb12/bver?tab=doc)

### cmd

A basic CLI for Heb12. It (will) include features to use the rest of the heb12 module, and for reading the Bible from the command line.

### config

Config is made for Heb12-specific configuration for using the other packages.

[More documentation](https://pkg.go.dev/code.heb12.com/heb12/heb12/config?tab=doc)

### manage

A manager for OSIS works in a directory. It provides information about them and manages the directory structure.

[More documetation](https://pkg.go.dev/code.heb12.com/heb12/heb12/manage?tab=doc)

### osis

An OSIS parser. It provides functions to parse OSIS Bible files, and give information about them and the Bible text itself.

[More documetation](https://pkg.go.dev/code.heb12.com/heb12/heb12/osis?tab=doc)

## License

Copyright (C) 2020 Ted Jameson

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

See COPYING.LESSER for more license information.
