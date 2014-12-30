{{ define "Content" }}
<div class="row">
    <div class="col-lg-12">
        <h1 class="page-header">
            Profile
        </h1>
    </div>
</div>
{{ with .user }}
<div class="row">
	<div class="col-lg-3">
		<img src="http://www.gravatar.com/avatar/{{ .Hash }}.png?s=200" />
		<br />
		<a href="#" data-toggle="modal" data-target="#gravatar">Change Image</a>
	</div>
	<div class="col-lg-5">
		<table class="table table-bordered table-hover table-striped">
			<tbody>
				<tr>
					<th>Name</th>
					<td>{{ .Name }}</td>
				</tr>
				<tr>
					<th>Email</th>
					<td>{{ .Email }}</td>
				</tr>
				<tr>
					<th>Hash</th>
					<td>{{ .Hash }}</td>
				</tr>
				<tr>
					<th>Token</th>
					<td>{{ .Token }}</td>
				</tr>
			</tbody>
		</table>
		<div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">Update Bio</h3>
			</div>
			<div class="panel-body">
				<form method="POST" action="/admin/user/profile/update/bio">
					<div class="form-group">
						<textarea cols="45" rows="12" name="bio">{{ .Bio }}</textarea>
					</div>
					<div class="form-group">
						<button type="submit" class="btn btn-primary">Update Bio</button>
					</div>
				</form>
			</div>
		</div>
	</div>
	<div class="col-lg-4">
		<div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">Regenerate Token</h3>
			</div>
			<div class="panel-body">
				<a style="width: 100%;" class="btn btn-primary" href="/admin/user/profile/token/regen">Go</a>
			</div>
		</div>
	</div>
	<div class="col-lg-4">
		<div class="panel panel-default">
			<div class="panel-heading">
				<h3 class="panel-title">Change Password</h3>
			</div>
			<div class="panel-body">
				<form method="POST" action="/admin/user/profile/update/password">
					<div class="form-group">
						<input type="password" name="old" class="form-control" placeholder="Old Password" />
					</div>
					<div class="form-group">
						<input type="password" name="new1" class="form-control" placeholder="New Password" />
					</div>
					<div class="form-group">
						<input type="password" name="new2" class="form-control" placeholder="Verify Password" />
					</div>
					<div class="form-group">
						<button type="submit" class="btn btn-primary">Update Password</button>
					</div>
				</form>
			</div>
		</div>
	</div>
</div>

<div class="modal fade" id="gravatar">
	<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
				<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
        		<h4 class="modal-title">Images Provided by Gravatar</h4>
			</div>
			<div class="modal-body">
				<p>We currently support gravatar for your profile images.</p>
				<p>Go to <a href="http://www.gravatar.com" target="_blank">Gravatar</a> and set up an account to change your avatar.</p>
			</div>
		</div>
	</div>
</div>
{{ else }}
	<p>{{ .error }}</p>
{{ end }}
{{ end }}

{{ define "Scripts" }}

{{ end }}