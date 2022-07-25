package tmpl

var PageTmpl = `
<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.2/css/bulma.min.css">
</head>

<body>

{{ range .Messages }}
<div class="card">
    <div class="card-header">[{{.ID}}]</div>
    <div class="card-content">{{.Body}}</div>
</div>
{{ end }}


{{if eq .No .LastNo }}
<div class="card is-three-fifths-widescreen">
    <div class="card-header">[NEW]</div>
    <div class="card-content">
        <form action="/{{.Title}}/{{.No}}" method="POST">
            <div>
                <textarea class="textarea" name="body"></textarea>
            </div>
            <div>
                <input class="input" type="submit">
            </div>

        </form>
    </div>
</div>
{{end}}

<nav class="pagination is-rounded" role="navigation" aria-label="pagination">
    {{if ne .PrevNo -1}}
    <a class="pagination-previous" href="/{{.Title}}/{{.PrevNo}}">Previous</a>
    {{end}}
    {{if ne .NextNo -1}}
    <a class="pagination-next" href="/{{.Title}}/{{.NextNo}}">Next</a>
    {{end}}
</nav>

</body>

</html>
`
