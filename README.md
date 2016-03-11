# Gonotes

A simple command line note/message/thoughts taking app.

Wrote this mainly to store memes/one line copypasta/ imgur links and shit lol

Drop the code into ```$GOPATH/src/gonotes```

```
cd $GOPATH/src/gonotes
go install
```
Drop the db into ```$GOPATH/bin``` or wherever you move your binary.

Create stuff:
```
./gonotes hello this is my first note.
```
Access the note:
```
./gonotes hello
```
Add notes to lists:
```
./gonotes @pointlessapps hello this is my first note.
```
Access notes in a list:
```
./gonotes -list @pointlessapps
```
Display all existing lists:
```
./gonotes -list
```

