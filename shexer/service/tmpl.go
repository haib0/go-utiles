package service

var indexTmpl = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>
    {{ range . }}
        <p>----------</p>
        <p></p>
        <a href="/run/{{ . }}">{{ . }}</a>
        <p></p>
        <p>----------</p>
    {{end}}
</body>
</html>
`
