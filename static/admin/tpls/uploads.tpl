{{ define "Content" }}
<div class="row">
    <div class="col-lg-12">
        <h1 class="page-header">
            Uploads
        </h1>
    </div>
</div>
<div class="row">
	<div class="col-lg-3">
	</div>
	<div class="col-lg-9">
		<div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">File Upload</h3>
			</div>
			<div class="panel-body">
				<form method="POST" action="/admin/upload" enctype="multipart/form-data">
					<input type="file" name="upload">
					<br />
    				<button type="submit" class="btn btn-primary">Upload</button>
				</form>
			</div>
		</div>
	</div>
</div>
<div class="row">
	<div class="col-lg-12">
		<table class="table table-bordered table-hover table-striped">
			<thead>
				<tr>
					<th>Thumbnail</th>
					<th>File Name</th>
					<th>Size</th>
					<th>Link</th>
					<th>Delete</th>
				</tr>
			</thead>
			<tbody>
				{{ range .files }}
					<tr>
						<td><a href="/uploads/{{ .Name }}"><img src="/uploads/{{ .Name }}" height="75" /></a></td>
						<td>{{ .Name }}</td>
						<td>{{ .Size }}</td>
						<td>/uploads/{{ .Name }}</td>
						<td><a href="/admin/uploads/{{ .Name }}/delete">Delete</a></td>
					</tr>
				{{ end }}
			</tbody>
		</table>
	</div>
</div>
{{ end }}

{{ define "Scripts" }}

{{ end }}