{{ define "Content" }}
<div class="row">
	<div class="col-lg-4">
		<h3>Series</h3>
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
{{ end }}

{{ define "Scripts" }}

{{ end }}