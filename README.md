# mempwgen

a password generator that makes secure but memorable passwords.

- 20-27 characters
- 1 symbol (`#`, `@`, `-`, `_`)
- 3 numbers
- 2 words

## install

```
$ go install -v code.nkcmr.net/mempwgen@latest
```

## sample

```
$ mempwgen
inundated686-inclined
overactive681@scotch
corporal565#interpret
repetition175#lapdog
circuits688_warriors
masculine913@equation
breathable975-factory
cheeseburger323_dandruff
neediness260-victoria
coroner119_drugstore
deputized627#restful
disruptions591-forbade
becoming12-anesthetic
fatso664@spontaneity
subjugation1-upholstery
...
```

the dictionary used contained some offensive words when i first found it. i did the best i could to scrub of everything that wouldn't be appropriate in a work context. if you find something offensive, please make an edit and open a PR, i'll be sure to take it out. 🙂

## license

```
Copyright (c) 2020 Nicholas Comer

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
