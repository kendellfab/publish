{{ define "Content" }}
<div class="row">
	<div class="col-lg-12">
		<h3>Categories</h3>
	</div>
</div>
<div class="row">
	<div class="col-lg-8">

		{{ if .error }}
			<p>{{ .error }}</p>
		{{ else }}
			<table class="table table-bordered table-hover table-striped">
				<thead>
					<tr>
						<th>Title</th>
						<th>Created</th>
					</tr>
				</thead>
				<tbody>
					{{ range .cats }}
						<tr>
							<td>{{ .Title }}</td>
							<td>{{ fmt_date .Created }}</td>
						</tr>
					{{ end }}
				</tbody>
			</table>
		{{ end }}

	</div>
	<div class="col-lg-4">
		<div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">New Category</h3>
			</div>
			<div class="panel-body">
				<form method="POST" action="/admin/cat/save">
					<div class="form-group">
						<input type="text" name="title" class="form-control" placeholder="Category Title" />
					</div>
					<button type="submit" class="btn btn-primary">Save</button>
				</form>
			</div>
		</div>
	</div>
</div>
{{ end }}

{{ define "Scripts" }}

{{ end }}