{{ define "Content" }}
<div class="row">
	<div class="col-lg-8">
		<h3>Posts</h3>
		<p>All: {{ len .posts }}<p>
	</div>
	<div class="col-lg-4">
		<div class="panel panel-default">
            <div class="panel-heading">
                <h3 class="panel-title"><i class="fa fa-fw fa-file-text-o"></i> New Post</h3>
            </div>
            <div class="panel-body">
				<form method="POST" action="/admin/post/start">
					<div class="form-group">
						<input type="text" name="title" class="form-control" placeholder="Post Title" />
					</div>
					<button type="submit" class="btn btn-primary">Save</button>
				</form>
			</div>
		</div>
	</div>
</div>
<div class="row">
	<div class="col-lg-12">
		{{ if .error }}
			<p>{{ .error }}</p>
		{{ else }}
			<table class="table table-bordered table-hover table-striped">
				<thead>
					<tr>
						<th>Title</th>
						<th>Author</th>
						<th>Created</th>
						<th>Category</th>
						<th>Published</th>
						<th>Edit</th>
					</tr>
				</thead>
				<tbody>
					{{ range .posts }}
						<tr>
							<td>{{ .Title }}</td>
							<td>{{ .Author.Name }}</td>
							<td>{{ fmt_date .Created }}</td>
							<td>{{ if .Category }}{{ .Category.Title }}{{ end }}</td>
							<td>{{ fmt_bool .Published }}</td>
							<td><a href="/admin/post/{{ .Id }}/edit">Edit</a></td>
						</tr>
					{{ end }}
				</tbody>
			</table>
		{{ end }}
	</div>
</div>
{{ end }}

{{ define "Scripts" }}

{{ end }}