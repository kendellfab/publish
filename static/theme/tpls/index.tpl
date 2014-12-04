{{ define "Content" }}
	<h1>Frontend Content</h1>
	{{ if .error }}
		<p>{{ .error }}</p>
	{{ else }}
		{{ range .posts }}
			<hr />
			{{ rend_md .Content }}
		{{ end }}
	{{ end }}
{{ end }}