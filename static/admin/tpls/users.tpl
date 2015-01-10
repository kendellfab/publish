{{ define "Content" }}
<div class="row">
    <div class="col-lg-12">
        <h1 class="page-header">
            Users
        </h1>
    </div>
</div>
<div class="row">
	<div class="col-lg-3 col-lg-offset-9">
		<button type="button" class="btn btn-success pull-right" data-toggle="modal" data-target="#userModal">Add User</button>
		<br />
		<br />
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
						<th>Name</th>
						<th>Email</th>
						<th>Role</th>
						<th>Delete</th>
					</tr>
				</thead>
				<tbody>
					{{ range .users }}
						<tr>
							<td>{{ .Name }}</td>
							<td>{{ .Email }}</td>
							<td>{{ .Role }}</td>
							<td><a href="/admin/users/{{ .Id }}/delete">Delete</a></td>
						</tr>
					{{ end }}
				</tbody>
			</table>
		{{ end }}
	</div>
</div>
<div class="modal fade" id="userModal">
  <div class="modal-dialog modal-admin">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
        <h4 class="modal-title">Add User</h4>
      </div>
      <div class="modal-body">
    	<form method="POST" action="/admin/users/add">
    		<div class="form-group">
				<label for="name">Name</label>
				<input type="text" class="form-control" name="name" id="name" placeholder="Name" />
    		</div>
    		<div class="form-group">
				<label for="email">Email</label>
				<input type="text" class="form-control" name="email" id="email" placeholder="Email" />
    		</div>
    		<div class="form-group">
				<label for="password">Password</label>
				<input type="text" class="form-control" name="password" id="password" placeholder="Password" />
    		</div>
			<div class="form-group">
				<label for="role">Role</label>
				<select id="role" name="role" class="form-control">
					{{ range $k, $v := .roles }}
						<option value="{{ $k }}">{{ $v }}</option>
					{{ end }}
				</select>
			</div>
			<div class="form-group">
				<button type="submit" class="btn btn-primary">Save</button>
			</div>
    	</form>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->
{{ end }}

{{ define "Scripts" }}

{{ end }}