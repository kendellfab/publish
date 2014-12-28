{{ define "Content" }}
<div class="row">
    <div class="col-lg-12">
        <h1 class="page-header">
            Pages
        </h1>
    </div>
</div>
<div class="row">
	<div class="col-lg-4">
	</div>
	<div class="col-lg-8">
		<div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">New Page</h3>
			</div>
			<div class="panel-body">
				<form method="POST" action="/admin/pages/start">
					<div class="row">
						<div class="col-lg-10">
							<input type="text" name="title" class="form-control" placeholder="Page Title" />
						</div>
						<div class="col-lg-2">
							<button type="submit" class="btn btn-primary">Save</button>
						</div>
					</div>
				</form>
			</div>
		</div>
	</div>
</div>
<div class="row">
	<div class="col-lg-12">
		{{ with .pages }}
			<table class="table table-bordered table-hover table-striped">
				<thead>
					<tr>
						<th>Title</th>
						<th>Created</th>
						<th>Published</th>
						<th>Edit</th>
					</tr>
				</thead>
				<tbody>
					{{ range .}}
						<tr>
							<td>{{ .Title }}</td>
							<td>{{ fmt_date .Created }}</td>
							<td>{{ fmt_bool .Published }}</td>
							<td><a href="/admin/pages/{{ .Id }}/edit">Edit</a></td>
						</tr>
					{{ end }}
				</tbody>
			</table>
		{{ else }}
			<p>No pages right now.</p>
		{{ end }}
	</div>
</div>
{{ end }}

{{ define "Scripts" }}

{{ end }}