{{ define "Content" }}
<div class="row">
    <div class="col-lg-12">
        <h1 class="page-header">
            Series
        </h1>
    </div>
</div>
<div class="row">
	<div class="col-lg-4">
	</div>
	<div class="col-lg-8">
		<div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">New Series</h3>
			</div>
			<div class="panel-body">
				<form method="POST" action="/admin/series/start">
					<div class="row">
						<div class="col-lg-10">
							<input type="text" name="title" class="form-control" placeholder="Series Title" />
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
		{{ with .series }}
			<table class="table table-bordered table-hover table-striped">
				<thead>
					<tr>
						<th>Title</th>
						<th>Created</th>
						<th>Edit</th>
					</tr>
				</thead>
				<tbody>
					{{ range . }}
						<tr>
							<td>{{ .Title }}</td>
							<td>{{ fmt_date .Created }}</td>
							<td><a href="/admin/series/{{ .Id }}/edit">Edit</a></td>
						</tr>
					{{ end }}
				</tbody>
			</table>
		{{ else }}
			<p>No Series</p>
		{{ end }}
	</div>
</div>
{{ end }}

{{ define "Scripts" }}

{{ end }}