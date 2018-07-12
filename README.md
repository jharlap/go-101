# Intro to Go Workshop

These materials are intended for an Intro to Go workshop that is suitable for
programmers who are comfortable with a modern programming language and wish to
learn Go. The workshop should take about 4 hours to run in a small classroom
setting (~12 students), giving everyone sufficient time to complete each
exercise.

To view the slides:

```bash
python -m SimpleHTTPServer 8081 # python2
python -m http.server 8081 # python3
```

And browse to http://localhost:8081/

## Editing/Building

Edit slides.embedded.md and then run `make` to update slides.md. Make sure
`embedmd` is installed first, `go get -u github.com/campoy/embedmd` or it won't
work.
