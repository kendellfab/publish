{{ define "Content" }}
<div class="row">
	<div class="col-lg-12">
		<h3>Uploads</h3>
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
					<th>File Name</th>
				</tr>
			</thead>
			<tbody>
				{{ range .files }}
					<tr>
						<td>{{ .Name }}</td>
					</tr>
				{{ end }}
			</tbody>
		</table>
	</div>
</div>
{{ end }}

{{ define "Scripts" }}

{{ end }}